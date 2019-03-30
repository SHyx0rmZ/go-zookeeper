package txn

type ErrorTxn struct {
	Err uint32 `jute:"err"`
}
