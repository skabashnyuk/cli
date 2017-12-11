package client

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestShouldBeAbleToGetListOfCheArtifacts(t22 *testing.T) {

	metadata, error := GetMetadata("http://maven.codenvycorp.com/content/groups/public", "org.eclipse.che", "assembly-main")
	if error != nil {
		t22.Error(error)
	}
	if len(metadata.Versioning.Versions) < 1 {
		t22.Error("Tags list should not be empty")
	}
	fmt.Printf("%v", metadata.Versioning.Versions)
}

func TestShouldBeAbleToGetTimeOfLastModified(t *testing.T) {
	expectedTime, _ := time.Parse(time.RFC1123, "Thu, 03 Mar 2016 20:39:25 UTC")
	expected := &ArtifactMetadata{
		GroupId:      "org.eclipse.che",
		ArtifactId:   "assembly-main",
		Version:      "4.0.0-RC11",
		LastModified: expectedTime,
		Size:         138847044,
	}

	metadata, error := GetArtifactMetadata(
		"http://maven.codenvycorp.com/content/groups/public",
		"org.eclipse.che",
		"assembly-main",
		"4.0.0-RC11",
		"tar.gz")
	if error != nil {
		t.Error(error)
	}

	// assert equality
	assert.Equal(t, expected, metadata, "they should be equal")

}

func TestShouldParseDate(t22 *testing.T) {
	lastModified, error := time.Parse("Mon Jan 02 15:04:05 MST 2006", "Thu Mar 03 20:39:25 UTC 2016")

	if error != nil {
		t22.Error(error)
	}
	fmt.Println(lastModified)
}
