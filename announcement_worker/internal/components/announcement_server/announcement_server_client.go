package announcementserver

import (
	"log"

	pb "announcement_worker/internal/announcement_proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAnnouncementsServerClient(announcement_client_address string) pb.AnnouncementsServiceClient {
	conn, err := grpc.Dial(announcement_client_address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server at %s: %v", announcement_client_address, err)
	}
	log.Printf("Connected to announcement server at %s", announcement_client_address)
	return pb.NewAnnouncementsServiceClient(conn)
}
