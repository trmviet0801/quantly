package customtype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Int64Slice []int64

// convert slice to json for storing inside mysql db
func (s Int64Slice) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// convert db value to struct
func (s *Int64Slice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("cannot convert %T to Int64Slice", value)
	}
	return json.Unmarshal(bytes, s)
}
