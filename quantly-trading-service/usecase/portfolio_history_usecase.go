package usecase

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/trmviet0801/quantly/convertor"
	"github.com/trmviet0801/quantly/database"
	"github.com/trmviet0801/quantly/dto"
	"github.com/trmviet0801/quantly/models"
	"github.com/trmviet0801/quantly/network"
	"github.com/trmviet0801/quantly/repos/model_repo"
)

func GetPortfolioHistoryOfAccount(accountId string) (*models.PortfolioHistory, error) {
	godotenv.Load()
	url := fmt.Sprintf("%s%s%s", os.Getenv("ALPACE_ORDER_BASE_URL"), accountId, os.Getenv("ALPACE_PORTFOLIO_HISTORY"))

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	response, err := network.SafeCall("GET", url, headers, nil)
	if err != nil {
		return nil, err
	}

	var portfolioHistoryDto dto.PortfolioHistoryDto
	isOk, err := network.OnResult(response, &portfolioHistoryDto)
	if isOk {
		portfolioHistory := convertor.ConvertDtoToModelPortfolioHistory(&portfolioHistoryDto, accountId)
		syncPortfolioHistoryWithDB(portfolioHistory)
		return portfolioHistory, nil
	}
	return nil, err
}

func syncPortfolioHistoryWithDB(history *models.PortfolioHistory) {
	portfolioHistoryRepo := model_repo.PortfolioHistoryRepo{
		DB: database.GetDatabase(),
	}

	portfolioHistoryRepo.Update(history)
}
