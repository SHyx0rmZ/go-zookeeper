package quorum

import (
	"code.witches.io/go/zookeeper/cmd/org/apache/zookeeper/data"
)

type QuorumPacket struct {
	Type     uint32    `jute:"type"`
	Zxid     uint64    `jute:"zxid"`
	Data     []uint8   `jute:"data"`
	Authinfo []data.Id `jute:"authinfo"`
}
