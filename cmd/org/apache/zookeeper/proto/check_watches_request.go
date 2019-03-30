package proto

type CheckWatchesRequest struct {
	Path string `jute:"path"`
	Type uint32 `jute:"type"`
}
