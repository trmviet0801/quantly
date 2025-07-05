package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/models"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/utils"
)

func PostStockToRedisDB(stock *models.Stock, rdb *redis.Client) error {
	key := fmt.Sprintf("stock:%s", stock.CompanyID)
	ctx := context.Background()

	_, err := rdb.JSONSet(ctx, key, "$", stock).Result()
	rdb.Expire(ctx, key, 1*time.Hour)
	if err != nil {
		utils.OnError(err)
		return err
	}
	return nil
}

// Post latest version of snapshot_overview to Redis
// Only 1 version of snapshot_overview in Redis (latest version)
func PostSnapshotOverviewToRedisDB(snapshotOverview *models.SnapshotOverview, rdb *redis.Client) error {
	key := "snapshot_overview:latest"
	ctx := context.Background()

	_, err := rdb.JSONSet(ctx, key, "$", snapshotOverview).Result()
	rdb.Expire(ctx, key, 1*time.Hour)
	if err != nil {
		utils.OnError(err)
		return err
	}
	return nil
}

func PostSnapshotToRedisDB(snapshot *models.Snapshot, rdb *redis.Client) error {
	key := fmt.Sprintf("snapshot:%s", snapshot.SnapshotID)
	ctx := context.Background()

	_, err := rdb.JSONSet(ctx, key, "$", snapshot).Result()
	rdb.Expire(ctx, key, 1*time.Hour)
	if err != nil {
		utils.OnError(err)
		return err
	}
	return nil
}

func FindStocksByCompanyID(companyId string, rdb *redis.Client) ([]*models.Stock, error) {
	ctx := context.Background()

	redisStocks, err := rdb.FTSearch(ctx, "idx:stocks", fmt.Sprintf("@company_id:{%s}", companyId)).Result()
	if err != nil {
		utils.OnError(err)
		return nil, err
	}

	stocks, err := utils.UnmarshallRedisReturn[models.Stock](&redisStocks)
	if err != nil {
		return nil, err
	}

	return stocks, nil
}

// Get by key: snapshotOverviews:latest from Redis
func GetLatestSnapshot(rdb *redis.Client) ([]*models.SnapshotOverview, error) {
	ctx := context.Background()

	redisSnapshotOverviews, err := rdb.JSONGet(ctx, "snapshot_overview:latest", "$").Result()
	if err != nil {
		utils.OnError(err)
		return nil, err
	}

	if redisSnapshotOverviews == "" {
		return nil, redis.Nil
	}

	snapshotOverviews, err := utils.UnmarshallRedisReturnString[models.SnapshotOverview](redisSnapshotOverviews)
	if err != nil {
		return nil, err
	}
	return snapshotOverviews, nil
}

func FindSnapshotById(snapshotId string, rdb *redis.Client) ([]*models.Snapshot, error) {
	ctx := context.Background()

	redisSnapshot, err := rdb.FTSearch(ctx, "idx:snapshot", fmt.Sprintf("@snapshot_id:{%s}", snapshotId)).Result()
	if err != nil {
		utils.OnError(err)
		return nil, err
	}

	snapshots, err := utils.UnmarshallRedisReturn[models.Snapshot](&redisSnapshot)
	if err != nil {
		return nil, err
	}
	return snapshots, nil
}
