package elasticsearch

import (
	pb "announcement_server/internal/announcement_proto"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type RepoIface interface {
	Create(ctx context.Context, announcement Announcement) error
	GetAllAnnouncements(ctx context.Context) (*pb.GetAllAnnouncementsResponse, error)
}

type es struct {
	client *elasticsearch.Client
}

var indexName = "announcement"

func NewElasticsearch(address string) (RepoIface, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			address,
		},
	}

	elasticSearchClient, err := tryToSetupElasticsearch(cfg)
	if err != nil {
		return nil, err
	}
	return &es{
		client: elasticSearchClient,
	}, nil
}

func tryToSetupElasticsearch(cfg elasticsearch.Config) (*elasticsearch.Client, error) {
	for i := 0; i < 10; i++ {
		elasticSearchClient, err := elasticsearch.NewClient(cfg)
		if err != nil {
			fmt.Printf("error creating client: %s, waiting 5 seconds\n", err)
			time.Sleep(5 * time.Second)
			continue
		}
		res, err := elasticSearchClient.Indices.Exists([]string{indexName})

		if err != nil {
			fmt.Printf("error checking if index exists: %s, waiting 5 seconds\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		if res.StatusCode == 404 {
			mapping := `{
				"mappings": {
				  "properties": {
					"id": {
					  "type": "keyword"
					},
					"categories": {
					  "properties": {
						"subcategory": {
						  "type": "keyword"
						}
					  }
					},
					"title": {
					  "type": "object",
					  "properties": {
						"ro": {
						  "type": "text",
						  "analyzer": "romanian"
						},
						"ru": {
						  "type": "text",
						  "analyzer": "russian"
						}
					  }
					},
					"type": {
					  "type": "keyword"
					},
					"posted": {
					  "type": "date"
					}
				  }
				}
			  }
			  `

			req := esapi.IndicesCreateRequest{
				Index: indexName,
				Body:  strings.NewReader(mapping),
			}

			res, err := req.Do(context.Background(), elasticSearchClient)
			if err != nil {
				return nil, fmt.Errorf("Error creating index: %s", err)
			}
			defer res.Body.Close()

			if res.IsError() {
				return nil, fmt.Errorf("Error creating index: %s", res.String())
			} else {
				return elasticSearchClient, nil
			}
		}
	}
	return nil, fmt.Errorf("error creating index: %s", "timeout")
}

func (e *es) Create(ctx context.Context, announcement Announcement) error {
	fmt.Printf("Creating announcement: %v\n", announcement)
	jsonAnnouncement, err := json.Marshal(announcement)
	if err != nil {
		return err
	}

	req := esapi.IndexRequest{
		Index:      indexName,
		DocumentID: announcement.ID,
		Body:       bytes.NewReader(jsonAnnouncement),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), e.client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error indexing document ID=%s, response: %s", announcement.ID, res.String())
	}

	return nil
}

func (e *es) GetAllAnnouncements(ctx context.Context) (*pb.GetAllAnnouncementsResponse, error) {

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	res, err := e.client.Search(
		e.client.Search.WithContext(context.Background()),
		e.client.Search.WithIndex(indexName),
		e.client.Search.WithBody(&buf),
		e.client.Search.WithTrackTotalHits(true),
		e.client.Search.WithPretty(),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error: %s", res.String())
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	return BuildGetAllAnnouncementsResponse(r), nil
}
