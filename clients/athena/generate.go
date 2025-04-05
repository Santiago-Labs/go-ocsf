package athena

import (
	"fmt"
	"log"
	"strings"

	"github.com/apache/arrow/go/v15/arrow"
)

func arrowTypeToAthena(t arrow.DataType) string {
	switch t.ID() {
	case arrow.STRING:
		return "string"
	case arrow.INT32:
		return "int"
	case arrow.INT64:
		return "bigint"
	case arrow.FLOAT32:
		return "float"
	case arrow.FLOAT64:
		return "double"
	case arrow.TIMESTAMP:
		return "timestamp"
	case arrow.BOOL:
		return "boolean"
	case arrow.STRUCT:
		structType := t.(*arrow.StructType)
		var fields []string
		for _, f := range structType.Fields() {
			fields = append(fields, fmt.Sprintf("%s:%s", f.Name, arrowTypeToAthena(f.Type)))
		}
		return fmt.Sprintf("struct<%s>", strings.Join(fields, ","))
	case arrow.LIST:
		listType := t.(*arrow.ListType)
		return fmt.Sprintf("array<%s>", arrowTypeToAthena(listType.Elem()))
	default:
		log.Fatalf("Unknown arrow type: %v", t.ID())
	}

	return "string"
}

func GenerateAthenaTable(fields []arrow.Field, tableName, s3Location string, partitions []string) string {
	var schemaDefs []string
	var partitionDefs []string

	partitionSet := make(map[string]bool)
	for _, p := range partitions {
		partitionSet[p] = true
	}

	for _, field := range fields {
		definition := fmt.Sprintf("  %s %s", field.Name, arrowTypeToAthena(field.Type))
		if partitionSet[field.Name] {
			partitionDefs = append(partitionDefs, definition)
		} else {
			schemaDefs = append(schemaDefs, definition)
		}
	}

	schemaStr := strings.Join(schemaDefs, ",\n")
	partitionStr := strings.Join(partitionDefs, ",\n")

	var partitionStmt string
	if len(partitions) > 0 {
		partitionStmt = fmt.Sprintf("PARTITIONED BY (\n%s\n)", partitionStr)
	}

	createStmt := fmt.Sprintf(`
	CREATE EXTERNAL TABLE %s (
	%s
	)
	%s
	STORED AS PARQUET
	LOCATION '%s';
	`, tableName, schemaStr, partitionStmt, s3Location)

	return createStmt
}
