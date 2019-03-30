package txn

type TxnHeader struct {
	ClientId uint64 `jute:"clientId"`
	Cxid     uint32 `jute:"cxid"`
	Zxid     uint64 `jute:"zxid"`
	Time     uint64 `jute:"time"`
	Type     uint32 `jute:"type"`
}
