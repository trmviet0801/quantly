from repo.mysql_repo import get_data_by_id
from models.stock import Stock

def test_get_stock_from_mysql():
    stock = get_data_by_id(Stock, 'NVDA')
    print(stock)
    assert stock.symbol == 'NVDA'