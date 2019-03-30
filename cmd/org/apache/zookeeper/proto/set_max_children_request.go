package proto

type SetMaxChildrenRequest struct {
	Path string `jute:"path"`
	Max  uint32 `jute:"max"`
}
