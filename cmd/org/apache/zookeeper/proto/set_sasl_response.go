package proto

type SetSASLResponse struct {
	Token []uint8 `jute:"token"`
}
