package models

type CompanyKeyExecutive struct {
	Name      string  `json:"name" redis:"name"`
	Title     string  `json:"title" redis:"title"`
	Pay       string  `json:"pay" redis:"pay"`
	Exercised *string `json:"exercised" redis:"exercised"`
	YearBorn  *int    `json:"year_born" redis:"year_born"`
}
