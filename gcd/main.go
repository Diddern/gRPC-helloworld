package main
import (
	"log"
	"net"
	"github.com/Diddern/gRPC-simpleGCDService/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/credentials"
)

type server struct{}

func main() {

	portNumber := ":3000"
	pathToCert := "gcd/server-cert.pem"
	pathToKey := "gcd/server-key.pem"

	//Load cert and key from file
	creds, err := credentials.NewServerTLSFromFile(pathToCert, pathToKey)
	if err != nil {
		log.Fatalf("Failed to setup tls: %v", err)
	}

	//Listen for incoming connections.
	lis, err := net.Listen("tcp", portNumber)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Listening on port %v", portNumber)

	//Create gRPC Server
	s := grpc.NewServer(
		grpc.Creds(creds),
	)
	pb.RegisterGCDServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) Compute(ctx context.Context, r *pb.GCDRequest) (*pb.GCDResponse, error) {
	a, b := r.A, r.B
	for b != 0 {
		a, b = b, a%b
	}
	return &pb.GCDResponse{Result: a}, nil
}