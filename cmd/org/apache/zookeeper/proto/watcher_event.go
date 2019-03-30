package proto

type WatcherEvent struct {
	Type  uint32 `jute:"type"`
	State uint32 `jute:"state"`
	Path  string `jute:"path"`
}
