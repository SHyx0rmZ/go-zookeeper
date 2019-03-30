package txn

type MultiTxn struct {
	Txns []Txn `jute:"txns"`
}
