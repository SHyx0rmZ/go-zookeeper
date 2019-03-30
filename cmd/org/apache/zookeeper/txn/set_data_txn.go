package txn

type SetDataTxn struct {
	Path    string  `jute:"path"`
	Data    []uint8 `jute:"data"`
	Version uint32  `jute:"version"`
}
