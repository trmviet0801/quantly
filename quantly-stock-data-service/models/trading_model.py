from sqlmodel import SQLModel, Field
from typing import Optional


class TradingModel(SQLModel, table=True):
    __tablename__ = "trading_models"

    trading_model_id: str = Field(primary_key=True, max_length=36)
    name: Optional[str] = Field(default=None, max_length=50)
    description: Optional[str] = None
    status: Optional[str] = Field(default=None, max_length=36)
    account_id: Optional[str] = Field(default=None, foreign_key="accounts.account_id", max_length=36)

    formula: Optional[str] = None
    entry_condition: Optional[str] = None
    exit_condition: Optional[str] = None
