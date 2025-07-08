import redis
from typing import Optional
from dotenv import load_dotenv
import os

_redis_instance: Optional[redis.Redis] = None
load_dotenv()

def get_redis_instance():
    global _redis_instance
    if _redis_instance is None:
        _redis_instance = redis.Redis(host=os.getenv('REDIS_HOST'),
                                      port=int(os.getenv('REDIS_PORT')),
                                      decode_responses=True,
                                      username=os.getenv('REDIS_USERNAME'),
                                      password=os.getenv('REDIS_PASSWORD'))
    return _redis_instance