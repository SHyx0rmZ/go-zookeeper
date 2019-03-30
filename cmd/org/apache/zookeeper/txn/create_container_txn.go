package txn

import (
	"code.witches.io/go/zookeeper/cmd/org/apache/zookeeper/data"
)

type CreateContainerTxn struct {
	Path           string     `jute:"path"`
	Data           []uint8    `jute:"data"`
	Acl            []data.ACL `jute:"acl"`
	ParentCVersion uint32     `jute:"parentCVersion"`
}
