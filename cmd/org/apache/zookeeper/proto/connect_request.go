package proto

type ConnectRequest struct {
	ProtocolVersion uint32  `jute:"protocolVersion"`
	LastZxidSeen    uint64  `jute:"lastZxidSeen"`
	TimeOut         uint32  `jute:"timeOut"`
	SessionId       uint64  `jute:"sessionId"`
	Passwd          []uint8 `jute:"passwd"`
}
