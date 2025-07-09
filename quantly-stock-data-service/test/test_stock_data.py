# from types import NoneType
#
# from models.stock import Stock
# from usecase.fetch_stock import fetch_stock, fetch_stocks
# from utils.stock_utils import get_stock_symbols, get_redis_key_of_stock
# from repo.redis_repo import redis_post_object
# from db.redis_conn import get_redis_instance
#
# def test_get_single_stock():
#     nvda_stock = fetch_stock('NVDA')
#     error_stock = fetch_stock('something')
#     assert nvda_stock.symbol == 'NVDA'
#     assert type(error_stock) is NoneType
#
# def test_get_stock_symbols():
#     symbols = get_stock_symbols()
#     assert len(symbols) > 0
#     assert type(symbols[0]) is str
#     assert symbols[0] == 'MMM'
#
# def test_fetch_stocks():
#     symbols = get_stock_symbols()
#     stocks = fetch_stocks(symbols)
#     for stock in stocks:
#         if stock:
#             redis_post_object(stock, get_redis_key_of_stock(stock.symbol), get_redis_instance())
#     assert len(stocks) > 0
#     assert type(stocks[0]) is Stock