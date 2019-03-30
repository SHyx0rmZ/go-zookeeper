package proto

type GetChildrenRequest struct {
	Path  string `jute:"path"`
	Watch bool   `jute:"watch"`
}
