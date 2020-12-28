package main

import (
    "context"
    "fmt"
    "net"

    "google.golang.org/grpc"
    "github.com/fmagellan/grpcDemo/demo1"
)

type demoServer struct {
    demo1.UnimplementedGreeterServer
}

func (server *demoServer) GetName(ctx context.Context, request *demo1.PersonName) (*demo1.PersonDetails, error) {
    fmt.Println("Received a query for the person:", request.GetName())
    response := &demo1.PersonDetails{
        Name: "Magellan",
        Age: 33,
        Year: 1519,
    }
    return response, nil
}

func main() {
    listener, err := net.Listen("tcp", ":50051")
    if err != nil {
        fmt.Println("Not able to listen.")
        fmt.Println(err.Error())
        return
    }

    server := grpc.NewServer()
    demo1.RegisterGreeterServer(server, &demoServer{})
    err = server.Serve(listener)
    if err != nil {
        fmt.Println("Not able to start gRPC server")
        return
    }

    fmt.Println("Done with server")
}
