package main

import (
	"fmt"
	"net"
	"protocolobuf/service/school"

	proto "google.golang.org/protobuf/proto"
)

func main() {
	list, err := net.Listen("tcp", ":4000")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := list.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println(conn.LocalAddr().String())
		go func() {
			buffer := make([]byte, 1024)
			_, err := conn.Read(buffer)
			if err != nil {
				conn.Close()
			}
			var message school.Professor
			if err := proto.Unmarshal(buffer, &message); err != nil {
				conn.Close()
			}
			fmt.Println(&message)
		}()
	}
}
