package proto

type GetSASLRequest struct {
	Token []uint8 `jute:"token"`
}
