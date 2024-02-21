package worker

import (
	pb "announcement_worker/internal/announcement_proto"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type AnnouncementsWorker interface {
	Work(ctx context.Context, path string) error
	GetAllAnnouncements() ([]*pb.Announcement, error)
}

type announcementsWorker struct {
	asc pb.AnnouncementsServiceClient
}

func NewWorker(asc pb.AnnouncementsServiceClient) AnnouncementsWorker {
	return &announcementsWorker{asc: asc}
}

func (f *announcementsWorker) Work(ctx context.Context, filename string) error {
	log.Printf("Reading file %s", filename)
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	log.Printf("File %s opened", filename)
	decoder := json.NewDecoder(file)
	_, err = decoder.Token()
	if err != nil {
		return err
	}
	// Here we can do one by one or send Announcements in chunks by x number of Announcements
	// For now I'll do one by one to keep it simple
	for decoder.More() {
		var a pb.Announcement
		err = decoder.Decode(&a)
		if err != nil {
			return err
		}

		response, err := f.asc.CreateAnnouncement(ctx, &a)
		if err != nil {
			return err
		}

		if response != nil && response.GetResponse() != 201 {
			fmt.Println("Received wrong response from gRPC server's CreateAnnouncement function: ", response.GetResponse())
			return fmt.Errorf("received wrong response from gRPC server's CreateAnnouncement function")
		}
		log.Printf("Announcement with ID %s created", a.Id)
	}

	return nil
}

func (f *announcementsWorker) GetAllAnnouncements() ([]*pb.Announcement, error) {
	response, err := f.asc.GetAllAnnouncements(context.Background(), &pb.GetAllAnnouncementsRequest{})
	if err != nil {
		return nil, err
	}

	return response.Announcements, nil
}
