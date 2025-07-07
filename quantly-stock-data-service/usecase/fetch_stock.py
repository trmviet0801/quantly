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