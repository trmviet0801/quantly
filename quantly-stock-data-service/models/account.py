import uuid

from sqlmodel import SQLModel, Field
from typing import Optional
from datetime import datetime

class Account(SQLModel, table=True):
    __tablename__ = "accounts"

    account_id: str = Field(primary_key=True, max_length=36)
    account_number: Optional[str] = Field(default=None, max_length=36)
    status: Optional[str] = Field(default=None, max_length=10)
    currency: Optional[str] = Field(default=None, max_length=10)
    last_equity: Optional[str] = Field(default=None, max_length=20)
    create_at: Optional[datetime] = None
    account_type: Optional[str] = Field(default=None, max_length=20)
    trading_type: Optional[str] = Field(default=None, max_length=20)
    user_id: Optional[str] = Field(default=None, foreign_key="users.user_id", max_length=36)
