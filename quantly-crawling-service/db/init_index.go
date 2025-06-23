package db

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/utils"
)

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
