package service

import (
	pb "announcement_worker/internal/announcement_proto"
	"announcement_worker/internal/components/worker"
	"context"
)

type AnnouncementService interface {
	ProcessFile(ctx context.Context, path string) error
	GetAllAnnouncements() ([]*pb.Announcement, error)
}

type announcementService struct {
	worker worker.AnnouncementsWorker
}

func NewAnnouncementService(worker worker.AnnouncementsWorker) AnnouncementService {
	return &announcementService{
		worker: worker,
	}
}

func (a *announcementService) ProcessFile(ctx context.Context, path string) error {
	return a.worker.Work(ctx, "./data.json")
}

func (a *announcementService) GetAllAnnouncements() ([]*pb.Announcement, error) {
	return a.worker.GetAllAnnouncements()
}
