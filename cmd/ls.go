package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"text/tabwriter"
	"text/template"
)

// serveCmd represents the serve command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List installed Eclipse Che versions",
	Run: func(cmd *cobra.Command, args []string) {
		runLs()
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}

const listCheTemplate = `{{"NAME	VERSION	UPDATED	STATUS	INFRASTRUCTURE"}}{{range .}}
	{{.Version}}	{{.Date.Format "Jan 02, 2006 15:04:05"}}	{{.Package}}	{{.Size}}{{end}}
`

func runLs() error {
	t := template.New("ls")
	t, _ = t.Parse(listCheTemplate)
	w := tabwriter.NewWriter(os.Stdout, 2, 2, 2, ' ', 0)
	//err = t.Execute(os.Stdout, tags)
	if err := t.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
	w.Flush()
	return nil
}
