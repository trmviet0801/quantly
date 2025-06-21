CREATE TABLE `users` (
  `user_id` VARCHAR(64) NOT NULL PRIMARY KEY,
  `username` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL
);

CREATE TABLE `accounts` (
  `account_id` VARCHAR(64) NOT NULL PRIMARY KEY,
  `user_id` VARCHAR(64) NOT NULL,
  `account_number` VARCHAR(255),
  `status` VARCHAR(50),
  `crypto_status` VARCHAR(50),
  `currency` VARCHAR(10),
  `last_equity` VARCHAR(50),
  `created_at` DATETIME,
  `account_type` VARCHAR(50),
  `enabled_assets` JSON,
  `partner_user_id` VARCHAR(255),
  `funding_instructions_url` TEXT,
  `pattern_day_trader` TINYINT(1),
  `kyc_completed_at` DATETIME,
  `kyc_status` VARCHAR(50),
  `account_atype` VARCHAR(50),
  `management_status` VARCHAR(50),
  `clearing_broker` VARCHAR(255),
  `clearing_account_number` VARCHAR(255)
);

CREATE TABLE `kyc_results` (
  `kyc_result_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `account_id` VARCHAR(64) NOT NULL,
  `status` VARCHAR(50),
  `internal_status` VARCHAR(50),
  `timestamp` DATETIME
);

CREATE TABLE `contacts` (
  `contact_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `account_id` VARCHAR(64) NOT NULL,
  `email_address` VARCHAR(255),
  `phone_number` VARCHAR(50),
  `street_address` JSON,
  `city` VARCHAR(100),
  `state` VARCHAR(100),
  `postal_code` VARCHAR(50),
  `country` VARCHAR(100),
  `given_name` VARCHAR(100),
  `middle_name` VARCHAR(100),
  `family_name` VARCHAR(100)
);

CREATE TABLE `trading_configurations` (
  `trading_configuration_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `account_id` VARCHAR(64) NOT NULL,
  `dtbp_check` VARCHAR(50),
  `no_shorting` TINYINT(1),
  `fractional_trading` TINYINT(1),
  `max_margin_multiplier` VARCHAR(50)
);

CREATE TABLE `trusted_contacts` (
  `trusted_contact_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `account_id` VARCHAR(64) NOT NULL,
  `given_name` VARCHAR(100),
  `family_name` VARCHAR(100),
  `email_address` VARCHAR(255)
);

CREATE TABLE `disclosures` (
  `disclosures_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `account_id` VARCHAR(64) NOT NULL,
  `is_control_person` TINYINT(1),
  `is_affiliated_exchange_or_finra` TINYINT(1),
  `is_politically_exposed` TINYINT(1),
  `immediate_family_exposed` TINYINT(1)
);

CREATE TABLE `identities` (
  `identity_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `account_id` VARCHAR(64) NOT NULL,
  `tax_id` VARCHAR(100),
  `tax_id_type` VARCHAR(50),
  `country_of_citizenship` VARCHAR(100),
  `country_of_birth` VARCHAR(100),
  `country_of_tax_residence` VARCHAR(100),
  `funding_source` JSON,
  `date_of_birth` DATE
);

CREATE TABLE `orders` (
  `order_id` VARCHAR(64) NOT NULL PRIMARY KEY,
  `account_id` VARCHAR(64) NOT NULL,
  `symbol` VARCHAR(10),
  `qty` DECIMAL(20,4),
  `notional` DECIMAL(30,4),
  `side` VARCHAR(10),
  `type` VARCHAR(20),
  `time_in_force` VARCHAR(10),
  `limit_price` DECIMAL(20,4),
  `stop_price` DECIMAL(20,4),
  `trail_price` DECIMAL(20,4),
  `trail_percent` DECIMAL(10,6),   
  `extended_hours` TINYINT(1),     -- boolean flag (0 or 1)
  `client_order_id` VARCHAR(255),
  `order_class` VARCHAR(50)
);

CREATE TABLE `stop_losses` (
  `stop_loss_id` VARCHAR(64) NOT NULL PRIMARY KEY,
  `order_id` VARCHAR(64) NOT NULL,
  `stop_price` DECIMAL(30,4),
  `limit_price` DECIMAL(30,4)
);

CREATE TABLE `take_profits` (
  `take_profit_id` VARCHAR(64) NOT NULL PRIMARY KEY,
  `order_id` VARCHAR(64) NOT NULL,
  `limit_price` DECIMAL(30, 4)
);

CREATE TABLE `quant_models` (
  `quant_model_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(255)
);

CREATE TABLE `trade_signals` (
  `trade_signal_id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `type` VARCHAR(50),
  `stock_symbol` VARCHAR(10),
  `quant` VARCHAR(50),
  `timestamp` DATETIME,
  `quant_model_id` BIGINT UNSIGNED,
  `account_id` VARCHAR(64) NOT NULL
);

CREATE TABLE `positions` (
  `position_id` VARCHAR(64) NOT NULL PRIMARY KEY,
  `account_id` VARCHAR(64) NOT NULL,
  `asset_id` VARCHAR(50),
  `symbol` VARCHAR(10),
  `exchange` VARCHAR(50),
  `asset_class` VARCHAR(50),
  `asset_marginable` TINYINT(1), 
  `qty` DECIMAL(20,4),
  `avg_entry_price` DECIMAL(20,4),
  `side` VARCHAR(10),
  `market_value` DECIMAL(30,4),
  `cost_basis` DECIMAL(30,4),
  `unrealized_pl` DECIMAL(30,4),
  `unrealized_plpc` DECIMAL(10,6),          
  `unrealized_intraday_pl` DECIMAL(30,4),
  `unrealized_intraday_plpc` DECIMAL(10,6), 
  `current_price` DECIMAL(20,4),
  `lastday_price` DECIMAL(20,4),
  `change_today` DECIMAL(20,4),
  `qty_available` DECIMAL(20,4)
);


CREATE TABLE `portfolio` (
  `portfolio_id` VARCHAR(64) NOT NULL PRIMARY KEY,
  `account_id` VARCHAR(64) NOT NULL,
  `current_value` VARCHAR(20),
  `profit_loss` VARCHAR(20)
);

CREATE TABLE `notifications` (
  `notification_id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `userid` VARCHAR(64) NOT NULL,
  `type` VARCHAR(50),
  `content` TEXT
);

CREATE TABLE `stocks` (
  `stock_symbol` VARCHAR(10) PRIMARY KEY,
  `name` VARCHAR(255),
  `ipo_year` INT,
  `country` VARCHAR(100),
  `current_price` DECIMAL(20,4),
  `price_change` DECIMAL(20,4),
  `change_percent` DECIMAL(6,4), 
  `open_price` DECIMAL(20,4),
  `day_range` VARCHAR(50), 
  `day_low` DECIMAL(20,4),
  `day_high` DECIMAL(20,4),
  `volume` BIGINT, 
  `latest_trade_time` DATETIME,
  `ticker` VARCHAR(20),
  `exchange` VARCHAR(50),
  `industry` VARCHAR(100),
  `sector` VARCHAR(100),
  `employees` INT,
  `headquarters` VARCHAR(255),
  `market_cap` DECIMAL(30,4), 
  `pe_ratio_ttm` DECIMAL(10,4),
  `eps_ttm` DECIMAL(20,4),
  `operating_margin` DECIMAL(6,4), 
  `dividend_yield` DECIMAL(6,4)
);


CREATE TABLE `balance_sheets` (
  `stock_symbol` VARCHAR(10),
  `total_assets` DECIMAL(20,4),
  `total_liabilities_net_minority_interest` DECIMAL(20,4),
  `total_equity_gross_minority_interest` DECIMAL(20,4),
  `total_capitalization` DECIMAL(20,4),
  `common_stock_equity` DECIMAL(20,4),
  `capital_lease_obligations` DECIMAL(20,4),
  `net_tangible_assets` DECIMAL(20,4),
  `working_capital` DECIMAL(20,4),
  `invested_capital` DECIMAL(20,4),
  `tangible_book_value` DECIMAL(20,4),
  `total_debt` DECIMAL(20,4),
  `net_debt` DECIMAL(20,4),
  `share_issued` DECIMAL(20,4),
  `ordinary_shares_number` DECIMAL(20,4)
);

CREATE TABLE `cash_flows` (
  `stock_symbol` VARCHAR(10),
  `operating_cash_flow` DECIMAL(20,4),
  `investing_cash_flow` DECIMAL(20,4),
  `financing_cash_flow` DECIMAL(20,4),
  `end_cash_position` DECIMAL(20,4),
  `income_tax_paid_supplemental_data` DECIMAL(20,4),
  `interest_paid_supplemental_data` DECIMAL(20,4),
  `capital_expenditure` DECIMAL(20,4),
  `issuance_of_debt` DECIMAL(20,4),
  `repayment_of_debt` DECIMAL(20,4),
  `repurchase_of_capital_stock` DECIMAL(20,4),
  `free_cash_flow` DECIMAL(20,4)
);

CREATE TABLE `incomes` (
  `stock_symbol` VARCHAR(10),
  `total_revenue` DECIMAL(20,4),
  `operating_revenue` DECIMAL(20,4),
  `cost_of_revenue` DECIMAL(20,4),
  `gross_profit` DECIMAL(20,4),
  `operating_expense` DECIMAL(20,4),
  `selling_general_and_administrative` DECIMAL(20,4),
  `research_and_development` DECIMAL(20,4),
  `operating_income` DECIMAL(20,4),
  `net_non_operating_interest_income_expense` DECIMAL(20,4),
  `interest_income_non_operating` DECIMAL(20,4),
  `interest_expense_non_operating` DECIMAL(20,4),
  `other_income_expense` DECIMAL(20,4),
  `special_income_charges` DECIMAL(20,4),
  `restructuring_and_mergers_acquisition` DECIMAL(20,4),
  `other_non_operating_income_expenses` DECIMAL(20,4),
  `pretax_income` DECIMAL(20,4),
  `tax_provision` DECIMAL(20,4),
  `net_income_common_stockholders` DECIMAL(20,4),
  `net_income` DECIMAL(20,4),
  `net_income_including_non_controlling_interests` DECIMAL(20,4),
  `net_income_continuous_operations` DECIMAL(20,4),
  `diluted_ni_available_to_com_stockholders` DECIMAL(20,4),
  `basic_eps` DECIMAL(20,4),
  `diluted_eps` DECIMAL(20,4),
  `basic_average_shares` DECIMAL(20,4),
  `diluted_average_shares` DECIMAL(20,4),
  `total_operating_income_as_reported` DECIMAL(20,4),
  `total_expenses` DECIMAL(20,4),
  `net_income_from_continuing_and_discontinued_operation` DECIMAL(20,4),
  `normalized_income` DECIMAL(20,4),
  `interest_income` DECIMAL(20,4),
  `interest_expense` DECIMAL(20,4),
  `net_interest_income` DECIMAL(20,4),
  `ebit` DECIMAL(20,4),
  `ebitda` DECIMAL(20,4),
  `reconciled_cost_of_revenue` DECIMAL(20,4),
  `reconciled_depreciation` DECIMAL(20,4),
  `net_income_from_continuing_operation_net_minority_interest` DECIMAL(20,4),
  `total_unusual_items_excluding_goodwill` DECIMAL(20,4),
  `total_unusual_items` DECIMAL(20,4),
  `normalized_ebitda` DECIMAL(20,4),
  `tax_rate_for_calcs` DECIMAL(6,4),
  `tax_effect_of_unusual_items` DECIMAL(20,4)
);


-- Foreign Keys
ALTER TABLE `accounts` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`);
ALTER TABLE `kyc_results` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`);
ALTER TABLE `contacts` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`);
ALTER TABLE `trading_configurations` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`);
ALTER TABLE `trusted_contacts` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`);
ALTER TABLE `disclosures` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`);
ALTER TABLE `identities` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`);
ALTER TABLE `orders` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`);
ALTER TABLE `stop_losses` ADD FOREIGN KEY (`order_id`) REFERENCES `orders` (`order_id`);
ALTER TABLE `take_profits` ADD FOREIGN KEY (`order_id`) REFERENCES `orders` (`order_id`);
ALTER TABLE `trade_signals` ADD FOREIGN KEY (`quant_model_id`) REFERENCES `quant_models` (`quant_model_id`);
ALTER TABLE `trade_signals` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`);
ALTER TABLE `positions` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`);
ALTER TABLE `portfolio` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`account_id`);
ALTER TABLE `notifications` ADD FOREIGN KEY (`userid`) REFERENCES `users` (`user_id`);
ALTER TABLE `balance_sheets` ADD FOREIGN KEY (`stock_symbol`) REFERENCES `stocks` (`stock_symbol`);
ALTER TABLE `cash_flows` ADD FOREIGN KEY (`stock_symbol`) REFERENCES `stocks` (`stock_symbol`);
ALTER TABLE `incomes` ADD FOREIGN KEY (`stock_symbol`) REFERENCES `stocks` (`stock_symbol`);
