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

//func TestTimeConsuming(t22 *testing.T) {
//
//	tags, error := DockerHub{}.GetTags("eclipse", "che-server")
//	if error != nil {
//		t22.Error(error)
//	}
//	//if testing.Short() {
//	//	t.Skip("skipping test in short mode.")
//	//}
//	fmt.Printf("%v", tags)
//
//	const tagsTemplate = `{{range .}}
//{{.Name}}	{{.LastUpdated.Format "Jan 02, 2006 15:04:05"}}{{end}}`
//
//	t := template.New("test")
//	t, _ = t.Parse(tagsTemplate)
//	w := tabwriter.NewWriter(os.Stdout, 2, 2, 2, ' ', 0)
//	//err = t.Execute(os.Stdout, tags)
//	if err := t.Execute(w, tags); err != nil {
//		log.Fatal(err)
//	}
//	w.Flush()
//
//	//if err := t.Execute(os.Stdout, tags); err != nil {
//	//	log.Fatal(err)
//	//}
//	//w.Flush()
//
//}
