package routing

import (
	"net/http"

	c "announcement_worker/internal/controllers"

	"github.com/gin-gonic/gin"
)

func CreateAnnouncementRoutes(announcementController c.AnnouncementController, engine *gin.Engine) {
	announcements := engine.Group("/announcements")

	announcements.Handle(http.MethodPost, "", announcementController.UploadAnnouncements)
	announcements.Handle(http.MethodGet, "", announcementController.GetAllAnnouncements)
}
