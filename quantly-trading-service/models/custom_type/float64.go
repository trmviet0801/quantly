package customtype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Float64Slice []float64

// Convert slice to json string to store in mysql
func (s Float64Slice) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// Convert data from db to struct
func (s *Float64Slice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("cannot convert %T to Float64Slice", value)
	}
	return json.Unmarshal(bytes, s)
}
