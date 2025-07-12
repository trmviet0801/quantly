import uuid
from sqlmodel import SQLModel, Field, Relationship
from typing import Optional
from .trading_model import TradingModel
from .condition import Condition

class CondModelIndex(SQLModel, table=True):
    __tablename__ = 'cond_model_index'

    cond_model_index_id: str = Field(primary_key=True, default_factory=lambda : str(uuid.uuid4()))
    condition_id: str = Field(foreign_key="conditions.condition_id")
    trading_model_id: str = Field(foreign_key="trading_models.trading_model_id")

    trading_model: Optional[TradingModel] = Relationship(back_populates="cond_model_indices")
    condition: Optional[Condition] = Relationship(back_populates="cond_model_indices")