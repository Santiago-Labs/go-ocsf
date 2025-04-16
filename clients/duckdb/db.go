package duckdb

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"sync"

	_ "github.com/marcboeker/go-duckdb/v2"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

var (
	db   *sqlx.DB
	once sync.Once

	reservedKeywords = map[string]bool{
		"type": true,
	}
)

func NewS3Client(ctx context.Context, bucketName string) (*sqlx.DB, error) {
	var err error
	once.Do(func() {

		var execDir string
		execDir, err = os.Executable()
		if err != nil {
			return
		}
		execDir = path.Dir(execDir)

		duckDBDir := path.Join(execDir, "duckdb_data")
		err = os.MkdirAll(duckDBDir, 0755)
		if err != nil {
			return
		}

		dbPath := path.Join(duckDBDir, "duckdb.db")
		db, err = sqlx.Connect("duckdb", dbPath)
		if err != nil {
			return
		}
		db.Mapper = reflectx.NewMapperTagFunc("json", nil, nil)

		_, err = db.Exec(fmt.Sprintf(`
			SET home_directory='%s'; 
			SET extension_directory='%s';

			FORCE INSTALL aws FROM core_nightly;
			FORCE INSTALL httpfs FROM core_nightly;
			FORCE INSTALL iceberg FROM core_nightly;

			CREATE SECRET (
				TYPE s3,
				PROVIDER credential_chain
			);

		`, duckDBDir, duckDBDir))
		if err != nil {
			return
		}
	})

	return db, err
}

func NewLocalClient(ctx context.Context) (*sqlx.DB, error) {
	var err error
	once.Do(func() {

		var execDir string
		execDir, err = os.Executable()
		if err != nil {
			return
		}
		execDir = path.Dir(execDir)

		duckDBDir := path.Join(execDir, "duckdb_data")
		err = os.MkdirAll(duckDBDir, 0755)
		if err != nil {
			return
		}

		dbPath := path.Join(duckDBDir, "duckdb.db")
		db, err = sqlx.Connect("duckdb", dbPath)
		if err != nil {
			return
		}
		db.Mapper = reflectx.NewMapperTagFunc("json", nil, nil)

		_, err = db.Exec(fmt.Sprintf(`
			SET home_directory='%s'; 
			SET extension_directory='%s';

		`, duckDBDir, duckDBDir))
		if err != nil {
			return
		}
	})

	return db, err
}

func arrowTypeToDuckDBType(t arrow.DataType) string {
	switch t.ID() {
	case arrow.STRING:
		return "VARCHAR"
	case arrow.INT32:
		return "INTEGER"
	case arrow.INT64:
		return "BIGINT"
	case arrow.FLOAT32:
		return "FLOAT"
	case arrow.FLOAT64:
		return "DOUBLE"
	case arrow.BOOL:
		return "BOOLEAN"
	case arrow.TIMESTAMP:
		tsType := t.(*arrow.TimestampType)
		switch tsType.Unit {
		case arrow.Second:
			return "TIMESTAMP_S"
		case arrow.Millisecond:
			return "TIMESTAMP_MS"
		case arrow.Microsecond:
			return "TIMESTAMP_US"
		case arrow.Nanosecond:
			return "TIMESTAMP_NS"
		default:
			log.Fatalf("Unknown timestamp unit: %v", tsType.Unit)
		}
	case arrow.LIST:
		listType := t.(*arrow.ListType)
		return fmt.Sprintf("%s[]", arrowTypeToDuckDBType(listType.Elem()))
	case arrow.STRUCT:
		structType := t.(*arrow.StructType)
		var fields []string
		for _, f := range structType.Fields() {
			fields = append(fields, fmt.Sprintf("%q %s", f.Name, arrowTypeToDuckDBType(f.Type)))
		}
		return fmt.Sprintf("STRUCT(%s)", strings.Join(fields, ", "))
	case arrow.MAP:
		mapType := t.(*arrow.MapType)
		return fmt.Sprintf("MAP(%s, %s)", arrowTypeToDuckDBType(mapType.KeyType()), arrowTypeToDuckDBType(mapType.ItemType()))
	case arrow.BINARY:
		return "BLOB"
	case arrow.FIXED_SIZE_BINARY:
		return "BLOB"
	case arrow.DATE32:
		return "DATE"
	case arrow.DATE64:
		return "TIMESTAMP_MS"
	default:
		log.Fatalf("Unsupported arrow type: %v", t)
	}

	return "VARCHAR" // default fallback
}

func GenerateDuckDBSelectFields(table, pattern string, fields []arrow.Field) string {
	var sb strings.Builder

	for i, field := range fields {
		comma := ","
		if i == len(fields)-1 {
			comma = ""
		}
		fieldType := arrowTypeToDuckDBType(field.Type)

		if field.Name == "filename" {
			continue
		}

		// Explicitly cast timestamp fields
		if fieldType == "TIMESTAMP_MS" {
			fmt.Fprintf(&sb, "  CAST(\"%s\" AS TIMESTAMP_MS) AS \"%s\"%s\n", field.Name, field.Name, comma)
		} else {
			fmt.Fprintf(&sb, "  \"%s\"%s\n", field.Name, comma)
		}
	}

	return sb.String()
}

func GenerateDuckDBColumnsDict(fields []arrow.Field) string {
	var sb strings.Builder

	sb.WriteString("columns = {\n")

	for i, field := range fields {
		comma := ","
		if i == len(fields)-1 {
			comma = ""
		}
		if field.Name == "filename" {
			continue
		}
		fmt.Fprintf(&sb, "  %s: '%s'%s\n", field.Name, arrowTypeToDuckDBType(field.Type), comma)
	}

	sb.WriteString("}")

	return sb.String()
}

func GenerateDuckDBNullView(table string, fields []arrow.Field) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("CREATE OR REPLACE VIEW %s AS\nSELECT\n", table))

	for i, field := range fields {
		comma := ","
		if i == len(fields)-1 {
			comma = ""
		}
		sb.WriteString(fmt.Sprintf("  NULL::%s AS \"%s\"%s\n", arrowTypeToDuckDBType(field.Type), field.Name, comma))
	}

	sb.WriteString("WHERE FALSE;")

	return sb.String()
}
