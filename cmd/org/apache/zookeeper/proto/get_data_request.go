package proto

type GetDataRequest struct {
	Path  string `jute:"path"`
	Watch bool   `jute:"watch"`
}
