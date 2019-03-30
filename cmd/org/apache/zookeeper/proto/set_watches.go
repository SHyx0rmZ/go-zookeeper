package proto

type SetWatches struct {
	RelativeZxid uint64   `jute:"relativeZxid"`
	DataWatches  []string `jute:"dataWatches"`
	ExistWatches []string `jute:"existWatches"`
	ChildWatches []string `jute:"childWatches"`
}
