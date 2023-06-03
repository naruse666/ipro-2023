package main

import (
	"context"
	"log"
	"time"

	pb "github.com/naruse666/ipro-2023/app/api_server/grpc"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSuicideServiceServer
}

func main() {
    conn, err := grpc.Dial("localhost:8010", grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewSuicideServiceClient(conn)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := c.SuicideRequest(ctx, &pb.Request{Request: "test"})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }
    //log.Printf("RECV: %s", r.ProtoMessage())
    log.Printf("RECV: %s", r)
}
