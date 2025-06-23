package models

import "fmt"

type ErrorCodeMap struct {
	CrawlFailed int `json:"crawl_failed"`
	DeadPage    int `json:"dead_page"`
}

func (e ErrorCodeMap) String() string {
	return fmt.Sprintf("CrawlFailed: %d, DeadPage: %d", e.CrawlFailed, e.DeadPage)
}
