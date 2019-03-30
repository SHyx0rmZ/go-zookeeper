package txn

type Txn struct {
	Type uint32  `jute:"type"`
	Data []uint8 `jute:"data"`
}
