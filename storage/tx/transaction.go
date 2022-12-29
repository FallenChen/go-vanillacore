package tx

type Transaction struct {
	TxNum    int64
	ReadOnly bool
}

func (tx *Transaction) Commit() {

}

func (tx *Transaction) Rollback() {

}

func (tx *Transaction) EndStatement() {

}
