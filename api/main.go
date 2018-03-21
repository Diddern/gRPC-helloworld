package main

import (
	"os"
	"log"
	"strconv"
	"golang.org/x/net/context"
	"github.com/Diddern/gRPC-helloworld/pb"
	"google.golang.org/grpc"

	"time"
)

func main() {
	// Connect to GCD service
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}

	//defer conn.Close()


	if len(os.Args) < 2 {
		log.Fatal("No arguments provided, please enter two numbers.")
		os.Exit(2)
	}
	a, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal("Invalid parameter A")
		return
	}
	b, err := strconv.ParseUint(os.Args[2], 10, 64)
	if err != nil {
		log.Fatal("Invalid parameter B")
		return
	}

	gcdClient := pb.NewGCDServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := gcdClient.Compute(ctx, &pb.GCDRequest{A: a, B: b})
	if err != nil {
		log.Fatal("could not compute: ", err)
	}
	log.Print("The GCD of " + os.Args[1] + " and " + os.Args[2] + " = ", r.Result)
}

