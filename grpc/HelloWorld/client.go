package main

import (
	"context"
	"io"
	"log"
	helloworld "test/grpc/protoc"

	"google.golang.org/grpc"
)

func SayRoute(client helloworld.GreeterClient, r *helloworld.HelloRequest) error {
	stream, err := client.SayRoute(context.Background())
	if err != nil {
		return err
	}

	for i := 0; i <= 6; i++ {
		err := stream.Send(r)
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

		log.Println(resp)
	}
	return nil
}

func main() {
	conn, _ := grpc.Dial(":8383", grpc.WithInsecure())
	defer conn.Close()

	client := helloworld.NewGreeterClient(conn)
	req := &helloworld.HelloRequest{
		Name: "hello laowang",
	}
	err := SayRoute(client, req)
	if err != nil {
		log.Fatal(err)
	}
}
