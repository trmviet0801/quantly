from models.cond_model_index import CondModelIndex

def test_unique_id():
    var1 = CondModelIndex()
    var2 = CondModelIndex()
    assert var1.cond_model_index_id != var2.cond_model_index_id