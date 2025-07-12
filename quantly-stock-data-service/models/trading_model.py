import uuid
from sqlmodel import SQLModel, Field, Relationship
from typing import Optional
from uuid import UUID
from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from .cond_model_index import CondModelIndex

class TradingModel(SQLModel, table=True):
    __tablename__ = "trading_models"

    trading_model_id: str = Field(primary_key=True, index=True, default_factory=lambda : str(uuid.uuid4()))
    name: Optional[str] = Field(default=None, max_length=50)
    description: Optional[str] = Field(default=None)
    status: Optional[str] = Field(default=None, max_length=36)
    account_id: str = Field(foreign_key="accounts.account_id")
    action: Optional[str] = Field(default=None, max_length=10)
    cond_model_indices: list["CondModelIndex"] = Relationship(back_populates="trading_model")
