package proto

type CheckVersionRequest struct {
	Path    string `jute:"path"`
	Version uint32 `jute:"version"`
}
