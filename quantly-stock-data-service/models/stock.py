from dataclasses import dataclass
from typing import Optional

from sqlmodel import Field, SQLModel

@dataclass
class Stock(SQLModel, table=True, namespace="stocks"):
    shortName: Optional[str]
    longName: Optional[str]
    currency: Optional[str]
    exchange: Optional[str]
    quoteType: Optional[str]
    marketCap: Optional[int]
    previousClose: Optional[float]
    open: Optional[float]
    dayHigh: Optional[float]
    dayLow: Optional[float]
    fiftyTwoWeekHigh: Optional[float]
    fiftyTwoWeekLow: Optional[float]
    trailingPE: Optional[float]
    forwardPE: Optional[float]
    priceToBook: Optional[float]
    beta: Optional[float]
    volume: Optional[int]
    averageVolume: Optional[int]
    dividendYield: Optional[float]
    earningsTimestamp: Optional[int]
    website: Optional[str]
    industry: Optional[str]
    sector: Optional[str]
    symbol: str = Field(primary_key=True)