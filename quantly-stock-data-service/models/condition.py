from sqlmodel import SQLModel, Field, Relationship
from typing import Optional
from uuid import UUID

from models.trading_model import TradingModel


class Condition(SQLModel, table=True):
    __tablename__ = "conditions"

    condition_id: UUID = Field(primary_key=True, index=True)
    trading_model_id: UUID = Field(foreign_key="trading_models.trading_model_id")

    indicator: str = Field(max_length=50)
    operation: str = Field(max_length=10)
    value: float = Field(sa_column_kwargs={"precision": 10, "scale": 4})
    period: int

    period_k: Optional[float] = Field(default=None, sa_column_kwargs={"precision": 10, "scale": 4})
    period_d: Optional[float] = Field(default=None, sa_column_kwargs={"precision": 10, "scale": 4})
    std_dev: Optional[float] = Field(default=None, sa_column_kwargs={"precision": 10, "scale": 4})
    band: Optional[str] = Field(default=None, max_length=20)

    # Optional relationship back to TradingModel
    trading_model: Optional["TradingModel"] = Relationship(back_populates="conditions")
