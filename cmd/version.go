package cmd

import (
	"github.com/docker/cli/templates"
	"github.com/skabashnyuk/cli/cli"
	"github.com/skabashnyuk/cli/cli/command"
	"github.com/spf13/cobra"
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
	BuildTime string `json:",omitempty"`
}

var versionTemplate = `Eclipse Che Cli:
 Version:      {{.Version}}
 Go version:   {{.GoVersion}}
 Git commit:   {{.GitCommit}}
 Built:        {{.BuildTime}}
 OS/Arch:      {{.Os}}/{{.Arch}}`

// NewVersionCommand creates a new cobra.Command for `che version`
func NewVersionCommand(cheCli *command.CheCli) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "version [OPTIONS]",
		Short: "Show the Eclipse Che2 cli version information",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runVersion(cheCli)
		},
	}

	return cmd
}

func runVersion(cheCli *command.CheCli) error {

	templateFormat := versionTemplate

	tmpl, err := templates.Parse(templateFormat)
	if err != nil {
		return cli.StatusError{StatusCode: 64,
			Status: "Template parsing error: " + err.Error()}
	}

	vd := versionInfo{
		Version:   cli.Version,
		GoVersion: runtime.Version(),
		GitCommit: cli.GitCommit,
		BuildTime: cli.BuildTime,
		Arch:      runtime.GOARCH,
		Os:        runtime.GOOS,
	}

	// first we need to make BuildTime more human friendly
	t, errTime := time.Parse(time.RFC3339Nano, vd.BuildTime)
	if errTime == nil {
		vd.BuildTime = t.Format(time.ANSIC)
	}

	if err2 := tmpl.Execute(cheCli.Out(), vd); err2 != nil && err == nil {
		err = err2
	}
	cheCli.Out().Write([]byte{'\n'})
	return err
}
