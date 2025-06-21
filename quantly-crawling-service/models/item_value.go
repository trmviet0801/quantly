package models

type ItemValue struct {
	ColumnName string      `json:"column_name" redis:"column_name"`
	Value      interface{} `json:"value" redis:"value"`
}
