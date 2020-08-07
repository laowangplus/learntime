package main

import (
	"io"
	"log"
	"net"
	helloworld "test/grpc/protoc"

	"google.golang.org/grpc"
)

type GreeterServer struct {
}

func (s *GreeterServer) SayRoute(stream helloworld.Greeter_SayRouteServer) error {
	n := 0
	for {
		err := stream.Send(&helloworld.HelloReply{
			Message: "hello route",
		})
		if err != nil {
			return err
		}

		resp, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		n++

		log.Println(resp)
	}
}

func main() {
	serv := grpc.NewServer()
	helloworld.RegisterGreeterServer(serv, &GreeterServer{})
	lis, _ := net.Listen("tcp", ":8383")
	serv.Serve(lis)

}
