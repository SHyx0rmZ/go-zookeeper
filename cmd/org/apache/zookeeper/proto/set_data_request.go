package proto

type SetDataRequest struct {
	Path    string  `jute:"path"`
	Data    []uint8 `jute:"data"`
	Version uint32  `jute:"version"`
}
