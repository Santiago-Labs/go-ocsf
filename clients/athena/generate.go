package athena

import (
	"fmt"
	"log"
	"strings"

	"github.com/apache/arrow/go/v15/arrow"
)

func GenerateAthenaCreateTable(database, table string, fields []arrow.Field) string {
	var sb strings.Builder

	fmt.Fprintf(&sb, "CREATE TABLE IF NOT EXISTS `%s`.%s (\n", database, table)

	for i, field := range fields {
		comma := ","
		if i == len(fields)-1 {
			comma = ""
		}

		if field.Name == "event_day" {
			fmt.Fprintf(&sb, "  `%s` DATE%s\n", field.Name, comma)
		} else {
			fmt.Fprintf(&sb, "  `%s` %s%s\n", field.Name, arrowTypeToAthenaType(field.Type), comma)
		}
	}

	fmt.Fprintf(&sb, ")\n")
	fmt.Fprintf(&sb, "PARTITIONED BY (event_day)\n")
	fmt.Fprintf(&sb, "TBLPROPERTIES ('table_type'='ICEBERG');")

	return sb.String()
}

func arrowTypeToAthenaType(t arrow.DataType) string {
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
	case arrow.DATE64:
		return "date"
	case arrow.LIST:
		listType := t.(*arrow.ListType)
		return fmt.Sprintf("array<%s>", arrowTypeToAthenaType(listType.Elem()))
	case arrow.STRUCT:
		structType := t.(*arrow.StructType)
		var fields []string
		for _, f := range structType.Fields() {
			fields = append(fields, fmt.Sprintf("`%s`:%s", f.Name, arrowTypeToAthenaType(f.Type)))
		}
		return fmt.Sprintf("struct<%s>", strings.Join(fields, ","))
	case arrow.MAP:
		mapType := t.(*arrow.MapType)
		return fmt.Sprintf("map<%s,%s>", arrowTypeToAthenaType(mapType.KeyType()), arrowTypeToAthenaType(mapType.ItemType()))
	default:
		log.Fatalf("Unknown arrow type: %v", t.ID())
	}

	return "string"
}
