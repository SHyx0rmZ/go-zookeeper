package proto

type MultiHeader struct {
	Type uint32 `jute:"type"`
	Done bool   `jute:"done"`
	Err  uint32 `jute:"err"`
}
