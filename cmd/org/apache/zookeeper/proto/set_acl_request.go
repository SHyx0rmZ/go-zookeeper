package proto

import (
	"code.witches.io/go/zookeeper/cmd/org/apache/zookeeper/data"
)

type SetACLRequest struct {
	Path    string     `jute:"path"`
	Acl     []data.ACL `jute:"acl"`
	Version uint32     `jute:"version"`
}
