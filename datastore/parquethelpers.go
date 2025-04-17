package datastore

import (
	"fmt"
	"reflect"

	"github.com/apache/arrow-go/v18/arrow/array"
)

func buildRecord(rb *array.RecordBuilder, row any) error {
	v := reflect.ValueOf(row)
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("expected struct, got %T", row)
	}

	for i := 0; i < v.NumField(); i++ {
		fieldVal := v.Field(i)
		b := rb.Field(i)

		switch fieldVal.Kind() {
		case reflect.Pointer:
			if fieldVal.IsNil() {
				b.(array.Builder).AppendNull()
				continue
			}
			fieldVal = fieldVal.Elem()

		case reflect.Int32:
			b.(*array.Int32Builder).Append(int32(fieldVal.Int()))
		case reflect.Int64:
			b.(*array.Int64Builder).Append(fieldVal.Int())
		case reflect.String:
			b.(*array.StringBuilder).AppendString(fieldVal.String())
		case reflect.Struct:
			sb := b.(*array.StructBuilder)
			sb.Append(true)
			if err := buildSubStruct(sb, fieldVal); err != nil {
				return err
			}
		case reflect.Slice:
			lb := b.(*array.ListBuilder)
			lb.Append(true)
			if err := appendList(lb, fieldVal); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unhandled kind %v in field %d", fieldVal.Kind(), i)
		}
	}
	return nil
}

func buildSubStruct(sb *array.StructBuilder, v reflect.Value) error {
	for i := 0; i < v.NumField(); i++ {
		fld := v.Field(i)
		child := sb.FieldBuilder(i)
		switch fld.Kind() {
		case reflect.Int32:
			child.(*array.Int32Builder).Append(int32(fld.Int()))
		case reflect.Int64:
			child.(*array.Int64Builder).Append(fld.Int())
		case reflect.String:
			child.(*array.StringBuilder).AppendString(fld.String())
		default:
			return fmt.Errorf("unsupported nested kind %v", fld.Kind())
		}
	}
	return nil
}

func appendList(lb *array.ListBuilder, s reflect.Value) error {
	vb := lb.ValueBuilder()
	for i := 0; i < s.Len(); i++ {
		sb := vb.(*array.StructBuilder)
		sb.Append(true)
		if err := buildSubStruct(sb, s.Index(i)); err != nil {
			return err
		}
	}
	return nil
}
