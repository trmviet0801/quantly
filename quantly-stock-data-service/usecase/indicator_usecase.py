import pandas as pd
import yfinance as yf
from pandas import Series
from logger.structlog_config import get_logger

# Interval     -> Recommended Minimum Period
# -------------------------------------------
# '1m'         -> '1d' to '7d'
# '2m'-'15m'   -> '5d+'
# '1d'         -> '1mo+'
# '1wk'        -> '3mo+'
# '1mo'        -> '6mo+'

# Simple moving average.
# Average of closing price
def cal_sma(stock: str, ticker: yf.Ticker, fetch_range: str, sma_window: int, interval: str) -> Series | None:
    try:
        history = ticker.history(period=fetch_range, interval=interval)
        if history.empty:
            raise Exception('No data found')

        close_prices = history['Close']
        return close_prices.rolling(window=sma_window).mean()
    except Exception as e:
        logger = get_logger()
        logger.error('can not calculate SMA', stock=stock, fetch_range=fetch_range, sma_window=sma_window, interval=interval, error=e)
        return None

# EMA - Exponential moving average
def cal_ema(stock: str, ticker: yf.Ticker, fetch_range: str, ema_span: int, interval: str) -> Series | None:
    try:
        history = ticker.history(period=fetch_range, interval=interval)
        if history.empty:
            raise Exception('No data found')

        close_prices = history['Close']
        return close_prices.ewm(span=ema_span, adjust=False).mean()
    except Exception as e:
        logger = get_logger()
        logger.error('can not calculate EMA', stock=stock, fetch_range=fetch_range, interval=interval, error=e)
        return None

# RSI - Relative Strength Index (Wilder’s method)
def cal_rsi(stock: str, ticker: yf.Ticker, fetch_range: str, rsi_period: int, interval: str) -> Series | None:
    try:
        history = ticker.history(period=fetch_range, interval=interval)
        if history.empty:
            raise Exception("No data found")

        close_prices = history["Close"]
        delta = close_prices.diff()

        gain = delta.where(delta > 0, 0.0)
        loss = -delta.where(delta < 0, 0.0)

        avg_gain = gain.ewm(alpha=1/rsi_period, adjust=False).mean()
        avg_loss = loss.ewm(alpha=1/rsi_period, adjust=False).mean()

        rs = avg_gain / avg_loss
        rsi = 100 - (100 / (1 + rs))

        return rsi

    except Exception as e:
        logger = get_logger()
        logger.error(
            f"Cannot calculate RSI for {stock} | range={fetch_range} | interval={interval} | period={rsi_period} | error: {e}"
        )
        return None

# Bollinger Bands - Upper, Middle, Lower
def cal_bollinger_bands(stock: str, ticker: yf.Ticker, fetch_range: str, window: int, interval: str, num_std: float = 2.0) -> tuple[Series, Series, Series] | None:
    try:
        history = ticker.history(period=fetch_range, interval=interval)
        if history.empty:
            raise Exception("No data found")

        close_prices = history["Close"]
        sma = close_prices.rolling(window=window).mean()
        std = close_prices.rolling(window=window).std()

        upper_band = sma + (num_std * std)
        lower_band = sma - (num_std * std)

        return upper_band, sma, lower_band

    except Exception as e:
        logger = get_logger()
        logger.error(
            f"Cannot calculate Bollinger Bands for {stock} | range={fetch_range} | interval={interval} | window={window} | error: {e}"
        )
        return None

# ATR - Average True Range (Wilder’s method)
def cal_atr(stock: str, ticker: yf.Ticker, fetch_range: str, atr_period: int, interval: str) -> Series | None:
    try:
        history = ticker.history(period=fetch_range, interval=interval)
        if history.empty or len(history) < atr_period + 1:
            raise Exception("Not enough data to calculate ATR")

        high = history["High"]
        low = history["Low"]
        close = history["Close"]

        prev_close = close.shift(1)

        tr = pd.concat([
            (high - low),
            (high - prev_close).abs(),
            (low - prev_close).abs()
        ], axis=1).max(axis=1)

        atr = tr.ewm(alpha=1/atr_period, adjust=False).mean()

        return atr

    except Exception as e:
        logger = get_logger()
        logger.error(
            f"Cannot calculate ATR for {stock} | range={fetch_range} | interval={interval} | period={atr_period} | error: {e}"
        )
        return None

# Stochastic Oscillator - %K and %D
def cal_stochastic_oscillator(
    stock: str,
    ticker: yf.Ticker,
    fetch_range: str,
    period_k: int,
    period_d: int,
    interval: str
) -> tuple[pd.Series, pd.Series] | None:
    try:
        history = ticker.history(period=fetch_range, interval=interval)
        if history.empty or len(history) < period_k + period_d:
            raise Exception("Not enough data to calculate Stochastic Oscillator")

        high = history["High"]
        low = history["Low"]
        close = history["Close"]

        lowest_low = low.rolling(window=period_k).min()
        highest_high = high.rolling(window=period_k).max()

        percent_k = ((close - lowest_low) / (highest_high - lowest_low)) * 100
        percent_d = percent_k.rolling(window=period_d).mean()

        return percent_k, percent_d

    except Exception as e:
        logger = get_logger()
        logger.error(
            f"Cannot calculate Stochastic Oscillator for {stock} | range={fetch_range} | interval={interval} | %K={period_k} | %D={period_d} | error: {e}"
        )
        return None

# Momentum - Price difference over a lookback period
def cal_momentum(stock: str, ticker: yf.Ticker, fetch_range: str, momentum_period: int, interval: str) -> pd.Series | None:
    try:
        history = ticker.history(period=fetch_range, interval=interval)
        if history.empty or len(history) <= momentum_period:
            raise Exception("Not enough data to calculate Momentum")

        close_prices = history["Close"]
        momentum = close_prices - close_prices.shift(momentum_period)

        return momentum

    except Exception as e:
        logger = get_logger()
        logger.error(
            f"Cannot calculate Momentum for {stock} | range={fetch_range} | interval={interval} | period={momentum_period} | error: {e}"
        )
        return None
