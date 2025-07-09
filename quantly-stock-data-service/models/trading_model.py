import uuid
from sqlmodel import SQLModel, Field, Relationship
from typing import Optional
from uuid import UUID

from models.condition import Condition


class TradingModel(SQLModel, table=True):
    __tablename__ = "trading_models"

    trading_model_id: UUID = Field(primary_key=True, index=True, default=uuid.uuid4)
    name: Optional[str] = Field(default=None, max_length=50)
    description: Optional[str] = Field(default=None)
    status: Optional[str] = Field(default=None, max_length=36)
    account_id: UUID = Field(foreign_key="accounts.account_id")
    action: Optional[str] = Field(default=None, max_length=10)

    conditions: list[Condition] = Relationship(back_populates="trading_model")
