package api

import (
	"strconv"
)

type Response struct {
	TotalResults string     `json:"total_results"`
	TotalPages   string     `json:"total_pages"`
	PrevURL      *string    `json:"prev_url"`
	NextURL      *string    `json:"next_url"`
	Resources    []Resource `json:"resources"`
}

type V3Response struct {
	Pagination Pagination   `json:"pagination"`
	Resources  []V3Resource `json:"resources"`
}

type V3Resource map[string]interface{}

type Pagination struct {
	First        Link    `json:"first"`
	Last         Link    `json:"first"`
	Next         *string `json:"next"`
	Previous     *string `json:"previous"`
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
}

type Link struct {
	Href string `json:"href"`
}

type Resource struct {
	Metadata Metadata    `json:"metadata"`
	Entity   interface{} `json:"entity"`
}

type Metadata struct {
	Guid      string `json:"guid"`
	URL       string `json:"url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

func newResponse(resources []Resource) Response {
	return Response{
		TotalResults: strconv.Itoa(len(resources)),
		TotalPages:   "1",
		PrevURL:      nil,
		NextURL:      nil,
		Resources:    resources,
	}
}

func newV3Response(resources []V3Resource, path string) V3Response {
	return V3Response{
		Pagination: Pagination{
			First: Link{
				Href: path,
			},
			Last: Link{
				Href: path,
			},
			TotalPages:   1,
			TotalResults: 1,
		},
		Resources: resources,
	}
}
