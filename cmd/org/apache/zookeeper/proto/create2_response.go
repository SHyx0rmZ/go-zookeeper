package proto

import (
	"code.witches.io/go/zookeeper/cmd/org/apache/zookeeper/data"
)

type Create2Response struct {
	Path string    `jute:"path"`
	Stat data.Stat `jute:"stat"`
}
