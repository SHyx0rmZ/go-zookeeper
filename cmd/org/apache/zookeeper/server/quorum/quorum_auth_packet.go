package quorum

type QuorumAuthPacket struct {
	Magic  uint64  `jute:"magic"`
	Status uint32  `jute:"status"`
	Token  []uint8 `jute:"token"`
}
