import pytest
from db.redis_conn import get_redis_instance

def test_post_data_redis():
    redis_conn = get_redis_instance()
    redis_conn.set(name='test', value='new value')
    assert redis_conn.get('test') == 'new value'