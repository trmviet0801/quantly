package models

import "fmt"

type SnapshotOverview struct {
	Id          string `json:"id" gorm:"primaryKey;column:snapshot_overview_id"`
	DatasetId   string `json:"dataset_id" gorm:"column:dataset_id"`
	Status      string `json:"status" gorm:"snapshot_status"`
	DatasetSize int    `json:"dataset_size gorm:dataset_size"`
	Created     string `json:"created gorm:created"`
}

// If not running -> ready
func (so *SnapshotOverview) IsRunning() bool {
	return so.Status == "running"
}

func (so *SnapshotOverview) IsReady() bool {
	return so.Status == "ready"
}

func (s *SnapshotOverview) String() string {
	return fmt.Sprintf(
		"SnapshotOverview{Id: %q, DatasetId: %q, Status: %q, Created: %q}",
		s.Id, s.DatasetId, s.Status, s.Created,
	)
}

// Snapshot that contains only 1 in data set is snapshot for single stock
func (s *SnapshotOverview) IsSingleStock() bool {
	return s.DatasetSize == 1
}
