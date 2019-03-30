package proto

type RemoveWatchesRequest struct {
	Path string `jute:"path"`
	Type uint32 `jute:"type"`
}
