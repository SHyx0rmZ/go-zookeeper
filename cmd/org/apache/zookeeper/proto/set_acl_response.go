package proto

import (
	"code.witches.io/go/zookeeper/cmd/org/apache/zookeeper/data"
)

type SetACLResponse struct {
	Stat data.Stat `jute:"stat"`
}
