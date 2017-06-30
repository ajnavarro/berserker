package main

import (
	"net"

	"fmt"
	"github.com/src-d/berserker/extractor"
	"google.golang.org/grpc"
)

func main() {
	// TODO parametrize
	lis, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	extractor.RegisterExtractorServiceServer(grpcServer, extractor.NewExtractorServiceServer())

	fmt.Println("STARTED")

	panic(grpcServer.Serve(lis))
}
