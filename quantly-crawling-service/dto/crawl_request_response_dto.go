package dto

// Response from BrightData after making POST new crawling process successfully
type CrawlRequestResponseDto struct {
	SnapshotId string `json:"snapshot_id"`
}
