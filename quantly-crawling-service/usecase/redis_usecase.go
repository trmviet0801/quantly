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

func PostSnapShotToRedisDB(snapshotOverview *models.SnapshotOverview, rdb *redis.Client) error {
	key := fmt.Sprintf("snapshotOverviews:%s", snapshotOverview.Id)
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
