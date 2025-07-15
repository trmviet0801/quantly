from models.condition_match_group import ConditionMatchGroup
from models.condition import Condition

def get_condition_key(condition: Condition) ->  str:
    """
    Get Redis key for specific condition (e,g, SMA with period == 2 -> SMA(20)
    Args: models.Condition
    Returns: redis key (string)
        - scalar indicators: [indicator]([period])([operation])([value]) (e.g. SMA(20)(>)(500))
        - BollingerBands: BollingerBands(lower | upper | sma)([operation]) (e.g. BollingerBands(lower)(<))
        - StochasticOscillator: StochasticOscillator([operation]) (e.g. StochasticOscillator(<))

    """
    indicator = condition.indicator
    if indicator == 'SMA' or indicator == 'EMA' or indicator == 'RSI' or indicator == 'ATR' or indicator == 'Momentum':
        return indicator + '(' + str(condition.period) + ')' + '(' + condition.operation + ')(' + str(condition.value) + ')'
    elif indicator == 'BollingerBands':
        return indicator + '(' + condition.band + ')(' + condition.operation + ')'
    else:
        return indicator + '(' + condition.operation + ')'


