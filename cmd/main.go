package main

import (
	"fmt"
	"net"

	"github.com/erteldg/grpchealthcheckservice/pkg/config"
	"github.com/erteldg/grpchealthcheckservice/pkg/model"
	pb "github.com/erteldg/grpchealthcheckservice/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	clientset, err := config.GetClientset()
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		panic(err)
	}

	fmt.Print("Server run on port :50053")

	s := grpc.NewServer()

	pb.RegisterStatusServiceServer(s, &model.Server{Clientset: clientset})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
