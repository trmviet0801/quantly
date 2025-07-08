from models.stock import Stock # adjust import as needed
import yfinance as yf

def ticker_to_stock(symbol: str, ticker: yf.Ticker) -> Stock | None:
    info = ticker.info

    # Validate equity type
    if not info or info.get("quoteType") != "EQUITY":
        return None

    return Stock(
        symbol=symbol.upper(),
        short_name=info.get("shortName"),
        long_name=info.get("longName"),
        sector=info.get("sector"),
        industry=info.get("industry"),
        market_cap=info.get("marketCap"),
        current_price=info.get("currentPrice"),
        trailing_pe=info.get("trailingPE"),
        forward_pe=info.get("forwardPE"),
        price_to_book=info.get("priceToBook"),
        dividend_yield=info.get("dividendYield"),
        beta=info.get("beta"),
        fifty_two_week_high=info.get("fiftyTwoWeekHigh"),
        fifty_two_week_low=info.get("fiftyTwoWeekLow"),
        average_volume=info.get("averageVolume"),
        last_close_price=info.get("previousClose"),
        last_volume=info.get("volume"),
        currency=info.get("currency"),
        exchange=info.get("exchange"),
    )
