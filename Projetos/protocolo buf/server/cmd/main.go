package main

import (
	"fmt"
	"log"
	"net"

	pb "protocolobuf/service/school"

	"google.golang.org/grpc"
)

type classServer struct {
	pb.UnimplementedClassServiceServer
}

func (s *classServer) ListAllStudents(req *pb.InputListAllStudents, stream pb.ClassService_ListAllStudentsServer) error {
	students := []*pb.Student{
		{Name: "Alice", Age: 20, Code: "S123"},
		{Name: "Bob", Age: 21, Code: "S456"},
	}
	for _, student := range students {
		if err := stream.Send(student); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterClassServiceServer(grpcServer, &classServer{})
	fmt.Println("Servidor gRPC ouvindo na porta 4000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
