package txn

import (
	"code.witches.io/go/zookeeper/cmd/org/apache/zookeeper/data"
)

type SetACLTxn struct {
	Path    string     `jute:"path"`
	Acl     []data.ACL `jute:"acl"`
	Version uint32     `jute:"version"`
}
