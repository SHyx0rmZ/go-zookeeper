package proto

import (
	"code.witches.io/go/zookeeper/cmd/org/apache/zookeeper/data"
)

type GetACLResponse struct {
	Acl  []data.ACL `jute:"acl"`
	Stat data.Stat  `jute:"stat"`
}
