package main

import (
	a "announcement_worker/internal/components/announcement_server"
	"announcement_worker/internal/components/worker"
	c "announcement_worker/internal/controllers"
	r "announcement_worker/internal/routing"
	s "announcement_worker/internal/service"

	"github.com/gin-gonic/gin"
)

const (
	service_port               = ":8080"
	annoncement_client_address = "announcement_server:5005"
)

func main() {
	router := gin.Default()
	announcementsServiceClient := a.NewAnnouncementsServerClient(annoncement_client_address)
	worker := worker.NewWorker(announcementsServiceClient)
	announcementService := s.NewAnnouncementService(worker)
	announcementController := c.NewAnnouncementController(announcementService)
	r.CreateAnnouncementRoutes(announcementController, router)
	router.Run(service_port)
}
