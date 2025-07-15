import pandas as pd
import yfinance as yf
from pandas import Series
from yfinance import Ticker

from logger.structlog_config import get_logger
from utils.indicator_utils import resolve_yfinance_params
from models.condition import Condition

# ==============================
# Valid 'period' (fetch_range) values:
# ==============================
# "1d"    -> Last 1 day
# "5d"    -> Last 5 days
# "7d"    -> Last 7 days (not official, but works)
# "1mo"   -> Last 1 month
# "3mo"   -> Last 3 months
# "6mo"   -> Last 6 months
# "1y"    -> Last 1 year
# "2y"    -> Last 2 years
# "5y"    -> Last 5 years
# "10y"   -> Last 10 years
# "ytd"   -> From Jan 1 of current year to today
# "max"   -> All available historical data

# ==============================
# Valid 'interval' values:
# ==============================
# "1m"    -> 1-minute candles (only for last 7 days, use period="1d" or "5d")
# "2m"    -> 2-minute candles (up to 60 days)
# "5m"    -> 5-minute candles (up to 60 days)
# "15m"   -> 15-minute candles (up to 60 days)
# "30m"   -> 30-minute candles (up to 60 days)
# "60m"   -> 60-minute candles (up to 60 days)
# "90m"   -> 90-minute candles (up to 60 days)
# "1d"    -> Daily candles
# "5d"    -> 5-day candles
# "1wk"   -> Weekly candles
# "1mo"   -> Monthly candles
# "3mo"   -> Quarterly candles

# ==============================
# Compatibility Notes:
# ==============================
# - Intraday intervals ("1m", "5m", etc.) only work with short periods ("1d", "5d", "7d")
# - Longer periods ("6mo", "1y", etc.) should use "1d", "1wk", "1mo", etc.
# - If you pass an invalid combination, the result will be empty
# - Do NOT mix 'period' with 'start'/'end' parameters

def if_stock_matched_condition(condition: Condition, indicator_value, current_price: float) -> bool:
    """
    Check if condition matches indicator condition.
    Args:
        condition: models.Condition
        indicator_value: indicator value of specific stock. Can be [float, tuple{Series, Series], tuple{Series, Series, Series}]
        current_price: current price of stock

    Returns: boolean (return False if error occurs)

    """
    if indicator_value is None:
        return False

    indicator = condition.indicator
    if indicator == 'SMA' or indicator == 'EMA' or indicator == 'RSI' or indicator == 'ATR' or indicator == 'Momentum':
        return {
            '<': indicator_value < condition.value,
            '>': indicator_value > condition.value,
            '<=': indicator_value <= condition.value,
            '>=': indicator_value >= condition.value,
            '=': indicator_value == condition.value,
        }.get(condition.operation, False)
    elif indicator == 'BollingerBands':
        band_to_index = {
            'upper': 0,
            'lower': 2,
        }
        index = band_to_index.get(condition.band, 1)
        return {
            '<': indicator_value[index].iloc[-1] < current_price,
            '>': indicator_value[index].iloc[-1] > current_price,
            '<=': indicator_value[index].iloc[-1] <= current_price,
            '>=': indicator_value[index].iloc[-1] >= current_price,
            '=': indicator_value[index].iloc[-1] == current_price,
        }.get(condition.operation, False)
    elif indicator == 'StochasticOscillator':
        return {
            '<': indicator_value[1] < condition.value,
            '>': indicator_value[1] > condition.value,
            '<=': indicator_value[1] <= condition.value,
            '>=': indicator_value[1] >= condition.value,
            '=': indicator_value[1] == condition.value,
        }.get(condition.operation, False)
    return False

def indicator_from_stock(stock: str, cond: Condition, ticker: yf.Ticker):
    """
        Calculating indicators for specific stock
    Args:
        stock: symbol of stock to calculate indicators for
        ticker: ticker from yFinance
        cond: condition to calculate indicators for

    Returns:
        - float for scalar indicators
        - tuple[pd.Series, pd.Series, pd.Series] for BollingerBands
        - tuple[pd.Series, pd.Series] for StochasticOscillator
        - 0 on failure

    """
    indicator = cond.indicator
    fetch_range = resolve_yfinance_params(period_length=cond.period)

    if fetch_range is None:
        return None

    if indicator == 'SMA':
        print(type(cal_sma(stock, ticker, fetch_range, cond.period)))
        return cal_sma(stock, ticker, fetch_range, cond.period).iloc[-1]
    elif indicator == 'EMA':
        return cal_ema(stock, ticker, fetch_range, cond.period).iloc[-1]
    elif indicator == 'RSI':
        return cal_rsi(stock, ticker, fetch_range, cond.period).iloc[-1]
    elif indicator == 'BollingerBands':
        return cal_bollinger_bands(stock, ticker, fetch_range, cond.period)
    elif indicator == 'ATR':
        return cal_atr(stock, ticker, fetch_range, cond.period).iloc[-1]
    elif indicator == 'StochasticOscillator':
        return cal_stochastic_oscillator(stock, ticker, fetch_range, int(cond.period_k), int(cond.period_d))
    elif indicator == 'Momentum':
        return cal_momentum(stock, ticker, fetch_range, cond.period).iloc[-1]
    return 0

# Simple moving average.
# Average of closing price
def cal_sma(stock: str, ticker: yf.Ticker, fetch_range: str, sma_window: int, interval: str = '1d') -> Series | None:
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
def cal_ema(stock: str, ticker: yf.Ticker, fetch_range: str, ema_span: int, interval: str = '1d') -> Series | None:
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
def cal_rsi(stock: str, ticker: yf.Ticker, fetch_range: str, rsi_period: int, interval: str = '1d') -> Series | None:
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
def cal_bollinger_bands(stock: str, ticker: yf.Ticker, fetch_range: str, window: int, interval: str = '1d', num_std: float = 2.0) -> tuple[Series, Series, Series] | None:
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
def cal_atr(stock: str, ticker: yf.Ticker, fetch_range: str, atr_period: int, interval: str = '1d') -> Series | None:
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
    interval: str = '1d'
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
def cal_momentum(stock: str, ticker: yf.Ticker, fetch_range: str, momentum_period: int, interval: str = '1d') -> pd.Series | None:
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

