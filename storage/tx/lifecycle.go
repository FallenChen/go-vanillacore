package tx

type TransactionLifecycle interface {
	onTxCommit(tx Transaction)

	onTxRollback(tx Transaction)

	onTxEndStatement(tx Transaction)
}
