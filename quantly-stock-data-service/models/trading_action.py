from dataclasses import dataclass

@dataclass()
class TradingAction():
    action: str = None
    symbol: str = None
    quantity: int = None