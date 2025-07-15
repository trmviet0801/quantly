from dataclasses import dataclass
from pandas import Series

@dataclass
class ConditionMatchGroup:
    condition: str
    matching_stocks: set[str]



