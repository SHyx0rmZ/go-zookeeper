package txn

type CreateSessionTxn struct {
	TimeOut uint32 `jute:"timeOut"`
}
