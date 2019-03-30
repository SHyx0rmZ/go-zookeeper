package persistence

type FileHeader struct {
	Magic   uint32 `jute:"magic"`
	Version uint32 `jute:"version"`
	Dbid    uint64 `jute:"dbid"`
}
