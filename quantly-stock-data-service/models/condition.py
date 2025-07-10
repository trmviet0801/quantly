import uuid

from sqlalchemy import Column, Numeric
from sqlmodel import SQLModel, Field, Relationship
from typing import Optional
from uuid import UUID

# from models.trading_model import TradingModel


class Condition(SQLModel, table=True):
    __tablename__ = "conditions"

    condition_id: UUID = Field(primary_key=True, index=True, default=uuid.uuid4)
    trading_model_id: UUID = Field(foreign_key="trading_models.trading_model_id")

    indicator: str = Field(max_length=50)
    operation: str = Field(max_length=10)
    value: float = None
    period: int

    period_k: Optional[float] = None
    period_d: Optional[float] = None
    std_dev: Optional[float] = None
    band: Optional[str] = Field(default=None, max_length=20)

    # Optional relationship back to TradingModel
    # trading_model: Optional["TradingModel"] = Relationship(back_populates="conditions")
