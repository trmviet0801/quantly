package usecase

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/trmviet0801/quantly/convertor"
	"github.com/trmviet0801/quantly/database"
	"github.com/trmviet0801/quantly/dto"
	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/network"
	"github.com/trmviet0801/quantly/repos/model_repo"
)

func GetOpenPositionsOfAccount(accountId string) ([]*models.Position, error) {
	godotenv.Load()

	url := fmt.Sprintf("%s%s%s", os.Getenv("ALPACE_ORDER_BASE_URL"), accountId, os.Getenv("ALPACE_POSITION"))

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	response, err := network.SafeCall("GET", url, headers, nil)
	if err != nil {
		return nil, err
	}

	var positionsDtos []dto.PositionDto
	isOk, err := network.OnResult(response, &positionsDtos)
	if isOk {
		positions := convertor.ConvertDtosToPositions(&positionsDtos, accountId)
		syncPositionsToDB(positions)
		return positions, nil
	}
	return nil, err
}

func CloseAllPositions(accountId string) ([]*dto.PositionCloseAllResponseDto, error) {
	godotenv.Load()

	url := fmt.Sprintf("%s%s%s", os.Getenv("ALPACE_ORDER_BASE_URL"), accountId, os.Getenv("ALPACE_POSITION"))

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	response, err := network.SafeCall("DELETE", url, headers, nil)
	if err != nil {
		return nil, err
	}

	var positionCLoseAllResponseDtos []dto.PositionCloseAllResponseDto
	isOk, err := network.OnResult(response, &positionCLoseAllResponseDtos)
	if isOk {
		result := convertor.ArrayToArrayOfPointer(positionCLoseAllResponseDtos)
		return result, nil
	}
	return nil, err
}

func syncPositionsToDB(positions []*models.Position) {
	for _, position := range positions {
		syncPositionToDB(position)
	}
}

func syncPositionToDB(position *models.Position) {
	positionRepo := model_repo.PositionRepo{
		DB: database.GetDatabase(),
	}
	positionDB, _ := positionRepo.GetById(position.PositionId)
	if positionDB == nil {
		position.PositionId = uuid.New().String()
		positionRepo.Create(position)
	} else {
		positionRepo.Update(position)
	}
}
