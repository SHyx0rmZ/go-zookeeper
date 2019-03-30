package proto

type GetChildren2Request struct {
	Path  string `jute:"path"`
	Watch bool   `jute:"watch"`
}
