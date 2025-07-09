import uuid
from sqlmodel import SQLModel, Field
from typing import Optional
from datetime import datetime

class Position(SQLModel, table=True):
    __tablename__ = "positions"

    position_id: str = Field(primary_key=True, max_length=36, default=uuid.uuid4)
    symbol: str = Field(foreign_key="stocks.symbol", max_length=10)
    account_id: str = Field(foreign_key="accounts.account_id", max_length=36)

    quantity: Optional[int] = None
    created_at: Optional[datetime] = None
    price: Optional[float] = None

    position_type: Optional[str] = Field(default=None, alias="type", max_length=10)
    status: Optional[str] = Field(default=None, max_length=10)
