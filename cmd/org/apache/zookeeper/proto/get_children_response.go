package proto

type GetChildrenResponse struct {
	Children []string `jute:"children"`
}
