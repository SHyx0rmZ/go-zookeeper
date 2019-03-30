package proto

type ConnectResponse struct {
	ProtocolVersion uint32  `jute:"protocolVersion"`
	TimeOut         uint32  `jute:"timeOut"`
	SessionId       uint64  `jute:"sessionId"`
	Passwd          []uint8 `jute:"passwd"`
}
