package main

import (
	"belajar/grpc/proto"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:4040"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewCalculateClient(conn)

	// Contact the server and print out its response.
	var numb1, numb2 int64
	if len(os.Args) > 2 {
		numbone, errNumb1 := strconv.ParseInt(os.Args[1], 10, 64)
		numbtwo, errNumb2 := strconv.ParseInt(os.Args[2], 10, 64)
		if errNumb1 != nil || errNumb2 != nil {
			fmt.Printf("invalid number: %v or %v\n", os.Args[1], os.Args[2])
			return
		} else {
			numb1 = numbone
			numb2 = numbtwo
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Add(ctx, &proto.Request{A: numb1, B: numb2})
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	fmt.Println(r.GetResult())
	r, err = c.Multiply(ctx, &proto.Request{A: numb1, B: numb2})
	if err != nil {
		log.Fatalf("could not multiply: %v", err)
	}
	fmt.Println(r.GetResult())
}
