import time
from dataclasses import dataclass
from typing import Optional
from sqlmodel import SQLModel, Field
from datetime import datetime

# Redis key: stock:[symbol]
@dataclass
class Stock(SQLModel, table=True):
    __tablename__ = "stocks"

    symbol: str = Field(primary_key=True, max_length=10)

    short_name: Optional[str] = Field(default=None, max_length=50)
    long_name: Optional[str] = Field(default=None, max_length=100)
    sector: Optional[str] = Field(default=None, max_length=50)
    industry: Optional[str] = Field(default=None, max_length=50)

    market_cap: Optional[int] = None
    current_price: Optional[float] = None
    trailing_pe: Optional[float] = None
    forward_pe: Optional[float] = None
    price_to_book: Optional[float] = None
    dividend_yield: Optional[float] = None
    beta: Optional[float] = None
    fifty_two_week_high: Optional[float] = None
    fifty_two_week_low: Optional[float] = None

    average_volume: Optional[int] = None
    last_close_price: Optional[float] = None
    last_volume: Optional[int] = None

    currency: Optional[str] = Field(default=None, max_length=24)
    exchange: Optional[str] = Field(default=None, max_length=24)

    updated_at: Optional[datetime] = Field(default=time.time())
