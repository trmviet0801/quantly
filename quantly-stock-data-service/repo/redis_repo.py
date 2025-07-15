import json
from typing import TypeVar, Type
from dataclasses import asdict, is_dataclass
from logger.structlog_config import get_logger

import redis

T = TypeVar('T')

def redis_post_object(data: T, key: str, redis_conn: redis.Redis) -> bool:
    try:
        if not is_dataclass(data):
            raise TypeError('Provided data is not a dataclass')

        if redis_conn is None:
            raise ValueError("Redis connection is None")

        data_dict = asdict(data)

        for k, v in data_dict.items():
            if isinstance(v, set):
                data_dict[k] = list(v)

        redis_conn.json().set(key, "$", data_dict)
        return True
    except Exception as e:
        logger = get_logger()
        logger.error("Redis post failed", reason=str(e), key=key)
        return False

# Handle response from Redis with data type = [list, dict]
def redis_get_object(key: str, cls: Type[T], redis_conn: redis.Redis) -> T | None:
    try:
        if redis_conn is None:
            raise ValueError("Redis connection is None")

        json_data = redis_conn.json().get(key)
        if isinstance(json_data, list) and len(json_data) > 0:
            obj_data = json_data[0]
            if isinstance(obj_data, dict):
                return cls(**obj_data)
            else:
                raise TypeError('Could not get object from redis')
        elif isinstance(json_data, dict):
            return cls(**json_data)
        else:
            logger = get_logger()
            logger.error("Redis get object failed", key=key, data_type=type(json_data), size=len(json_data))
            return None
    except Exception as e:
        logger = get_logger()
        logger.error("Redis get failed", reason=str(e), key=key)
        return None
