package stockrepo

import (
	"database/sql"
	"errors"

	"github.com/trmviet0801/quantly/models"
)

type StockRepo struct {
	db *sql.DB
}

func (stockRepo *StockRepo) GetById(id string) (stock *models.Stock, err error) {
	query := `SELECT name, stock_symbol, ipo_year, country, current_price, price_change, change_percent,
       open_price, day_range, day_low, day_high, volume, latest_trade_time, ticker,
       exchange, industry, sector, employees, headquarters, market_cap, pe_ratio_ttm, eps_ttm FROM stocks WHERE stock_symbol = ?`
	stock = &models.Stock{}
	err = stockRepo.db.QueryRow(query, id).Scan(
		&stock.Name,
		&stock.Symbol,
		&stock.IpoYear,
		&stock.Country,
		&stock.CurrentPrice,
		&stock.PriceChange,
		&stock.ChangePercent,
		&stock.OpenPrice,
		&stock.DayRange,
		&stock.DayLow,
		&stock.DayHigh,
		&stock.Volume,
		&stock.LatestTradeTime,
		&stock.Ticker,
		&stock.Exchange,
		&stock.Industry,
		&stock.Sector,
		&stock.Employees,
		&stock.Headquarters,
		&stock.MarketCap,
		&stock.PERatioTtm,
		&stock.EPSTtm,
	)
	if err != nil {
		return nil, errors.New("stock not found")
	}
	return stock, nil
}

func (stockRepo *StockRepo) Create(stock *models.Stock) error {
	query := `INSERT INTO stocks (name, stock_symbol, ipo_year, country, current_price, price_change, change_percent,
	   open_price, day_range, day_low, day_high, volume, latest_trade_time, ticker,
	   exchange, industry, sector, employees, headquarters, market_cap, pe_ratio_ttm, eps_ttm)
	   VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := stockRepo.db.Exec(query,
		stock.Name,
		stock.Symbol,
		stock.IpoYear,
		stock.Country,
		stock.CurrentPrice,
		stock.PriceChange,
		stock.ChangePercent,
		stock.OpenPrice,
		stock.DayRange,
		stock.DayLow,
		stock.DayHigh,
		stock.Volume,
		stock.LatestTradeTime,
		stock.Ticker,
		stock.Exchange,
		stock.Industry,
		stock.Sector,
		stock.Employees,
		stock.Headquarters,
		stock.MarketCap,
		stock.PERatioTtm,
		stock.EPSTtm,
	)
	if err != nil {
		return errors.New("failed to create stock")
	}
	return nil
}

func (stockRepo *StockRepo) Update(stock *models.Stock) error {
	query := `UPDATE stocks SET name = ?, ipo_year = ?, country = ?, current_price = ?, price_change = ?,
	   change_percent = ?, open_price = ?, day_range = ?, day_low = ?, day_high = ?, volume = ?,
	   latest_trade_time = ?, ticker = ?, exchange = ?, industry = ?, sector = ?, employees = ?,
	   headquarters = ?, market_cap = ?, pe_ratio_ttm = ?, eps_ttm = ? WHERE stock_symbol = ?`
	_, err := stockRepo.db.Exec(query,
		stock.Name,
		stock.IpoYear,
		stock.Country,
		stock.CurrentPrice,
		stock.PriceChange,
		stock.ChangePercent,
		stock.OpenPrice,
		stock.DayRange,
		stock.DayLow,
		stock.DayHigh,
		stock.Volume,
		stock.LatestTradeTime,
		stock.Ticker,
		stock.Exchange,
		stock.Industry,
		stock.Sector,
		stock.Employees,
		stock.Headquarters,
		stock.MarketCap,
		stock.PERatioTtm,
		stock.EPSTtm,
		stock.Symbol)
	if err != nil {
		return errors.New("failed to update stock")
	}
	return nil
}

func (stockRepo *StockRepo) DeleteById(id string) error {
	query := `DELETE FROM stocks WHERE stock_symbol = ?`
	_, err := stockRepo.db.Exec(query, id)
	if err != nil {
		return errors.New("failed to delete stock")
	}
	return nil
}
