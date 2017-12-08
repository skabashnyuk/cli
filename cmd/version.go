package cmd

import (
	"github.com/docker/cli/templates"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
	"runtime"
	"time"
)

var (
	Version    = "unknown-version"
	GitCommit  = "unknown-commit"
	GitSummary = "unknown-summary"
	GitBranch  = "unknown-branch"
	BuildDate  = "unknown-buildtime"
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

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version [OPTIONS]",
	Short: "Show the Eclipse Che2 cli version information",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runVersion()
	},
}

func runVersion() error {

	templateFormat := versionTemplate

	tmpl, err := templates.Parse(templateFormat)
	if err != nil {
		return errors.Errorf("Template parsing error: %s", err.Error())
	}

	vd := versionInfo{
		Version:   Version,
		GoVersion: runtime.Version(),
		GitCommit: GitCommit,
		BuildDate: BuildDate,
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
