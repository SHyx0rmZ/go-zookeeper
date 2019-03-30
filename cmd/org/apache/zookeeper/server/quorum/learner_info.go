package quorum

type LearnerInfo struct {
	Serverid        uint64 `jute:"serverid"`
	ProtocolVersion uint32 `jute:"protocolVersion"`
	ConfigVersion   uint64 `jute:"configVersion"`
}
