package client

import (
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetArtifactMetadata(repository string, groupId string, artifactId string, version string, extension string) (*ArtifactMetadata, error) {
	uri := fmt.Sprintf("%s/%s/%s/%s", repository, strings.Replace(groupId, ".", "/", -1), artifactId, version)
	fileName := fmt.Sprintf("%s-%s.%s", artifactId, version, extension)
	resp, err := http.Get(uri)
	defer resp.Body.Close()

	result := &ArtifactMetadata{
		GroupId:    groupId,
		ArtifactId: artifactId,
		Version:    version,
	}

	if err != nil {
		return result, err
	}

	if code := resp.StatusCode; code != http.StatusOK {
		return nil, fmt.Errorf("%d GET %s", code, uri)
	}

	htmlTokens := html.NewTokenizer(resp.Body)
loop:
	for {
		tt := htmlTokens.Next()
		switch tt {
		case html.ErrorToken:
			break loop
		case html.TextToken:
			t := htmlTokens.Token()
			if t.Data == fileName {
				htmlTokens.Next()
				htmlTokens.Next()
				htmlTokens.Next()
				htmlTokens.Next()
				htmlTokens.Next()
				artifactDate := strings.TrimSpace(htmlTokens.Token().Data)
				lastModified, error := time.Parse("Mon Jan 02 15:04:05 MST 2006", artifactDate)
				if error != nil {
					return result, error
				}
				result.LastModified = lastModified
				htmlTokens.Next()
				htmlTokens.Next()
				htmlTokens.Next()
				htmlTokens.Next()
				size, error := strconv.ParseInt(strings.TrimSpace(htmlTokens.Token().Data), 10, 64)
				if error != nil {
					return result, error
				}
				result.Size = size
				break loop
			}
			break
		case html.SelfClosingTagToken:
			break
		case html.StartTagToken:
			break
		}
	}
	return result, nil
}

func GetMetadata(repository string, groupId string, artifactId string) (Metadata, error) {
	var metadata Metadata

	uri := fmt.Sprintf("%s/%s/%s/maven-metadata.xml", repository, strings.Replace(groupId, ".", "/", -1), artifactId)
	resp, err := http.Get(uri)
	defer resp.Body.Close()
	if err != nil {
		return metadata, err
	}

	if code := resp.StatusCode; code != http.StatusOK {
		return metadata, fmt.Errorf("%d GET %s", code, uri)
	}

	if err := xml.NewDecoder(resp.Body).Decode(&metadata); err != nil {
		return metadata, err
	}

	return metadata, nil
}

type ArtifactMetadata struct {
	GroupId      string
	ArtifactId   string
	Version      string
	LastModified time.Time
	Size         int64
}

type Metadata struct {
	XMLName    xml.Name   `xml:"metadata"`
	GroupID    string     `xml:"groupId"`
	ArtifactID string     `xml:"artifactId"`
	Versioning Versioning `xml:"versioning`
}

type Versioning struct {
	XMLName     xml.Name `xml:"versioning"`
	Latest      string   `xml:"latest"`
	Release     string   `xml:"release"`
	Versions    []string `xml:"versions>version"`
	LastUpdated string   `xml:"lastUpdated"`
}
