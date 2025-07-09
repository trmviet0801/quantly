import os
import dotenv

from sqlmodel import SQLModel, create_engine, Session

dotenv.load_dotenv()

engine = create_engine(os.getenv('MYSQL_DB_URL'))

def init_db():
    from models.account import Account
    from models.stock import Stock
    from models.trading_model import TradingModel
    from models.stock_price import StockPrice
    from models.position import Position
    from models.condition import Condition
    SQLModel.metadata.create_all(engine)

def get_session() -> Session:
    return Session(engine)

