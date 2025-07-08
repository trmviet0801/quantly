from dataclasses import dataclass
from typing import Optional

from sqlmodel import Field, SQLModel

# Redis key: stock:[symbol]
@dataclass
class Stock(SQLModel, table=True, namespace="stocks"):
    shortName: Optional[str] = None
    longName: Optional[str] = None
    currency: Optional[str] = None
    exchange: Optional[str] = None
    quoteType: Optional[str] = None
    marketCap: Optional[int] = None
    previousClose: Optional[float] = None
    open: Optional[float] = None
    dayHigh: Optional[float] = None
    dayLow: Optional[float] = None
    fiftyTwoWeekHigh: Optional[float] = None
    fiftyTwoWeekLow: Optional[float] = None
    trailingPE: Optional[float] = None
    forwardPE: Optional[float] = None
    priceToBook: Optional[float] = None
    beta: Optional[float] = None
    volume: Optional[int] = None
    averageVolume: Optional[int] = None
    dividendYield: Optional[float] = None
    earningsTimestamp: Optional[int] = None
    website: Optional[str] = None
    industry: Optional[str] = None
    sector: Optional[str] = None
    currenPrice: Optional[float] = None
    symbol: str = Field(primary_key=True)