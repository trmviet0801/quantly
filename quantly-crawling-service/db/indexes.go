package db

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/utils"
)

// IdxName: idx:snapshotOverviews
// Columns: [id, dataset_id, status]
func SetUpSnapshotOverviewIndex(rdb *redis.Client) error {
	ctx := context.Background()
	_, err := rdb.FTCreate(
		ctx,
		"idx:snapshotOverviews",
		&redis.FTCreateOptions{
			OnJSON: true,
			Prefix: []interface{}{"snapshot_overview:"},
		},
		&redis.FieldSchema{
			FieldName: "$.id",
			As:        "id",
			FieldType: redis.SearchFieldTypeText,
		},
		&redis.FieldSchema{
			FieldName: "$.dataset_id",
			As:        "dataset_id",
			FieldType: redis.SearchFieldTypeText,
		},
		&redis.FieldSchema{
			FieldName: "$.status",
			As:        "status",
			FieldType: redis.SearchFieldTypeTag,
		},
	).Result()

	if err != nil {
		utils.OnError(fmt.Errorf("can not create redis indexes: %w", err))
		return err
	}
	return nil
}

// IdxName = idx:stocks
// Columns: [company_id, name]
func SetUpStockIndex(rdb *redis.Client) error {
	ctx := context.Background()
	_, err := rdb.FTCreate(
		ctx,
		"idx:stocks",
		&redis.FTCreateOptions{
			OnJSON: true,
			Prefix: []interface{}{"stock:"},
		},
		&redis.FieldSchema{
			FieldName: "$.company_id",
			As:        "company_id",
			FieldType: redis.SearchFieldTypeText,
		},
		&redis.FieldSchema{
			FieldName: "$.name",
			As:        "name",
			FieldType: redis.SearchFieldTypeText,
		},
	).Result()

	if err != nil {
		utils.OnError(fmt.Errorf("can not create redis indexes: %w", err))
		return err
	}
	return nil
}

// IdxName = idx:snapshot
// Columns: [status, snapshot_id]
func SetUpSnapshotIndex(rdb *redis.Client) error {
	ctx := context.Background()
	_, err := rdb.FTCreate(
		ctx,
		"idx:snapshot",
		&redis.FTCreateOptions{
			OnJSON: true,
			Prefix: []interface{}{"snapshot:"},
		},
		&redis.FieldSchema{
			FieldName: "$.status",
			As:        "status",
			FieldType: redis.SearchFieldTypeText,
		},
		&redis.FieldSchema{
			FieldName: "$.snapshot_id",
			As:        "snapshot_id",
			FieldType: redis.SearchFieldTypeText,
		},
	).Result()

	if err != nil {
		utils.OnError(fmt.Errorf("can not create redis indexes: %w", err))
		return err
	}
	return nil
}
