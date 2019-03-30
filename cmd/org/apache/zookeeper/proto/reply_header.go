package proto

type ReplyHeader struct {
	Xid  uint32 `jute:"xid"`
	Zxid uint64 `jute:"zxid"`
	Err  uint32 `jute:"err"`
}
