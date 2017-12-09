package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	tagsURL = "https://hub.docker.com/v2/repositories/%s/%s/tags?page_size=200&page=1"
)

// TagsResponse defines data structure for JSON response of "Listing Image Tags"
type TagsResponse struct {
	Count    int64  `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Tag  `json:"results"`
}

// Tag defines data structure for each tag in TagsResponse
type Tag struct {
	Name        string    `json:"name"`
	FullSize    int64     `json:"full_size"`
	ID          int64     `json:"id"`
	Repository  int64     `json:"repository"`
	Creator     int64     `json:"creator"`
	LastUpdater int64     `json:"last_updater"`
	LastUpdated time.Time `json:"last_updated"`
	ImageID     string    `json:"image_id"`
	V2          bool      `json:"v2"`
}

// GetTags function is listing image tags on http://hub.docker.com
func GetTags(organization string, repo string) ([]Tag, error) {
	reqURL := fmt.Sprintf(tagsURL, organization, repo)
	tags := make([]Tag, 0, 10)
	var err error //We create this here, otherwise url will be rescoped with :=
	var response TagsResponse
	for err == nil {
		response.Next = ""

		res, err := http.Get(reqURL)
		if err != nil {
			log.Printf("Failed to call REST api: %s, %s", reqURL, err)
			return nil, err
		}

		data, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			log.Printf("Failed to get image tags: %s", err)
			return nil, err
		}

		err = json.Unmarshal(data, &response)
		tags = append(tags, response.Results...)
		if response.Next == "" {
			return tags, nil
		}
		reqURL = response.Next
	}

	return nil, err

	//res, err := http.Get(reqURL)
	//if err != nil {
	//	log.Printf("Failed to call REST api: %s, %s", reqURL, err)
	//	return nil, err
	//}
	//
	//data, err := ioutil.ReadAll(res.Body)
	//res.Body.Close()
	//if err != nil {
	//	log.Printf("Failed to get image tags: %s", err)
	//	return nil, err
	//}
	//var tagsRes TagsResponse
	//err = json.Unmarshal(data, &tagsRes)
	//return tagsRes.Results, nil
}

//
///Tags Returns tags for a given user/organization and repository
//func (registry *DockerHubRegistry) Tags(user, repository string) ([]string, error) {
//	url := registry.url("/v2/repositories/%s/%s/tags", user, repository)
//	tags := make([]string, 0, 10)
//	var err error //We create this here, otherwise url will be rescoped with :=
//	var response repositoriesResponse
//	for err == nil {
//		response.Next = ""
//		url, err = registry.getDockerHubPaginatedJson(url, &response)
//		for _, r := range response.Results {
//			tags = append(tags, r.Name)
//		}
//	}
//	if err != ErrNoMorePages {
//		return nil, err
//	}
//	return tags, nil
//}
