package ocsf

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type DBTime struct {
	time.Time
}

func (t *DBTime) Scan(value interface{}) error {
	if value == nil {
		*t = DBTime{Time: time.Time{}}
		return nil
	}

	str, ok := value.(string)
	if !ok {
		return nil
	}

	parsed, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}

	*t = DBTime{parsed}
	return nil
}

func (t DBTime) Value() (driver.Value, error) {
	return t.Time.Format(time.RFC3339), nil
}

type JSONB json.RawMessage

func (j *JSONB) Scan(src interface{}) error {
	if src == nil {
		*j = JSONB("null")
		return nil
	}

	switch data := src.(type) {
	case []byte:
		if len(data) == 0 {
			*j = JSONB("null")
		} else {
			*j = JSONB(data)
		}
	case string:
		if data == "" {
			*j = JSONB("null")
		} else {
			*j = JSONB(data)
		}
	case map[string]interface{}:
		if len(data) == 0 {
			*j = JSONB("null")
		} else {
			bytes, err := json.Marshal(data)
			if err != nil {
				return fmt.Errorf("failed to marshal map[string]interface{}: %w", err)
			}
			*j = JSONB(bytes)
		}
	case []interface{}:
		if len(data) == 0 {
			*j = JSONB("null")
		} else {
			bytes, err := json.Marshal(data)
			if err != nil {
				return fmt.Errorf("failed to marshal []interface{}: %w", err)
			}
			*j = JSONB(bytes)
		}
	default:
		return fmt.Errorf("unsupported type: %T", src)
	}

	return nil
}

func (j JSONB) MarshalJSON() ([]byte, error) {
	if len(j) == 0 {
		return []byte("null"), nil
	}
	return j, nil
}

func (j *JSONB) UnmarshalJSON(data []byte) error {
	if j == nil {
		return fmt.Errorf("JSONB: UnmarshalJSON on nil pointer")
	}
	*j = append((*j)[0:0], data...)
	return nil
}
