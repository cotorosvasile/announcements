package service

import (
	pb "announcement_server/internal/announcement_proto"
	"context"
	"fmt"
	"log"

	repo "announcement_server/internal/components/elasticsearch"
)

type announcementService struct {
	repo repo.RepoIface
	pb.UnimplementedAnnouncementsServiceServer
}

func NewAnnouncementService(es repo.RepoIface) announcementService {
	return announcementService{repo: es}
}

func (a *announcementService) CreateAnnouncement(ctx context.Context, announcement *pb.Announcement) (*pb.AnnouncementResponse, error) {
	err := a.repo.Create(ctx, *repo.BuildAnnouncementDto(announcement))
	if err != nil {
		log.Printf("error creating announcement: %v", err)
		return nil, fmt.Errorf("error creating announcement: %v", err)
	}
	return &pb.AnnouncementResponse{
		Response: 201,
	}, nil
}

func (a *announcementService) GetAllAnnouncements(ctx context.Context, _ *pb.GetAllAnnouncementsRequest) (*pb.GetAllAnnouncementsResponse, error) {
	response, err := a.repo.GetAllAnnouncements(ctx)
	if err != nil {
		return nil, fmt.Errorf("error retrieving announcements: %v", err)
	}
	return response, nil
}
