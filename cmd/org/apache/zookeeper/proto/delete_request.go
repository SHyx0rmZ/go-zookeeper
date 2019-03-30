package proto

type DeleteRequest struct {
	Path    string `jute:"path"`
	Version uint32 `jute:"version"`
}
