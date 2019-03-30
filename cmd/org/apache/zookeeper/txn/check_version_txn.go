package txn

type CheckVersionTxn struct {
	Path    string `jute:"path"`
	Version uint32 `jute:"version"`
}
