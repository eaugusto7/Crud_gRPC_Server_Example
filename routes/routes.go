package routes

import (
	"log"
	"net"
	"users/grpc/controllers"
	"users/grpc/pb"

	"google.golang.org/grpc"
)

func StartServerGRPC() {
	lis, err := net.Listen("tcp", "localhost:9000")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterUserServerServer(s, &controllers.Server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
