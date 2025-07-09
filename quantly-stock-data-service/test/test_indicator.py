import usecase.indicator_usecase as cal
import yfinance as yf

def test_sma():
    ticker = yf.Ticker('NVDA')
    print(cal.cal_sma(stock='NVDA', ticker=ticker, fetch_range='1d', sma_window=3, interval='1m'))

def test_ema():
    ticker = yf.Ticker('NVDA')
    print(cal.cal_ema(stock='NVDA', ticker=ticker, fetch_range='1d', ema_span=3, interval='1m'))

def test_rsi():
    ticker = yf.Ticker('NVDA')
    print(cal.cal_rsi(stock='NVDA', ticker=ticker, fetch_range='1d', rsi_period=3, interval='1m'))

def test_bollinger_bands():
    ticker = yf.Ticker('NVDA')
    result = cal.cal_bollinger_bands(
        stock='NVDA',
        ticker=ticker,
        fetch_range='5d',
        window=20,
        interval='1m'
    )
    if result:
        upper, middle, lower = result
        print("Upper Band:\n", upper.tail())
        print("Middle Band (SMA):\n", middle.tail())
        print("Lower Band:\n", lower.tail())
    else:
        print("Bollinger Bands calculation failed.")

def test_atr():
    ticker = yf.Ticker('NVDA')
    result = cal.cal_atr(
        stock='NVDA',
        ticker=ticker,
        fetch_range='1mo',
        atr_period=14,
        interval='1d'
    )
    if result is not None:
        print("ATR:\n", result.tail())
    else:
        print("ATR calculation failed.")

def test_stochastic_oscillator():
    ticker = yf.Ticker('NVDA')
    result = cal.cal_stochastic_oscillator(
        stock='NVDA',
        ticker=ticker,
        fetch_range='1mo',
        period_k=14,
        period_d=3,
        interval='1d'
    )
    if result is not None:
        percent_k, percent_d = result
        print("Stochastic %K:\n", percent_k.tail())
        print("Stochastic %D:\n", percent_d.tail())
    else:
        print("Stochastic Oscillator calculation failed.")