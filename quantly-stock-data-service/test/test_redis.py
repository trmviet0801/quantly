import pytest
from db.redis_conn import get_redis_instance
from repo.redis_repo import  redis_post_object, redis_get_object
from usecase.fetch_stock import fetch_stock
from models.stock import Stock
import redis.commands.json.path as Path

def test_post_data_redis():
    redis_conn = get_redis_instance()
    redis_conn.set(name='test', value='new value')
    assert redis_conn.get('test') == 'new value'

def test_post_stock_redis():
    redis_conn = get_redis_instance()
    stock = Stock(symbol='test')
    redis_post_object(stock, 'stock:test', redis_conn)
    assert redis_conn.json().get('stock:test', Path.Path('.symbol')) == 'test'

def test_get_stock_redis():
    redis_conn = get_redis_instance()
    assert redis_get_object('stock:test', Stock, redis_conn).symbol == 'test'
    assert redis_get_object('stock:none', Stock, redis_conn) is None
