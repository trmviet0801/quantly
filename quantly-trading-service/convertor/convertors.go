package convertor

import (
	"strconv"

	"github.com/trmviet0801/quantly/dto"
	"github.com/trmviet0801/quantly/models"
)

// ConvertDtoToPosition maps a PositionDto to a Position model.
func ConvertDtoToPosition(dto *dto.PositionDto, accountId string) *models.Position {
	return &models.Position{
		AccountId:              accountId,
		AssetId:                dto.AssetId,
		Symbol:                 dto.Symbol,
		Exchange:               dto.Exchange,
		AssetClass:             dto.AssetClass,
		AssetMarginable:        dto.AssetMarginable,
		Quantity:               parseFloat(dto.Quantity),
		AverageEntryPrice:      parseFloat(dto.AverageEntryPrice),
		Side:                   dto.Side,
		MarketValue:            parseFloat(dto.MarketValue),
		CostBasis:              parseFloat(dto.CostBasis),
		UnrealizedPL:           parseFloat(dto.UnrealizedPL),
		UnrealizedPLPC:         parseFloat(dto.UnrealizedPLPC),
		UnrealizedIntradayPL:   parseFloat(dto.UnrealizedIntradayPL),
		UnrealizedIntradayPLPC: parseFloat(dto.UnrealizedIntradayPLPC),
		CurrentPrice:           parseFloat(dto.CurrentPrice),
		LastDayPrice:           parseFloat(dto.LastDayPrice),
		ChangeToday:            parseFloat(dto.ChangeToday),
		QuantityAvailable:      parseFloat(dto.QuantityAvailable),
	}
}

// ConvertPositionToDto maps a Position model to a PositionDto with string fields.
func ConvertPositionToDto(position *models.Position) *dto.PositionDto {
	return &dto.PositionDto{
		AssetId:                position.AssetId,
		Symbol:                 position.Symbol,
		Exchange:               position.Exchange,
		AssetClass:             position.AssetClass,
		AssetMarginable:        position.AssetMarginable,
		Quantity:               formatFloat(position.Quantity),
		AverageEntryPrice:      formatFloat(position.AverageEntryPrice),
		Side:                   position.Side,
		MarketValue:            formatFloat(position.MarketValue),
		CostBasis:              formatFloat(position.CostBasis),
		UnrealizedPL:           formatFloat(position.UnrealizedPL),
		UnrealizedPLPC:         formatFloat(position.UnrealizedPLPC),
		UnrealizedIntradayPL:   formatFloat(position.UnrealizedIntradayPL),
		UnrealizedIntradayPLPC: formatFloat(position.UnrealizedIntradayPLPC),
		CurrentPrice:           formatFloat(position.CurrentPrice),
		LastDayPrice:           formatFloat(position.LastDayPrice),
		ChangeToday:            formatFloat(position.ChangeToday),
		QuantityAvailable:      formatFloat(position.QuantityAvailable),
	}
}

// converts float64 to string with 6-digit precision
func formatFloat(val float64) string {
	return strconv.FormatFloat(val, 'f', -1, 64)
}
func ConvertDtosToPositions(dtos *[]dto.PositionDto, accountId string) []*models.Position {
	positions := make([]*models.Position, 0, len(*dtos))
	for _, dto := range *dtos {
		position := ConvertDtoToPosition(&dto, accountId)
		position.AccountId = accountId
		positions = append(positions, position)
	}
	return positions
}

func parseFloat(val string) float64 {
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0
	}
	return f
}

func ArrayToArrayOfPointer[T any](arr []T) []*T {
	result := make([]*T, len(arr))

	for i, item := range arr {
		val := item
		result[i] = &val
	}

	return result
}

func ConvertDtoToModelPortfolioHistory(dto *dto.PortfolioHistoryDto, accountId string) *models.PortfolioHistory {
	return &models.PortfolioHistory{
		AccountId:     accountId,
		Timestamp:     dto.Timestamp,
		Equity:        dto.Equity,
		ProfitLoss:    dto.ProfitLoss,
		ProfitLossPct: dto.ProfitLossPct,
		BaseValue:     dto.BaseValue,
		BaseValueAsof: dto.BaseValueAsof,
		Timeframe:     dto.Timeframe,
	}
}

func ConvertModelToDtoPortfolioHistory(model *models.PortfolioHistory) *dto.PortfolioHistoryDto {
	return &dto.PortfolioHistoryDto{
		Timestamp:     model.Timestamp,
		Equity:        model.Equity,
		ProfitLoss:    model.ProfitLoss,
		ProfitLossPct: model.ProfitLossPct,
		BaseValue:     model.BaseValue,
		BaseValueAsof: model.BaseValueAsof,
		Timeframe:     model.Timeframe,
	}
}
