package tx

type TransactionMgr struct {
}

func (txM *TransactionMgr) onTxCommit(tx Transaction) {

}

func (txM *TransactionMgr) onTxRollback(tx Transaction) {

}

func (txM *TransactionMgr) onTxEndStatement(tx Transaction) {

}

func (txM *TransactionMgr) NewTransaction(isolationLevel int, readOnly bool) *Transaction {
	return nil
}

func (txM *TransactionMgr) NewTransactionWithTxNum(isolationLevel int, readOnly bool, txNum int64) *Transaction {
	return nil
}
