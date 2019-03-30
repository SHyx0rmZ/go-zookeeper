package proto

import (
	"code.witches.io/go/zookeeper/cmd/org/apache/zookeeper/data"
)

type CreateTTLRequest struct {
	Path  string     `jute:"path"`
	Data  []uint8    `jute:"data"`
	Acl   []data.ACL `jute:"acl"`
	Flags uint32     `jute:"flags"`
	Ttl   uint64     `jute:"ttl"`
}
