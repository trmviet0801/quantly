package usecase

import (
	"fmt"
	"os"

	"github.com/trmviet0801/quantly/quantly-crawling-serivce/models"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/network"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/utils"
)

func GetProcessStatus(snapshotId string) (*models.Snapshot, error) {
	url := fmt.Sprintf("%s%s%s", os.Getenv("BRIGHT_DATA_BASE_URL"), os.Getenv("BRIGHT_DATA_GET_SNAPSHOT_STATUS_SUB_URL"), snapshotId)
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("BRIGHT_DATA_BEARER_TOKEN")),
		"Content-Type":  "application/json",
	}

	response, err := network.SafeCall(url, "GET", headers, nil)
	if err != nil {
		return nil, err
	}

	responseBody, err := network.Result[models.Snapshot](response)
	if err != nil {
		return nil, err
	}

	return &responseBody, nil
}

func GetAllSnapshotOverview() ([]*models.SnapshotOverview, error) {
	url := fmt.Sprintf("%s%s", os.Getenv("BRIGHT_DATA_BASE_URL"), os.Getenv("BRIGHT_DATA_GET_ALL_SNAPSHOTS_OVERVIEW_SUB_URL"))
	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", os.Getenv("BRIGHT_DATA_BEARER_TOKEN")),
	}

	response, err := network.SafeCall(url, "GET", headers, nil)
	if err != nil {
		utils.OnError(fmt.Errorf("can not get all snapshots overview: url %s | err: %w", url, err))
		return nil, err
	}

	body, err := network.Result[[]models.SnapshotOverview](response)
	if err != nil {
		utils.OnError(fmt.Errorf("can not unmarshall snapshots overview: url %s | err: %w", url, err))
		return nil, err
	}

	return utils.ToPointerArray(body), nil
}

// Returns (nil, nil) if dataset is empty
func GetLatestSnapshotOverview() (*models.SnapshotOverview, error) {
	snapshotoverviews, err := GetAllSnapshotOverview()
	if err != nil {
		return nil, err
	}

	if len(snapshotoverviews) == 0 {
		return nil, nil
	}
	return snapshotoverviews[0], nil
}
