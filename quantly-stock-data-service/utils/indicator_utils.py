from pandas import Series


def average_close_price(data: Series) -> float:
    num = len(data)
    if num == 0:
        return 0
    sum = 0
    for price in data:
        sum += price
    return sum / num