import pandas

def get_redis_key_of_stock(symbol: str) -> str:
    return "stock:" + symbol

# Returns all S&P500 stock's symbols
def get_stock_symbols() -> list[str]:
    url = 'https://en.wikipedia.org/wiki/List_of_S%26P_500_companies'
    table = pandas.read_html(url, header=0)[0]
    return table["Symbol"].tolist()