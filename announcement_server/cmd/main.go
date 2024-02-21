package main

import (
	pb "announcement_server/internal/announcement_proto"
	repo "announcement_server/internal/components/elasticsearch"
	"announcement_server/internal/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	service_port         = ":5005"
	elasticsearchAddress = "http://elasticsearch:9200"
)

func main() {
	lis, err := net.Listen("tcp", service_port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", service_port, err)
	}

	s := grpc.NewServer()

	elasticsearch, err := repo.NewElasticsearch(elasticsearchAddress)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	service := service.NewAnnouncementService(elasticsearch)
	pb.RegisterAnnouncementsServiceServer(s, &service)

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
