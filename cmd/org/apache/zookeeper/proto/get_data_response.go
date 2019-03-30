package proto

import (
	"code.witches.io/go/zookeeper/cmd/org/apache/zookeeper/data"
)

type GetDataResponse struct {
	Data []uint8   `jute:"data"`
	Stat data.Stat `jute:"stat"`
}
