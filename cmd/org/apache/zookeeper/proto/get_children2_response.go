package proto

import (
	"code.witches.io/go/zookeeper/cmd/org/apache/zookeeper/data"
)

type GetChildren2Response struct {
	Children []string  `jute:"children"`
	Stat     data.Stat `jute:"stat"`
}
