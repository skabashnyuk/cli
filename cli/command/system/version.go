package system

import (
	"github.com/docker/cli/templates"
	"github.com/skabashnyuk/cli/cli"
	"github.com/spf13/cobra"
	"os"
	"runtime"
	"time"
)

// versionInfo contains version information of both the Client, and Server
type versionInfo struct {
	Version   string
	GitCommit string
	GoVersion string
	Os        string
	Arch      string
	BuildDate string `json:",omitempty"`
}

var versionTemplate = `Eclipse Che Cli:
 Version:      {{.Version}}
 Go version:   {{.GoVersion}}
 Git commit:   {{.GitCommit}}
 Built:        {{.BuildDate}}
 OS/Arch:      {{.Os}}/{{.Arch}}`

// NewVersionCommand creates a new cobra.Command for `che version`
func NewVersionCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "version [OPTIONS]",
		Short: "Show the Eclipse Che2 cli version information",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runVersion()
		},
	}

	return cmd
}

func runVersion() error {

	templateFormat := versionTemplate

	tmpl, err := templates.Parse(templateFormat)
	if err != nil {
		return cli.StatusError{StatusCode: 64,
			Status: "Template parsing error: " + err.Error()}
	}

	vd := versionInfo{
		Version:   cli.GitSummary,
		GoVersion: runtime.Version(),
		GitCommit: cli.GitCommit,
		BuildDate: cli.BuildDate,
		Arch:      runtime.GOARCH,
		Os:        runtime.GOOS,
	}

	// first we need to make BuildTime more human friendly
	t, errTime := time.Parse(time.RFC3339Nano, vd.BuildDate)
	if errTime == nil {
		vd.BuildDate = t.Format(time.ANSIC)
	}

	if err2 := tmpl.Execute(os.Stdout, vd); err2 != nil && err == nil {
		err = err2
	}
	os.Stdout.Write([]byte{'\n'})
	return err
}
