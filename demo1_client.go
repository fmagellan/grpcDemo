package main

import (
    "context"
    "fmt"
    "time"

    "google.golang.org/grpc"
    "github.com/fmagellan/grpcDemo/demo1"
)

const (
    address = "localhost:50051"
    defaultName = "Magellan"
)

func main() {
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        fmt.Println("Not able to connect.")
        return
    }
    defer conn.Close()

    client := demo1.NewGreeterClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    request := &demo1.PersonName{ Name: "Magellan" }
    response, err := client.GetName(ctx, request)
    if err != nil {
        fmt.Println("Error in executing the RPC.")
        return
    }

    fmt.Println("Got the below details from the server")
    fmt.Println("Name:", response.GetName())
    fmt.Println("Age:", response.GetAge())
    fmt.Println("Year:", response.GetYear())
}
