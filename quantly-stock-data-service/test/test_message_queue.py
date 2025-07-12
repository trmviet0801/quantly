from message_queue.rabbitmq import post_message
from models.trading_action import TradingAction

def test_post_message():
    trading_action = TradingAction('buy', 'NVDA', 200)
    is_ok = post_message(trading_action)
    assert is_ok is True