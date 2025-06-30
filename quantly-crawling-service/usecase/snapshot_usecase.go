package usecase

import (
	"fmt"
	"os"

	"github.com/trmviet0801/quantly/quantly-crawling-serivce/models"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/network"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/utils"
)

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
