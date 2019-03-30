package proto

type SetSASLRequest struct {
	Token []uint8 `jute:"token"`
}
