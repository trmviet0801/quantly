package usecase

import (
	"fmt"
	"os"

	"github.com/trmviet0801/quantly/quantly-crawling-serivce/models"
	"github.com/trmviet0801/quantly/quantly-crawling-serivce/network"
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
