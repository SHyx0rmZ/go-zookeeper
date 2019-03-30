package proto

type ExistsRequest struct {
	Path  string `jute:"path"`
	Watch bool   `jute:"watch"`
}
