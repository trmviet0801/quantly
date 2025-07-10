from typing import TypeVar, Type, cast

from sqlmodel import select, SQLModel

from db.mysql_config import get_session
from logger.structlog_config import get_logger


T = TypeVar("T", bound=SQLModel)

# overwrite if duplicated
def insert_data(data: T) -> T | None:
    try:
        with get_session() as session:
            merged = session.merge(data)
            session.commit()
            session.refresh(merged)
            return merged
    except Exception as e:
        logger = get_logger()
        logger.error("Insert data to Mysql DB failed", error=str(e))
        return None

def get_data_by_id(cls: Type[T], data_id: str) -> T | None:
    try:
        with get_session() as session:
            return session.get(cls, data_id)
    except Exception as e:
        logger = get_logger()
        logger.error("Get data by id failed", error=str(e))
        return None

def get_all_records(cls: Type[T]) -> list[T] | None:
    try:
        with get_session() as session:
            statement = select(cls)
            result = session.exec(statement).all()
            return cast(list[T], result)
    except Exception as e:
        logger = get_logger()
        logger.error("Get all records failed", error=str(e))
        return None

