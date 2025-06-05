package accountrepo

import (
	"database/sql"
	"errors"

	"github.com/trmviet0801/quantly/models"
)

type AccountRepo struct {
	db *sql.DB
}

func (accountRepo *AccountRepo) GetById(accountId string) (account *models.Account, err error) {
	query := `
	SELECT 
		account_id,
		user_id,
		account_number,
		status,
		crypto_status,
		currency,
		last_equity,
		created_at,
		account_type,
		enabled_assets,
		partner_user_id,
		funding_instructions_url,
		pattern_day_trader,
		kyc_completed_at,
		kyc_status,
		account_atype,
		management_status,
		clearing_broker,
		clearing_account_number
	FROM accounts
	WHERE account_id = ?`

	row := accountRepo.db.QueryRow(query, accountId)
	account = &models.Account{}
	err = row.Scan(
		&account.AccountId,
		&account.UserId,
		&account.AccountNumber,
		&account.Status,
		&account.CryptoStatus,
		&account.Currency,
		&account.LastEquity,
		&account.CreatedAt,
		&account.AccountType,
		&account.EnabledAssets,
		&account.PartnerUserId,
		&account.FundingInstructionsUrl,
		&account.PatternDayTrader,
		&account.KycCompletedAt,
		&account.KycStatus,
		&account.AccountAtype,
		&account.ManagementStatus,
		&account.ClearingBroker,
		&account.ClearingAccountNumber,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.New("account not found for account ID: " + accountId)
	}

	return account, nil
}

func (accountRepo *AccountRepo) Create(account *models.Account) error {
	query := `
	INSERT INTO accounts (
		user_id,
		account_number,
		status,
		crypto_status,
		currency,
		last_equity,
		created_at,
		account_type,
		enabled_assets,
		partner_user_id,
		funding_instructions_url,
		pattern_day_trader,
		kyc_completed_at,
		kyc_status,
		account_atype,
		management_status,
		clearing_broker,
		clearing_account_number
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := accountRepo.db.Exec(query,
		account.UserId,
		account.AccountNumber,
		account.Status,
		account.CryptoStatus,
		account.Currency,
		account.LastEquity,
		account.CreatedAt,
		account.AccountType,
		account.EnabledAssets,
		account.PartnerUserId,
		account.FundingInstructionsUrl,
		account.PatternDayTrader,
		account.KycCompletedAt,
		account.KycStatus,
		account.AccountAtype,
		account.ManagementStatus,
		account.ClearingBroker,
		account.ClearingAccountNumber,
	)
	if err != nil {
		return errors.New("failed to create account: " + err.Error())
	}
	return nil
}

func (accountRepo *AccountRepo) Update(account *models.Account) error {
	query := `
	UPDATE accounts SET
		user_id = ?,
		account_number = ?,
		status = ?,
		crypto_status = ?,
		currency = ?,
		last_equity = ?,
		created_at = ?,
		account_type = ?,
		enabled_assets = ?,
		partner_user_id = ?,
		funding_instructions_url = ?,
		pattern_day_trader = ?,
		kyc_completed_at = ?,
		kyc_status = ?,
		account_atype = ?,
		management_status = ?,
		clearing_broker = ?,
		clearing_account_number = ?
	WHERE account_id = ?`

	_, err := accountRepo.db.Exec(query,
		account.UserId,
		account.AccountNumber,
		account.Status,
		account.CryptoStatus,
		account.Currency,
		account.LastEquity,
		account.CreatedAt,
		account.AccountType,
		account.EnabledAssets,
		account.PartnerUserId,
		account.FundingInstructionsUrl,
		account.PatternDayTrader,
		account.KycCompletedAt,
		account.KycStatus,
		account.AccountAtype,
		account.ManagementStatus,
		account.ClearingBroker,
		account.ClearingAccountNumber,
		account.AccountId)
	if err != nil {
		return errors.New("failed to update account: " + err.Error())
	}
	return nil
}

func (accountRepo *AccountRepo) DeleteById(accountId string) error {
	query := `DELETE FROM accounts WHERE account_id = ?`
	_, err := accountRepo.db.Exec(query, accountId)
	if err != nil {
		return errors.New("failed to delete account: " + err.Error())
	}
	return nil
}
