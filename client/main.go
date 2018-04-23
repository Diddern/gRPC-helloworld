package main

import (
	"os"
	"log"
	"time"
	"strconv"
	"golang.org/x/net/context"
	"github.com/Diddern/gRPC-simpleGCDService/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	creds, err := credentials.NewClientTLSFromFile("../gcd/server-cert.pem", "")
	if err != nil {
		log.Fatalf("cert load error: %s", err)
	}

	// Connect securely to GCD service
	conn, err := grpc.Dial("localhost:5001", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to start gRPC connection: %v", err)
	}
	defer conn.Close()


	//Check for OS arguments
	if len(os.Args) != 3 {
		log.Fatal("Wrong number of arguments, please enter two integers.")
		os.Exit(2)
	}

	//Parse first argument
	a, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		log.Fatal("Invalid parameter A")
		return
	}
	//parse second argument
	b, err := strconv.ParseUint(os.Args[2], 10, 64)
	if err != nil {
		log.Fatal("Invalid parameter B")
		return
	}

	gcdClient := pb.NewGCDServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	gcdResponse, err := gcdClient.Compute(ctx, &pb.GCDRequest{A: a, B: b})
	if err != nil {
		log.Fatal("could not compute: ", err)
	}
	log.Print("The GCD of " + os.Args[1] + " and " + os.Args[2] + " = ", gcdResponse.Result)
}
