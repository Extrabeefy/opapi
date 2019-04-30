package main

import (
	"log"
	"net"

	pb "github.com/opAPIProgression/progression-service/proto/progression"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type IRepository interface {
	Create(*pb.Boss) (*pb.Boss, error)
	GetAll() []*pb.Boss
}

type Repository struct {
	bosses []*pb.Boss
}

func (repo *Repository) Create(boss *pb.Boss) (*pb.Boss, error) {
	updated := append(repo.bosses, boss)
	repo.bosses = updated
	return boss, nil
}

func (repo *Repository) GetAll() []*pb.Boss {
	return repo.bosses
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	repo IRepository
}

// CreateBoss - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) CreateBoss(ctx context.Context, req *pb.Boss) (*pb.Response, error) {

	// Save our boss
	boss, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	return &pb.Response{Created: true, Boss: boss}, nil
}

func (s *service) GetBosses(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	bosses := s.repo.GetAll()
	return &pb.Response{Bosses: bosses}, nil
}

func main() {
	repo := &Repository{}

	// Set-up our gRPC server.
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	pb.RegisterProgressionServiceServer(s, &service{repo})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
