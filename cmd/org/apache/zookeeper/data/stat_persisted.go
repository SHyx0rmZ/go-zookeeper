package data

type StatPersisted struct {
	Czxid          uint64 `jute:"czxid"`
	Mzxid          uint64 `jute:"mzxid"`
	Ctime          uint64 `jute:"ctime"`
	Mtime          uint64 `jute:"mtime"`
	Version        uint32 `jute:"version"`
	Cversion       uint32 `jute:"cversion"`
	Aversion       uint32 `jute:"aversion"`
	EphemeralOwner uint64 `jute:"ephemeralOwner"`
	Pzxid          uint64 `jute:"pzxid"`
}
