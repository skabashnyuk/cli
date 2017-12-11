package client

import (
	"testing"
)

func TestShouldBeAbleToGetListOfDockerHubTags(t22 *testing.T) {

	tags, error := GetTags("eclipse", "che-server")
	if error != nil {
		t22.Error(error)
	}
	if len(tags) < 1 {
		t22.Error("Tags list should not be empty")
	}
}
