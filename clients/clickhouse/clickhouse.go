package clickhouse

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/Santiago-Labs/go-ocsf/ocsf"
)

type Client struct {
	conn   driver.Conn
	dbName string
}

type Options struct {
	DSN      string
	Database string
	Username string
	Password string
}

func New(ctx context.Context, opts Options) (*Client, error) {
	var err error
	var conn driver.Conn

	once := sync.Once{}

	var clickhouseClient *Client

	var dbName string
	if opts.Database != "" {
		dbName = opts.Database
	} else {
		dbName = "telophase"
	}

	once.Do(func() {
		var auth clickhouse.Auth
		if opts.Username != "" && opts.Password != "" {
			auth = clickhouse.Auth{
				Database: opts.Database,
				Username: opts.Username,
				Password: opts.Password,
			}
		}
		if opts.DSN == "" {
			opts.DSN = "localhost:9000"
		}

		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{opts.DSN},
			Auth: auth,
		})
		if err != nil {
			return
		}

		clickhouseClient = &Client{
			conn:   conn,
			dbName: dbName,
		}

		err = clickhouseClient.createDB(ctx, dbName)
		if err != nil {
			return
		}

	})
	if err != nil {
		return nil, err
	}

	return clickhouseClient, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) createDB(ctx context.Context, db string) error {
	return c.conn.Exec(ctx, fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", db))
}

func (c *Client) Query(ctx context.Context, query string) (driver.Rows, error) {
	return c.conn.Query(ctx, query)
}

func (c *Client) DBName() string {
	return c.dbName
}

func (c *Client) InsertFindings(ctx context.Context, findings []ocsf.VulnerabilityFinding) error {
	if len(findings) == 0 {
		return nil
	}

	// Assuming you have a table for findings
	fullTableName := fmt.Sprintf("%s.%s", c.dbName, "vulnerability_findings")

	// Prepare batch for insertion
	batch, err := c.conn.PrepareBatch(ctx, fmt.Sprintf("INSERT INTO %s", fullTableName))
	if err != nil {
		return err
	}

	// Add each finding to the batch as a whole struct
	for _, finding := range findings {
		err = batch.AppendStruct(&finding)
		if err != nil {
			return fmt.Errorf("failed to append finding to batch: %w", err)
		}
	}

	// Execute the batch insertion
	if err := batch.Send(); err != nil {
		return fmt.Errorf("failed to send findings batch: %w", err)
	}

	return nil
}

func (c *Client) InsertAPIActivities(ctx context.Context, apiActivities []ocsf.APIActivity) error {
	if len(apiActivities) == 0 {
		return nil
	}

	fullTableName := fmt.Sprintf("%s.%s", c.dbName, "api_activities")

	// Prepare batch for insertion
	batch, err := c.conn.PrepareBatch(ctx, fmt.Sprintf("INSERT INTO %s", fullTableName))
	if err != nil {
		return err
	}

	type Row struct {
		ActivityID   int
		ActivityName *string
	}
	it := 0
	// Add each API activity to the batch as a whole struct
	for _, activity := range apiActivities {
		if activity.ActivityID == 0 {
			fmt.Println("ACTIVITY ID ISNT SET")
			continue
		}
		fmt.Println("ACTIVITY ID IS SET", it, "ACTIVITY ID", activity.ActivityID)
		err = batch.AppendStruct(&activity)
		if err != nil {
			return fmt.Errorf("failed to append API activity to batch: %w", err)
		}
	}

	// Execute the batch insertion
	if err := batch.Send(); err != nil {
		return fmt.Errorf("failed to send API activities batch: %w", err)
	}

	return nil
}

// CreateTableFromStruct creates a ClickHouse table based on the structure of a Go struct
func (c *Client) CreateTableFromStruct(ctx context.Context, tableName, primaryKey, orderBy string, structType interface{}) error {
	// If table exists, then we don't need to create it
	existsQuery := fmt.Sprintf("EXISTS TABLE %s.%s", c.dbName, tableName)
	fmt.Println("Checking if table exists...", existsQuery)
	exists, err := c.conn.Query(ctx, existsQuery)
	if err != nil {
		return fmt.Errorf("failed to check if table exists: %w", err)
	}
	var tableExists uint8
	if exists.Next() {
		if err := exists.Scan(&tableExists); err != nil {
			return fmt.Errorf("failed to scan exists result: %w", err)
		}
		if tableExists == 1 {
			fmt.Println("Table already exists, skipping creation")
			return nil
		}
	}

	// Don't forget to close the result set
	if err := exists.Close(); err != nil {
		return fmt.Errorf("failed to close exists query: %w", err)
	}

	schema, err := generateClickHouseSchema(structType)
	if err != nil {
		return fmt.Errorf("failed to generate schema: %w", err)
	}
	var orderByClause, primaryKeyClause string
	if primaryKey != "" {
		primaryKeyClause = fmt.Sprintf("PRIMARY KEY %s", primaryKey)
	}
	if orderBy != "" {
		orderByClause = fmt.Sprintf("ORDER BY %s", orderBy)
	}

	query := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s.%s
		(
			%s
		) ENGINE = MergeTree() %s %s
	`, c.dbName, tableName, schema, orderByClause, primaryKeyClause)

	fmt.Println("Creating table...\n", query)

	if err := c.conn.Exec(ctx, query); err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}

// generateClickHouseSchema converts a Go struct to ClickHouse column definitions
func generateClickHouseSchema(structType interface{}) (string, error) {
	t := reflect.TypeOf(structType)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("expected struct, got %s", t.Kind())
	}

	columns, err := generateColumns(t, "")
	if err != nil {
		return "", err
	}

	return strings.Join(columns, ",\n"), nil
}

// generateColumns recursively generates column definitions for a struct type
func generateColumns(t reflect.Type, prefix string) ([]string, error) {
	var columns []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Skip unexported fields
		if field.PkgPath != "" {
			continue
		}

		// Get the column name from the json tag or use the field name
		columnName := field.Name
		// Add prefix if we're in a nested field
		if prefix != "" {
			columnName = prefix + columnName
		}

		switch field.Type.Kind() {
		case reflect.Struct:
			if field.Type.String() == "time.Time" {
				columns = append(columns, fmt.Sprintf("`%s` Nullable(DateTime64(9, 'UTC'))", columnName))
			} else {
				// For nested structs, create a Tuple
				nestedColumns, err := generateNestedTuple(field.Type)
				if err != nil {
					return nil, err
				}
				columns = append(columns, fmt.Sprintf("`%s` %s", columnName, nestedColumns))
			}
		case reflect.Slice, reflect.Array:
			elemType := field.Type.Elem()
			if elemType.Kind() == reflect.Struct {
				// For arrays of structs, create Array(Tuple(...))
				nestedColumns, err := generateNestedTuple(elemType)
				if err != nil {
					return nil, err
				}
				columns = append(columns, fmt.Sprintf("`%s` Array(%s)", columnName, nestedColumns))
			} else if elemType.Kind() == reflect.Ptr && elemType.Elem().Kind() == reflect.Struct {
				// For arrays of struct pointers
				nestedColumns, err := generateNestedTuple(elemType.Elem())
				if err != nil {
					return nil, err
				}
				columns = append(columns, fmt.Sprintf("`%s` Array(%s)", columnName, nestedColumns))
			} else {
				// For arrays of primitive types
				clickhouseType, err := goTypeToClickHouseType(elemType)
				if err != nil {
					return nil, err
				}
				columns = append(columns, fmt.Sprintf("`%s` Array(Nullable(%s))", columnName, clickhouseType))
			}
		case reflect.Ptr:
			// For pointer types, make them Nullable
			elemType := field.Type.Elem()
			if elemType.Kind() == reflect.Struct {
				if elemType.String() == "time.Time" {
					columns = append(columns, fmt.Sprintf("`%s` Nullable(DateTime64(9, 'UTC'))", columnName))
				} else {
					nestedColumns, err := generateNestedTuple(elemType)
					if err != nil {
						return nil, err
					}
					columns = append(columns, fmt.Sprintf("`%s` %s", columnName, nestedColumns))
				}
			} else {
				clickhouseType, err := goTypeToClickHouseType(elemType)
				if err != nil {
					return nil, err
				}
				columns = append(columns, fmt.Sprintf("`%s` Nullable(%s)", columnName, clickhouseType))
			}
		case reflect.Map:
			// Maps are serialized as JSON strings
			columns = append(columns, fmt.Sprintf("`%s` String", columnName))
		default:
			clickhouseType, err := goTypeToClickHouseType(field.Type)
			if err != nil {
				return nil, err
			}
			columns = append(columns, fmt.Sprintf("`%s` %s", columnName, clickhouseType))
		}
	}

	return columns, nil
}

// generateNestedTuple generates a ClickHouse Tuple type for a nested struct
func generateNestedTuple(t reflect.Type) (string, error) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("expected struct for Tuple, got %s", t.Kind())
	}

	var fields []string
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Skip unexported fields
		if field.PkgPath != "" {
			continue
		}

		// Get the field name from the json tag or use the field name
		fieldName := field.Name

		switch field.Type.Kind() {
		case reflect.Struct:
			if field.Type.String() == "time.Time" {
				fields = append(fields, fmt.Sprintf("%s Nullable(DateTime64(9, 'UTC'))", fieldName))
			} else {
				nestedTuple, err := generateNestedTuple(field.Type)
				if err != nil {
					return "", err
				}
				// Don't wrap Tuple in Nullable
				fields = append(fields, fmt.Sprintf("`%s` %s", fieldName, nestedTuple))
			}
		case reflect.Slice, reflect.Array:
			elemType := field.Type.Elem()
			if elemType.Kind() == reflect.Struct || (elemType.Kind() == reflect.Ptr && elemType.Elem().Kind() == reflect.Struct) {
				actualType := elemType
				if elemType.Kind() == reflect.Ptr {
					actualType = elemType.Elem()
				}
				nestedTuple, err := generateNestedTuple(actualType)
				if err != nil {
					return "", err
				}
				// Don't wrap Array(Tuple) in Nullable
				fields = append(fields, fmt.Sprintf("`%s` Array(%s)", fieldName, nestedTuple))
			} else {
				clickhouseType, err := goTypeToClickHouseType(elemType)
				if err != nil {
					return "", err
				}
				fields = append(fields, fmt.Sprintf("`%s` Array(Nullable(%s))", fieldName, clickhouseType))
			}
		case reflect.Ptr:
			elemType := field.Type.Elem()
			if elemType.Kind() == reflect.Struct {
				if elemType.String() == "time.Time" {
					fields = append(fields, fmt.Sprintf("`%s` Nullable(DateTime64(9, 'UTC'))", fieldName))
				} else {
					nestedTuple, err := generateNestedTuple(elemType)
					if err != nil {
						return "", err
					}
					// Don't wrap Tuple in Nullable
					fields = append(fields, fmt.Sprintf("`%s` %s", fieldName, nestedTuple))
				}
			} else {
				clickhouseType, err := goTypeToClickHouseType(elemType)
				if err != nil {
					return "", err
				}
				fields = append(fields, fmt.Sprintf("`%s` Nullable(%s)", fieldName, clickhouseType))
			}
		default:
			clickhouseType, err := goTypeToClickHouseType(field.Type)
			if err != nil {
				return "", err
			}
			fields = append(fields, fmt.Sprintf("`%s` Nullable(%s)", fieldName, clickhouseType))
		}
	}

	return fmt.Sprintf("Tuple(%s)", strings.Join(fields, ", ")), nil
}

// goTypeToClickHouseType maps Go types to ClickHouse data types
func goTypeToClickHouseType(t reflect.Type) (string, error) {
	switch t.Kind() {
	case reflect.Bool:
		return "Bool", nil
	case reflect.Int, reflect.Int64:
		return "Int64", nil
	case reflect.Int8, reflect.Int16, reflect.Int32:
		return "Int32", nil
	case reflect.Uint8, reflect.Uint16, reflect.Uint32:
		return "UInt32", nil
	case reflect.Uint, reflect.Uint64:
		return "UInt64", nil
	case reflect.Float32:
		return "Float32", nil
	case reflect.Float64:
		return "Float64", nil
	case reflect.String:
		return "String", nil
	case reflect.Struct:
		if t.String() == "time.Time" {
			return "DateTime64(9, 'UTC')", nil
		}
		return "String", nil // Fallback for other structs
	case reflect.Map:
		return "String", nil // Maps as JSON strings
	case reflect.Ptr:
		return goTypeToClickHouseType(t.Elem())
	default:
		return "", fmt.Errorf("unsupported type: %s", t.Kind())
	}
}
