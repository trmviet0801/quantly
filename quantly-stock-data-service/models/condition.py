import uuid

from sqlmodel import SQLModel, Field, Relationship
from typing import Optional, TYPE_CHECKING

if TYPE_CHECKING:
    from .cond_model_index import CondModelIndex

class Condition(SQLModel, table=True):
    __tablename__ = "conditions"

    condition_id: str = Field(primary_key=True, index=True, default_factory=lambda : str(uuid.uuid4()))

    indicator: str = Field(max_length=50)
    operation: str = Field(max_length=10)
    value: Optional[float] = None
    period: int

    period_k: Optional[float] = None
    period_d: Optional[float] = None
    std_dev: Optional[float] = None
    band: Optional[str] = Field(default=None, max_length=20)

    cond_model_indices: list["CondModelIndex"] = Relationship(back_populates="condition")