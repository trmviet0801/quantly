import yfinance

from usecase.trading_mode_usecase import evaluate_conditions_and_cache_matches

def test_evaluate_conditions_and_cache_matches():
    evaluate_conditions_and_cache_matches('AAPL', yfinance.Ticker('AAPL'))
    evaluate_conditions_and_cache_matches('MSFT', yfinance.Ticker('MSFT'))
    evaluate_conditions_and_cache_matches('PENN', yfinance.Ticker('PENN'))