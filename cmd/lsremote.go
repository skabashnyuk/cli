package cmd

import (
	"github.com/blang/semver"
	"github.com/pkg/errors"
	"github.com/skabashnyuk/cli/client"
	"github.com/spf13/cobra"
	"log"
	"os"
	"sort"
	"text/tabwriter"
	"text/template"
	"time"
)

// serveCmd represents the serve command
var lsRemoteCmd = &cobra.Command{
	Use:   "ls-remote",
	Short: "List remote Eclipse Che versions available for install",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runLsRemote()
	},
}

const tagsTemplate = `{{"	VERSION	DATE	PACKAGE	Size"}}{{range .}}
	{{.Version}}	{{.Date.Format "Jan 02, 2006 15:04:05"}}	{{.Package}}	{{.Size}}{{end}}
`

type RemoteChePackage struct {
	Version semver.Version
	Date    time.Time
	Package string
	Size    int64
}

func init() {
	rootCmd.AddCommand(lsRemoteCmd)
}

func runLsRemote() error {

	tags, error := client.GetTags("eclipse", "che-server")
	if error != nil {
		return errors.Errorf("Fail to get dockerhub tags : %s", error.Error())
	}
	metadata, error := client.GetMetadata("http://maven.codenvycorp.com/content/groups/public", "org.eclipse.che", "assembly-main")
	if error != nil {
		return errors.Errorf("Fail to get  maven.codenvycorp.com artifacts: %s", error.Error())
	}

	var data []RemoteChePackage

	for i := range tags {
		// assuming little endian
		tagVersion, error := semver.ParseTolerant(tags[i].Name)
		if error != nil {
			//fmt.Printf("Not able to parse version %s reason %s\n", tags[i].Name, error)
		} else {
			data = append(data, RemoteChePackage{
				Version: tagVersion,
				Date:    tags[i].LastUpdated,
				Package: "Docker",
				Size:    tags[i].FullSize,
			})
		}
	}
	for i := range metadata.Versioning.Versions {
		// assuming little endian
		versionString := metadata.Versioning.Versions[i]
		mavenArtifactVersion, _ := semver.ParseTolerant(versionString)
		//mavenMethaData, error := client.GetArtifactMetadata("http://maven.codenvycorp.com/content/groups/public", "org.eclipse.che", "assembly-main", versionString, "tar.gz")
		if error != nil {
			return errors.Errorf("Fail to get artifact version : %s", error.Error())
		}
		data = append(data, RemoteChePackage{
			Version: mavenArtifactVersion,
			Date:    time.Now(),
			Size:    0,
			Package: "Maven",
		})
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Version.LE(data[j].Version)
	})
	t := template.New("test")
	t, _ = t.Parse(tagsTemplate)
	w := tabwriter.NewWriter(os.Stdout, 2, 2, 2, ' ', 0)
	//err = t.Execute(os.Stdout, tags)
	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
	w.Flush()
	return nil
}
