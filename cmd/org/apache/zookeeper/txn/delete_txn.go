package txn

type DeleteTxn struct {
	Path string `jute:"path"`
}
