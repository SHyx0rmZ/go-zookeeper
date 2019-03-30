package proto

type RequestHeader struct {
	Xid  uint32 `jute:"xid"`
	Type uint32 `jute:"type"`
}
