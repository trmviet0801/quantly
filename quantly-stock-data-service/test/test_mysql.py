from repo.mysql_repo import get_data_by_id, get_all_records
from models.stock import Stock
from models.cond_model_index import CondModelIndex

def test_get_stock_from_mysql():
    stock = get_data_by_id(Stock, 'NVDA')
    print(stock)
    assert stock.symbol == 'NVDA'

def test_get_all_data_from_mysql():
    data = get_all_records(CondModelIndex)
    assert len(data) > 1