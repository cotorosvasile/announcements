package controllers

import (
	s "announcement_worker/internal/service"

	"github.com/gin-gonic/gin"
)

type AnnouncementController interface {
	UploadAnnouncements(ctx *gin.Context)
	GetAllAnnouncements(ctx *gin.Context)
}

type announcementController struct {
	announcementService s.AnnouncementService
}

func NewAnnouncementController(announcementService s.AnnouncementService) AnnouncementController {
	return &announcementController{
		announcementService: announcementService,
	}
}

func (a *announcementController) UploadAnnouncements(ctx *gin.Context) {
	if err := a.announcementService.ProcessFile(ctx, "./data.json"); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
}

func (a *announcementController) GetAllAnnouncements(ctx *gin.Context) {
	announcements, err := a.announcementService.GetAllAnnouncements()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, announcements)
}
