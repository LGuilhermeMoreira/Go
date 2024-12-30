package main

import (
	"fmt"
	"log"
	"net"
	"protocolobuf/service/school"

	"google.golang.org/protobuf/proto"
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
			professor := school.Professor{
				Name:       "Thigas",
				Discipline: "POO",
				Code:       "QXD004521",
				Age:        51,
			}
			data, err := proto.Marshal(&professor)
			if err != nil {
				log.Println(err)
				conn.Close()
			}
			conn.Write(data)
		}()
	}
}
