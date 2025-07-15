from pandas import Series


def average_close_price(data: Series) -> float:
    num = len(data)
    if num == 0:
        return 0
    sum = 0
    for price in data:
        sum += price
    return sum / num

# Helper to resolve appropriate 'period' (fetch_range) for yfinance
def resolve_yfinance_params(
    period_length: int,
    interval: str = '1d',
    multiplier: int = 3
) -> str | None:
    """
    Determine an appropriate 'period' (fetch_range) for yfinance.history(),
    given the required indicator window (e.g., SMA(20)) and interval.

    Args:
        period_length (int): Number of data points needed (e.g., 20 for SMA(20))
        interval (str): Granularity (e.g., '1d', '5m')
        multiplier (int): How many times more data to fetch than needed

    Returns:
        str: A valid 'period' string for yfinance
    """
    if period_length is None:
        return None
    total_days = period_length * multiplier

    # Intraday intervals are limited to short history
    if interval in {"1m", "2m", "5m", "15m", "30m", "60m", "90m"}:
        if total_days <= 1:
            return "1d"
        elif total_days <= 5:
            return "5d"
        else:
            return "7d"  # max safe window for intraday
    else:
        if total_days <= 5:
            return "5d"
        elif total_days <= 30:
            return "1mo"
        elif total_days <= 90:
            return "3mo"
        elif total_days <= 180:
            return "6mo"
        elif total_days <= 365:
            return "1y"
        elif total_days <= 730:
            return "2y"
        elif total_days <= 1825:
            return "5y"
        elif total_days <= 3650:
            return "10y"
        else:
            return "max"