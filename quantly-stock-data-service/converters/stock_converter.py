import yfinance as yf
from models.stock import Stock

def ticker_to_stock(symbol: str, ticker: yf.Ticker) -> Stock | None:
    info = ticker.info
    if not info or not info.get("quoteType") or info.get("quoteType") != "EQUITY":
        return None
    return Stock(
        shortName=info.get("shortName"),
        longName=info.get("longName"),
        currency=info.get("currency"),
        exchange=info.get("exchange"),
        quoteType=info.get("quoteType"),
        marketCap=info.get("marketCap"),
        previousClose=info.get("previousClose"),
        open=info.get("open"),
        dayHigh=info.get("dayHigh"),
        dayLow=info.get("dayLow"),
        fiftyTwoWeekHigh=info.get("fiftyTwoWeekHigh"),
        fiftyTwoWeekLow=info.get("fiftyTwoWeekLow"),
        trailingPE=info.get("trailingPE"),
        forwardPE=info.get("forwardPE"),
        priceToBook=info.get("priceToBook"),
        beta=info.get("beta"),
        volume=info.get("volume"),
        averageVolume=info.get("averageVolume"),
        dividendYield=info.get("dividendYield"),
        earningsTimestamp=info.get("earningsTimestamp"),
        website=info.get("website"),
        industry=info.get("industry"),
        sector=info.get("sector"),
        symbol=symbol.upper(),
        currenPrice=info.get("currentPrice"),
    )
