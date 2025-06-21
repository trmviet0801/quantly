package models

type Financial struct {
	Breakdown string      `json:"Breakdown" redis:"Breakdown"`
	Value     []ItemValue `json:"value" redis:"value"`
}
