package txn

type SetMaxChildrenTxn struct {
	Path string `jute:"path"`
	Max  uint32 `jute:"max"`
}
