package model_repo

import (
	"database/sql"
	"errors"

	"github.com/trmviet0801/quantly/models"
)

type BalanceSheetRepo struct {
	db *sql.DB
}

func (balanceSheetRepo *BalanceSheetRepo) GetById(stockSymbol string) (balanceSheet *models.BalanceSheet, err error) {
	query := `
	SELECT 
		stock_symbol,
		total_assets,
		total_liabilities_net_minority_interest,
		total_equity_gross_minority_interest,
		total_capitalization,
		common_stock_equity,
		capital_lease_obligations,
		net_tangible_assets,
		working_capital,
		invested_capital,
		tangible_book_value,
		total_debt,
		net_debt,
		share_issued,
		ordinary_shares_number
	FROM balance_sheets
	WHERE stock_symbol = ?`

	row := balanceSheetRepo.db.QueryRow(query, stockSymbol)
	balanceSheet = &models.BalanceSheet{}
	err = row.Scan(
		&balanceSheet.StockSymbol,
		&balanceSheet.TotalAssets,
		&balanceSheet.TotalLiabilitiesNetMinorityInterest,
		&balanceSheet.TotalEquityGrossMinorityInterest,
		&balanceSheet.TotalCapitalization,
		&balanceSheet.CommonStockEquity,
		&balanceSheet.CapitalLeaseObligations,
		&balanceSheet.NetTangibleAssets,
		&balanceSheet.WorkingCapital,
		&balanceSheet.InvestedCapital,
		&balanceSheet.TangibleBookValue,
		&balanceSheet.TotalDebt,
		&balanceSheet.NetDebt,
		&balanceSheet.ShareIssued,
		&balanceSheet.OrdinarySharesNumber,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.New("balance sheet not found for stock symbol: " + stockSymbol)
	}

	return balanceSheet, nil

}

func (balanceSheetRepo *BalanceSheetRepo) Create(balanceSheet *models.BalanceSheet) error {
	query := `
	INSERT INTO balance_sheets (
		stock_symbol,
		total_assets,
		total_liabilities_net_minority_interest,
		total_equity_gross_minority_interest,
		total_capitalization,
		common_stock_equity,
		capital_lease_obligations,
		net_tangible_assets,
		working_capital,
		invested_capital,
		tangible_book_value,
		total_debt,
		net_debt,
		share_issued,
		ordinary_shares_number
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := balanceSheetRepo.db.Exec(query,
		balanceSheet.StockSymbol,
		balanceSheet.TotalAssets,
		balanceSheet.TotalLiabilitiesNetMinorityInterest,
		balanceSheet.TotalEquityGrossMinorityInterest,
		balanceSheet.TotalCapitalization,
		balanceSheet.CommonStockEquity,
		balanceSheet.CapitalLeaseObligations,
		balanceSheet.NetTangibleAssets,
		balanceSheet.WorkingCapital,
		balanceSheet.InvestedCapital,
		balanceSheet.TangibleBookValue,
		balanceSheet.TotalDebt,
		balanceSheet.NetDebt,
		balanceSheet.ShareIssued,
		balanceSheet.OrdinarySharesNumber)
	if err != nil {
		return errors.New("failed to create balance sheet: " + err.Error())
	}
	return nil
}

func (balanceSheetRepo *BalanceSheetRepo) Update(balanceSheet *models.BalanceSheet) error {
	query := `
	UPDATE balance_sheets SET
		total_assets = ?,
		total_liabilities_net_minority_interest = ?,
		total_equity_gross_minority_interest = ?,
		total_capitalization = ?,
		common_stock_equity = ?,
		capital_lease_obligations = ?,
		net_tangible_assets = ?,
		working_capital = ?,
		invested_capital = ?,
		tangible_book_value = ?,
		total_debt = ?,
		net_debt = ?,
		share_issued = ?,
		ordinary_shares_number = ?
	WHERE stock_symbol = ?`
	_, err := balanceSheetRepo.db.Exec(query,
		balanceSheet.TotalAssets,
		balanceSheet.TotalLiabilitiesNetMinorityInterest,
		balanceSheet.TotalEquityGrossMinorityInterest,
		balanceSheet.TotalCapitalization,
		balanceSheet.CommonStockEquity,
		balanceSheet.CapitalLeaseObligations,
		balanceSheet.NetTangibleAssets,
		balanceSheet.WorkingCapital,
		balanceSheet.InvestedCapital,
		balanceSheet.TangibleBookValue,
		balanceSheet.TotalDebt,
		balanceSheet.NetDebt,
		balanceSheet.ShareIssued,
		balanceSheet.OrdinarySharesNumber,
		balanceSheet.StockSymbol)
	if err != nil {
		return errors.New("failed to update balance sheet: " + err.Error())
	}
	return nil
}

func (balanceSheetRepo *BalanceSheetRepo) Delete(stockSymbol string) error {
	query := `DELETE FROM balance_sheets WHERE stock_symbol = ?`
	_, err := balanceSheetRepo.db.Exec(query, stockSymbol)
	if err != nil {
		return errors.New("failed to delete balance sheet: " + err.Error())
	}
	return nil
}
