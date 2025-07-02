package models

import "fmt"

type SnapshotOverview struct {
	Id        string `json:"id"`
	DatasetId string `json:"dataset_id"`
	Status    string `json:"status"`
	Created   string `json:"created"`
}

// If not running -> ready
func (so *SnapshotOverview) IsRunning() bool {
	return so.Status == "running"
}

func (so *SnapshotOverview) IsReady() bool {
	return so.Status == "ready"
}

func (s SnapshotOverview) String() string {
	return fmt.Sprintf(
		"SnapshotOverview{Id: %q, DatasetId: %q, Status: %q, Created: %q}",
		s.Id, s.DatasetId, s.Status, s.Created,
	)
}
