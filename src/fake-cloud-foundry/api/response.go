package api

import (
	"strconv"
	"time"
)

type Response struct {
	TotalResults string     `json:"total_results"`
	TotalPages   string     `json:"total_pages"`
	PrevURL      *string    `json:"prev_url"`
	NextURL      *string    `json:"next_url"`
	Resources    []Resource `json:"resources"`
}

type Resource struct {
	Metadata Metadata    `json:"metadata"`
	Entity   interface{} `json:"entity"`
}

type Metadata struct {
	Guid      string `json:"guid"`
	URL       string `json:"url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func newResponse(resources []Resource) Response {
	now := time.Now().Format(time.RFC3339)

	for _, resource := range resources {
		resource.Metadata.CreatedAt = now
		resource.Metadata.UpdatedAt = now
	}

	return Response{
		TotalResults: strconv.Itoa(len(resources)),
		TotalPages:   "1",
		PrevURL:      nil,
		NextURL:      nil,
		Resources:    resources,
	}
}