package txn

import (
	"code.witches.io/go/zookeeper/cmd/org/apache/zookeeper/data"
)

type CreateTxnV0 struct {
	Path      string     `jute:"path"`
	Data      []uint8    `jute:"data"`
	Acl       []data.ACL `jute:"acl"`
	Ephemeral bool       `jute:"ephemeral"`
}
