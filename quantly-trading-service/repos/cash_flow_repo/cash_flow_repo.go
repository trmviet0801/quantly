package cashflowrepo

import (
	"database/sql"
	"errors"

	"github.com/trmviet0801/quantly/models"
)

type CashFlowRepo struct {
	db *sql.DB
}

func (cashFlowRepo *CashFlowRepo) GetById(stockSymbol string) (cashFlow *models.CashFlow, err error) {
	query := `
	SELECT 
		stock_symbol,
		operating_cash_flow,
		investing_cash_flow,
		financing_cash_flow,
		end_cash_position,
		income_tax_paid_supplemental_data,
		interest_paid_supplemental_data,
		capital_expenditure,
		issuance_of_debt,
		repayment_of_debt,
		repurchase_of_capital_stock,
		free_cash_flow
	FROM cash_flows
	WHERE stock_symbol = ?`

	row := cashFlowRepo.db.QueryRow(query, stockSymbol)
	cashFlow = &models.CashFlow{}
	err = row.Scan(
		&cashFlow.StockSymbol,
		&cashFlow.OperatingCashFlow,
		&cashFlow.InvestingCashFlow,
		&cashFlow.FinancingCashFlow,
		&cashFlow.EndCashPosition,
		&cashFlow.IncomeTaxPaidSupplementalData,
		&cashFlow.InterestPaidSupplementalData,
		&cashFlow.CapitalExpenditure,
		&cashFlow.IssuanceOfDebt,
		&cashFlow.RepaymentOfDebt,
		&cashFlow.RepurchaseOfCapitalStock,
		&cashFlow.FreeCashFlow,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.New("cash flow not found for stock symbol: " + stockSymbol)
	}

	return cashFlow, nil
}

func (csshFlowRepo *CashFlowRepo) Create(cashFlow *models.CashFlow) error {
	query := `
	INSERT INTO cash_flows (
		stock_symbol,
		operating_cash_flow,
		investing_cash_flow,
		financing_cash_flow,
		end_cash_position,
		income_tax_paid_supplemental_data,
		interest_paid_supplemental_data,
		capital_expenditure,
		issuance_of_debt,
		repayment_of_debt,
		repurchase_of_capital_stock,
		free_cash_flow
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := csshFlowRepo.db.Exec(query,
		cashFlow.StockSymbol,
		cashFlow.OperatingCashFlow,
		cashFlow.InvestingCashFlow,
		cashFlow.FinancingCashFlow,
		cashFlow.EndCashPosition,
		cashFlow.IncomeTaxPaidSupplementalData,
		cashFlow.InterestPaidSupplementalData,
		cashFlow.CapitalExpenditure,
		cashFlow.IssuanceOfDebt,
		cashFlow.RepaymentOfDebt,
		cashFlow.RepurchaseOfCapitalStock,
		cashFlow.FreeCashFlow,
	)
	if err != nil {
		return errors.New("failed to create cash flow: " + err.Error())
	}
	return nil
}

func (cashFlowRepo *CashFlowRepo) Update(cashFlow *models.CashFlow) error {
	query := `
	UPDATE cash_flows SET
		operating_cash_flow = ?,
		investing_cash_flow = ?,
		financing_cash_flow = ?,
		end_cash_position = ?,
		income_tax_paid_supplemental_data = ?,
		interest_paid_supplemental_data = ?,
		capital_expenditure = ?,
		issuance_of_debt = ?,
		repayment_of_debt = ?,
		repurchase_of_capital_stock = ?,
		free_cash_flow = ?
	WHERE stock_symbol = ?`

	_, err := cashFlowRepo.db.Exec(query,
		cashFlow.OperatingCashFlow,
		cashFlow.InvestingCashFlow,
		cashFlow.FinancingCashFlow,
		cashFlow.EndCashPosition,
		cashFlow.IncomeTaxPaidSupplementalData,
		cashFlow.InterestPaidSupplementalData,
		cashFlow.CapitalExpenditure,
		cashFlow.IssuanceOfDebt,
		cashFlow.RepaymentOfDebt,
		cashFlow.RepurchaseOfCapitalStock,
		cashFlow.FreeCashFlow,
		cashFlow.StockSymbol,
	)
	if err != nil {
		return errors.New("failed to update cash flow: " + err.Error())
	}
	return nil
}

func (cashFlowRepo *CashFlowRepo) Delete(stockSymbol string) error {
	query := `DELETE FROM cash_flows WHERE stock_symbol = ?`
	_, err := cashFlowRepo.db.Exec(query, stockSymbol)
	if err != nil {
		return errors.New("failed to delete cash flow for stock symbol: " + stockSymbol + " - " + err.Error())
	}
	return nil
}
