package main

import (
	"bytes"
	"code.witches.io/go/zookeeper/cmd/org/apache/zookeeper/proto"
	"code.witches.io/go/zookeeper/jute"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "[::1]:2181")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	req := proto.ConnectRequest{
		TimeOut: 30000,
		Passwd:  make([]byte, 16),
	}

	bs, err := jute.Marshal(&req)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(bs), bs)

	var length [4]byte

	binary.BigEndian.PutUint32(length[:], uint32(len(bs)))

	_, err = io.Copy(conn, io.MultiReader(bytes.NewReader(length[:]), bytes.NewReader(bs)))
	if err != nil {
		panic(err)
	}

	_, err = io.ReadFull(conn, length[:])
	if err != nil {
		panic(err)
	}

	bs = make([]byte, binary.BigEndian.Uint32(length[:]))

	_, err = io.ReadFull(conn, bs)
	if err != nil {
		panic(err)
	}

	var resp proto.ConnectResponse

	err = jute.Unmarshal(bs, &resp)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", resp)

	req2 := struct {
		proto.RequestHeader
		proto.GetChildrenRequest
	}{
		proto.RequestHeader{
			Xid:  1,
			Type: 8,
		},
		proto.GetChildrenRequest{
			Path:  "/",
			Watch: false,
		},
	}

	bs2, err := jute.Marshal(&req2)
	if err != nil {
		panic(err)
	}

	length = [4]byte{}

	binary.BigEndian.PutUint32(length[:], uint32(len(bs2)))

	_, err = io.Copy(conn, io.MultiReader(bytes.NewReader(length[:]), bytes.NewReader(bs2)))
	if err != nil {
		panic(err)
	}

	_, err = io.ReadFull(conn, length[:])
	if err != nil {
		panic(err)
	}

	bs2 = make([]byte, binary.BigEndian.Uint32(length[:]))

	_, err = io.ReadFull(conn, bs2)
	if err != nil {
		panic(err)
	}

	var resp2 struct {
		proto.ReplyHeader
		proto.GetChildrenResponse
	}

	err = jute.Unmarshal(bs2, &resp2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", resp2)
}

func main2() {
	conn, err := net.Dial("tcp", "[::1]:2181")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buf := make([]byte, 49)
	binary.BigEndian.PutUint32(buf[0x00:], uint32(len(buf)-4))
	binary.BigEndian.PutUint32(buf[0x04:], 0)                     // protocol version
	binary.BigEndian.PutUint64(buf[0x0C:], 0)                     // last zxid seen
	binary.BigEndian.PutUint32(buf[0x10:], 30000)                 // timeout
	binary.BigEndian.PutUint64(buf[0x14:], 0)                     // session id
	binary.BigEndian.PutUint32(buf[0x1C:], uint32(len(buf)-32-1)) // buffer size

	_, err = io.Copy(conn, bytes.NewReader(buf))
	if err != nil {
		panic(err)
	}

	var size [4]byte

	_, err = io.ReadFull(conn, size[:])
	if err != nil {
		panic(err)
	}
	fmt.Println(binary.BigEndian.Uint32(size[:]))

	buf = make([]byte, binary.BigEndian.Uint32(size[:]))
	_, err = io.ReadFull(conn, buf)
	fmt.Println(binary.BigEndian.Uint32(buf[0x00:])) // protocol version
	fmt.Println(binary.BigEndian.Uint32(buf[0x04:])) // timeout
	fmt.Println(binary.BigEndian.Uint64(buf[0x08:])) // session id
	bufSize := binary.BigEndian.Uint32(buf[0x10:])
	fmt.Println(bufSize)
	bufData := make([]byte, bufSize+1)
	n := copy(bufData, buf[0x14:])
	if n != int(bufSize)+1 {
		panic("mismatch")
	}
	fmt.Println(bufData)
}

type ConnectRequest struct {
	ProtocolVersion uint32
	LastZXIDSeen    uint64
	Timeout         uint32
	SessionID       uint64
	Password        [16]byte
}

type ConnectResponse struct {
	ProtocolVersion uint32
	Timeout         uint32
	SessionID       uint64
	Password        [16]byte
}
