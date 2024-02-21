package elasticsearch

import pb "announcement_server/internal/announcement_proto"

type Announcement struct {
	ID         string     `json:"id"`
	Categories Categories `json:"categories"`
	Title      Title      `json:"title"`
	Type       string     `json:"type"`
	Posted     float64    `json:"posted"`
}

type Categories struct {
	Subcategory string `json:"subcategory"`
}

type Title struct {
	Ro string `json:"ro"`
	Ru string `json:"ru"`
}

func BuildAnnouncementDto(announcement *pb.Announcement) *Announcement {
	return &Announcement{
		ID:         announcement.Id,
		Categories: Categories{Subcategory: announcement.Categories.Subcategory},
		Title:      Title{Ro: announcement.Title.Ro, Ru: announcement.Title.Ru},
		Type:       announcement.Type,
		Posted:     announcement.Posted,
	}
}

func BuildGetAllAnnouncementsResponse(r map[string]interface{}) *pb.GetAllAnnouncementsResponse {
	response := &pb.GetAllAnnouncementsResponse{}
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		announcement := &pb.Announcement{
			Id: source["id"].(string),
			Categories: &pb.Announcement_Categories{
				Subcategory: source["categories"].(map[string]interface{})["subcategory"].(string),
			},
			Title: &pb.Announcement_Title{
				Ro: source["title"].(map[string]interface{})["ro"].(string),
				Ru: source["title"].(map[string]interface{})["ru"].(string),
			},
			Type:   source["type"].(string),
			Posted: source["posted"].(float64),
		}
		response.Announcements = append(response.Announcements, announcement)
	}
	return response
}
