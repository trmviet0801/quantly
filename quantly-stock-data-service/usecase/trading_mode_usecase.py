from yfinance import Ticker

from repo.mysql_repo import get_all_records
from models.condition import Condition
from usecase.indicator_usecase import indicator_from_stock, if_stock_matched_condition
from repo.redis_repo import redis_get_object, redis_post_object
from db.redis_conn import get_redis_instance
from utils.trading_model_utils import get_condition_key
from models.condition_match_group import ConditionMatchGroup
from logger.structlog_config import get_logger

def evaluate_conditions_and_cache_matches(stock: str, ticker: Ticker):
    """
    Evaluate all existed conditions with specific stock.
    Args:
        stock: stock symbol
        ticker: ticker from yFinance
    Returns:
        - Matched -> store in Redis ->
        - Not matched -> if exist on redis -> remove
    """
    conditions = get_all_records(Condition)
    redis_conn = get_redis_instance()
    logger = get_logger()

    if ticker is None:
        logger.error('[evaluate_conditions_and_cache_matches] No ticker provided', stock={stock})
        return

    for condition in conditions:
        key = get_condition_key(condition)
        changed = False
        matched = False
        current_matching_stocks = redis_get_object(key, ConditionMatchGroup,
                                                   redis_conn)
        if if_stock_matched_condition(condition, indicator_from_stock(stock, condition, ticker), ticker.fast_info['lastPrice']):
            if current_matching_stocks is not None:
                if stock not in current_matching_stocks.matching_stocks:
                    current_matching_stocks.matching_stocks = set(current_matching_stocks.matching_stocks)
                    current_matching_stocks.matching_stocks.add(stock)
                    matched = True
                    changed = True
            else:
                current_matching_stocks = ConditionMatchGroup(key, {stock})
                changed = True
                matched = True
        else:
            if current_matching_stocks is not None and stock in current_matching_stocks.matching_stocks:
                current_matching_stocks.matching_stocks.remove(stock)
                changed = True

        if changed:
            if redis_post_object(current_matching_stocks, key, redis_conn):
                logger.info(
                    f"[Redis Updated] Stock: {stock} | Condition: {key} | Action: {'added' if matched else 'removed'}"
                )
            else:
                logger.error(
                    f"[Redis Update Failed]: {stock} | Condition: {key} | Action: {'failed'}"
                )








