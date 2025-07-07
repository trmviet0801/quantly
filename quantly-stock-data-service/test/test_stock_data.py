from types import NoneType

from usecase.fetch_stock import fetch_stock
from logger.structlog_config import get_logger


def test_get_single_stock():
    nvda_stock = fetch_stock('NVDA')
    error_stock = fetch_stock('something')

    logger = get_logger()
    logger.info(type(nvda_stock))

    assert nvda_stock.symbol == 'NVDA'
    assert type(error_stock) is NoneType