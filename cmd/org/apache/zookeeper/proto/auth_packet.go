package proto

type AuthPacket struct {
	Type   uint32  `jute:"type"`
	Scheme string  `jute:"scheme"`
	Auth   []uint8 `jute:"auth"`
}
