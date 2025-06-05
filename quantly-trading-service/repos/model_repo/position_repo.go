package model_repo

import (
	"database/sql"
	"errors"

	"github.com/trmviet0801/quantly/models"
)

type PositionRepo struct {
	db *sql.DB
}

func (positionRepo *PositionRepo) GetById(positionId int64) (position *models.Position, err error) {
	query := `
	SELECT 
		position_id,
		account_id,
		asset_id,
		symbol,
		exchange,
		asset_class,
		asset_marginable,
		qty,
		avg_entry_price,
		side,
		market_value,
		cost_basis,
		unrealized_pl,
		unrealized_plpc,
		unrealized_intraday_pl,
		unrealized_intraday_plpc,
		current_price,
		lastday_price,
		change_today,
		qty_available
	FROM positions
	WHERE position_id = ?`

	row := positionRepo.db.QueryRow(query, positionId)
	position = &models.Position{}
	err = row.Scan(
		&position.PositionId,
		&position.AccountId,
		&position.AssetId,
		&position.Symbol,
		&position.Exchange,
		&position.AssetClass,
		&position.AssetMarginable,
		&position.Quantity,
		&position.AverageEntryPrice,
		&position.Side,
		&position.MarketValue,
		&position.CostBasis,
		&position.UnrealizedPL,
		&position.UnrealizedPLPC,
		&position.UnrealizedIntradayPL,
		&position.UnrealizedIntradayPLPC,
		&position.CurrentPrice,
		&position.LastDayPrice,
		&position.ChangeToday,
		&position.QuantityAvailable,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return position, nil
}

func (positionRepo *PositionRepo) Create(position *models.Position) error {
	query := `
	INSERT INTO positions (
		account_id,
		asset_id,
		symbol,
		exchange,
		asset_class,
		asset_marginable,
		qty,
		avg_entry_price,
		side,
		market_value,
		cost_basis,
		unrealized_pl,
		unrealized_plpc,
		unrealized_intraday_pl,
		unrealized_intraday_plpc,
		current_price,
		lastday_price,
		change_today,
		qty_available
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := positionRepo.db.Exec(query,
		position.AccountId,
		position.AssetId,
		position.Symbol,
		position.Exchange,
		position.AssetClass,
		position.AssetMarginable,
		position.Quantity,
		position.AverageEntryPrice,
		position.Side,
		position.MarketValue,
		position.CostBasis,
		position.UnrealizedPL,
		position.UnrealizedPLPC,
		position.UnrealizedIntradayPL,
		position.UnrealizedIntradayPLPC,
		position.CurrentPrice,
		position.LastDayPrice,
		position.ChangeToday,
		position.QuantityAvailable)
	if err != nil {
		return errors.New("failed to create position: " + err.Error())
	}
	return nil
}

func (positionRepo *PositionRepo) Update(position *models.Position) error {
	query := `
	UPDATE positions SET
		account_id = ?,
		asset_id = ?,
		symbol = ?,
		exchange = ?,
		asset_class = ?,
		asset_marginable = ?,
		qty = ?,
		avg_entry_price = ?,
		side = ?,
		market_value = ?,
		cost_basis = ?,
		unrealized_pl = ?,
		unrealized_plpc = ?,
		unrealized_intraday_pl = ?,
		unrealized_intraday_plpc = ?,
		current_price = ?,
		lastday_price = ?,
		change_today = ?,
		qty_available = ?
	WHERE position_id = ?`

	_, err := positionRepo.db.Exec(query,
		position.AccountId,
		position.AssetId,
		position.Symbol,
		position.Exchange,
		position.AssetClass,
		position.AssetMarginable,
		position.Quantity,
		position.AverageEntryPrice,
		position.Side,
		position.MarketValue,
		position.CostBasis,
		position.UnrealizedPL,
		position.UnrealizedPLPC,
		position.UnrealizedIntradayPL,
		position.UnrealizedIntradayPLPC,
		position.CurrentPrice,
		position.LastDayPrice,
		position.ChangeToday,
		position.QuantityAvailable,
		position.PositionId)
	if err != nil {
		return errors.New("failed to update position: " + err.Error())
	}
	return nil
}

func (positionRepo *PositionRepo) DeleteById(positionId int64) error {
	query := `DELETE FROM positions WHERE position_id = ?`
	_, err := positionRepo.db.Exec(query, positionId)
	if err != nil {
		return errors.New("failed to delete position: " + err.Error())
	}
	return nil
}
