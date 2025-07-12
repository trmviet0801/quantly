import json

import pika
from typing import TypeVar
from logger.structlog_config import get_logger
from dataclasses import asdict

T = TypeVar('T')

logger = get_logger()

# Post message object to message queue
# Default queue = 'trading_action_queue'
# Default host = localhost
def post_message(data: T, queue_name: str = 'trading_action_queue', host: str = 'localhost') -> bool:
    try:
        with pika.BlockingConnection(pika.ConnectionParameters(host)) as conn:
            channel = conn.channel()
            channel.queue_declare(queue=queue_name, durable=True)
            channel.basic_publish(
                exchange='',
                routing_key=queue_name,
                body=json.dumps(asdict(data)),
                properties=pika.BasicProperties()
            )
            return True
    except Exception as e:
        logger.error("Insert data to RabbitMQ DB failed", queue=queue_name, error=str(e))
        return False