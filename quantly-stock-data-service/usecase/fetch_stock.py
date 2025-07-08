import yfinance
import yfinance as yf
from models.stock import Stock
from logger.structlog_config import get_logger
from converters.stock_converter import ticker_to_stock

# Fetch single stock via symbol
# Couldn't find stock -> returns NoneType
def fetch_stock(symbol: str) -> Stock | None:
    try:
        ticker = yf.Ticker(symbol)
        return ticker_to_stock(symbol, ticker)
    except Exception as e:
        logger = get_logger()
        logger.error('can not get stock data', stock=symbol, exception=e)
        return None

# Fetch multiple stocks
# Limit rate: ~ 200 stocks per time
def fetch_stocks(symbols: list[str], batch: int = 200) -> list[Stock] | None:
    try:
        i = 0
        result = []
        if len(symbols) <= 0:
            raise Exception('symbols is empty')
        while i < len(symbols):
            fetched_stocks = symbols[i: i + batch]
            tickers = yfinance.Tickers(' '.join(fetched_stocks))
            if not tickers.tickers:
                raise Exception('no ticker found')
            for symbol in fetched_stocks:
                if tickers.tickers[symbol]:
                    result.append(ticker_to_stock(symbol, tickers.tickers[symbol]))
            i += batch
        return result
    except Exception as e:
        logger = get_logger()
        logger.error('can not get stocks data', exception=e)
        return None