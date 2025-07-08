-- database creation
create database quantly;

-- table createtion
create table users (
	user_id varchar(36) PRIMARY KEY NOT NULL,
    username varchar(36),
    `password` varchar(36),
    created_at timestamp
);

create table accounts (
	account_id varchar(36) PRIMARY KEY NOT NULL,
    account_number varchar(36),
    `status` varchar(10),
    currency varchar(10),
    last_equity varchar(20),
    create_at timestamp,
    account_type varchar(20),
    trading_type varchar(20),
    user_id varchar(36),
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

create table stocks (
	symbol varchar(10) PRIMARY KEY NOT NULL,
    short_name varchar(50),
    long_name varchar(100),
    sector varchar(50),
    industry varchar(50),
    market_cap bigint,
    current_price decimal(10,4),
    trailing_pe decimal(10,4),
    forward_pe decimal(10,4),
    price_to_book decimal(10,4),
    dividend_yield decimal(10,4),
    beta decimal(10, 4),
    fifty_two_week_high decimal(10,4),
    fifty_two_week_low decimal(10,4),
    average_volume bigint,
    last_close_price decimal(10,4),
    last_volume bigint,
    currency varchar(24),
    `exchange` varchar(24),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
create table stock_prices (
	stock_price_id varchar(36) primary key not null,
    symbol varchar(10),
    `date`datetime,
    `open` decimal(10,4),
    `close` decimal(10,4),
    high decimal(10,4),
    low decimal(10,4),
    volume bigint,
    foreign key (symbol) references stocks(symbol)
);

create table trading_models (
	trading_model_id varchar(36) primary key not null,
    `name` varchar(50),
    `description` text,
    `status` varchar(36),
    account_id varchar(36),
    formula text,
    entry_condition text,
    exit_condition text,
    foreign key (account_id) references accounts(account_id)
);

CREATE TABLE positions (
    position_id VARCHAR(36) PRIMARY KEY NOT NULL,
    symbol VARCHAR(10) NOT NULL,
    account_id VARCHAR(36) NOT NULL,
    quantity INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    price DECIMAL(10,4),
    `type` VARCHAR(10),      -- e.g. 'buy', 'sell', etc.
    `status` VARCHAR(10),    -- e.g. 'open', 'closed'

    FOREIGN KEY (account_id) REFERENCES accounts(account_id) ON DELETE CASCADE,
    FOREIGN KEY (symbol) REFERENCES stocks(symbol) ON DELETE CASCADE
);
