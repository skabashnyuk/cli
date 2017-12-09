package cmd

import (
	"github.com/pkg/errors"
	"github.com/skabashnyuk/cli/client"
	"github.com/spf13/cobra"
	"log"
	"os"
	"text/tabwriter"
	"text/template"
)

// serveCmd represents the serve command
var lsRemoteCmd = &cobra.Command{
	Use:   "ls-remote",
	Short: "List remote Eclipse Che versions available for install",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runLsRemote()
	},
}

const tagsTemplate = `{{range .}}
	{{.Name}}	{{.LastUpdated.Format "Jan 02, 2006 15:04:05"}}{{end}}
`

func init() {
	rootCmd.AddCommand(lsRemoteCmd)
}

func runLsRemote() error {

	tags, error := client.GetTags("eclipse", "che-server")
	if error != nil {
		return errors.Errorf("Fail to get dockerhub tags : %s", error.Error())
	}

	t := template.New("test")
	t, _ = t.Parse(tagsTemplate)
	w := tabwriter.NewWriter(os.Stdout, 2, 2, 2, ' ', 0)
	//err = t.Execute(os.Stdout, tags)
	if err := t.Execute(w, tags); err != nil {
		log.Fatal(err)
	}
	w.Flush()
	return nil
	//if err := t.Execute(os.Stdout, tags); err != nil {
	//	log.Fatal(err)
	//}
	//w.Flush()
}
