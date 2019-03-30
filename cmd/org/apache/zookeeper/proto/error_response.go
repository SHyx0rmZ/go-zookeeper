package proto

type ErrorResponse struct {
	Err uint32 `jute:"err"`
}
