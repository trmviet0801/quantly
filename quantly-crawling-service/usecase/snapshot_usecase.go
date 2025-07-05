package usecase

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/models"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/network"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/utils"
	"go.uber.org/zap"
)

// Get all snapshot_overviews on BrightData
func GetSnapshotOverviews() ([]*models.SnapshotOverview, error) {
	url := fmt.Sprintf("%s%s", os.Getenv("BRIGHT_DATA_BASE_URL"), os.Getenv("BRIGHT_DATA_GET_ALL_SNAPSHOTS_OVERVIEW_SUB_URL"))

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("BRIGHT_DATA_BEARER_TOKEN")),
		"Content-Type":  "application/json",
	}

	response, err := network.SafeCall(url, "GET", headers, nil)
	if err != nil {
		return nil, err
	}

	snapshotOverviews, err := network.Result[[]models.SnapshotOverview](response)
	if err != nil {
		return nil, err
	}
	return utils.ToPointerArray(snapshotOverviews), nil
}

func GetRunningSnapshot(snapshotOverviews []*models.SnapshotOverview) ([]*models.SnapshotOverview, error) {
	if len(snapshotOverviews) == 0 {
		err := fmt.Errorf("snapshotoverviews is empty")
		utils.OnError(err)
		return nil, err
	}
	var result []*models.SnapshotOverview
	for _, snapshot := range snapshotOverviews {
		if snapshot.IsRunning() {
			result = append(result, snapshot)
		}
	}
	return result, nil
}

func GetReadySnapshotOverviews(snapshotOverviews []*models.SnapshotOverview) ([]*models.SnapshotOverview, error) {
	if len(snapshotOverviews) == 0 {
		err := fmt.Errorf("snapshotoverviews is empty")
		utils.OnError(err)
		return nil, err
	}
	var result []*models.SnapshotOverview
	for _, snapshot := range snapshotOverviews {
		if snapshot.IsReady() {
			result = append(result, snapshot)
		}
	}
	return result, nil
}

// Get latest snapshot_overview in BrightData
func GetLatestReadySnapshotOverviewOfFullStocksInBD(conn *redis.Client) (*models.SnapshotOverview, error) {
	snapshotOverviews, err := GetSnapshotOverviews()
	if err != nil {
		return nil, err
	}

	// Sorted by timestamp by default from BrightData | DESC order
	readySnapshotOverviews, err := GetReadySnapshotOverviews(snapshotOverviews)
	if err != nil {
		return nil, err
	}
	if len(readySnapshotOverviews) == 0 {
		err := fmt.Errorf("no ready snapshot")
		utils.OnError(err)
		return nil, err
	}

	for _, snapshotOverview := range readySnapshotOverviews {
		if snapshotOverview.DatasetSize > 1 {
			return snapshotOverview, nil
		}
	}

	err = fmt.Errorf("no ready snapshot overview for full stocks")
	zap.L().Error(err.Error())
	return nil, err
}

// Get all ready snapshot overviews of single stock
// Return empty list if no snapshot overview of single stock
func GetReadySnapshotOverviewOfSingleStock(conn *redis.Client) ([]*models.SnapshotOverview, error) {
	snapshotOverviews, err := GetSnapshotOverviews()
	if err != nil {
		err = fmt.Errorf("can not get all snapshot overviews from BrightData: %s", err.Error())
		utils.OnError(err)
		return nil, err
	}

	var result []*models.SnapshotOverview
	for _, snapshotOverview := range snapshotOverviews {
		if snapshotOverview.IsRunning() {
			result = append(result, snapshotOverview)
		}
	}

	return result, nil
}

// Get latest snapshot_overview from BrightData -> Post to Redis
func PostLatestSnapshotOverview(conn *redis.Client) error {
	latestSnapshotOverview, err := GetLatestReadySnapshotOverviewOfFullStocksInBD(conn)
	if err != nil {
		return err
	}

	err = PostSnapshotOverviewToRedisDB(latestSnapshotOverview, conn)
	if err != nil {
		return err
	}
	zap.L().Info("Redis Insertion Successfully", zap.String("Data", "Latest Snapshot Overview"))
	return nil
}
