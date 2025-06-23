package models

import "fmt"

type Snapshot struct {
	Status             string       `json:"status"`
	SnapshotID         string       `json:"snapshot_id"`
	DatasetID          string       `json:"dataset_id"`
	ErrorCodes         ErrorCodeMap `json:"error_codes"`
	Records            int          `json:"records"`
	Errors             int          `json:"errors"`
	CollectionDuration int64        `json:"collection_duration"`
}

func (s Snapshot) String() string {
	return fmt.Sprintf(
		"Snapshot:\n  Status: %s\n  SnapshotID: %s\n  DatasetID: %s\n  Records: %d\n  Errors: %d\n  CollectionDuration: %d ms\n  ErrorCodes: %s",
		s.Status,
		s.SnapshotID,
		s.DatasetID,
		s.Records,
		s.Errors,
		s.CollectionDuration,
		s.ErrorCodes.String(),
	)
}
