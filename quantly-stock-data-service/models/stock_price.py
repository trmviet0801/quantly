from sqlmodel import SQLModel, Field
from typing import Optional
from datetime import datetime

class StockPrice(SQLModel, table=True):
    __tablename__ = "stock_prices"

    stock_price_id: str = Field(primary_key=True, max_length=36)
    symbol: str = Field(foreign_key="stocks.symbol", max_length=10)
    date: datetime

    open: Optional[float] = None
    close: Optional[float] = None
    high: Optional[float] = None
    low: Optional[float] = None
    volume: Optional[int] = None
