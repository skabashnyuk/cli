// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"github.com/docker/docker/pkg/term"
	"github.com/sirupsen/logrus"
	"github.com/skabashnyuk/cli/cli"
	"github.com/skabashnyuk/cli/cli/command"
	"github.com/skabashnyuk/cli/cli/flags"
	commands "github.com/skabashnyuk/cli/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
)

func newCheCommand(CheCli *command.CheCli) *cobra.Command {
	opts := flags.NewCheCliOptions()
	var flags *pflag.FlagSet

	cmd := &cobra.Command{
		Use:              "che [OPTIONS] COMMAND [ARG...]",
		Short:            "Eclipse Che cli",
		SilenceUsage:     true,
		SilenceErrors:    true,
		TraverseChildren: true,
		Args:             noArgs,
	}
	cli.SetupRootCommand(cmd)

	flags = cmd.Flags()
	flags.BoolVarP(&opts.Version, "version", "v", false, "Print version information and quit")
	//flags.StringVar(&opts.ConfigDir, "config", cliconfig.Dir(), "Location of client config files")

	cmd.SetOutput(CheCli.Out())
	commands.AddCommands(cmd, CheCli)

	return cmd
}

func initializeCheCli(CheCli *command.CheCli, flags *pflag.FlagSet) {

}

// visitAll will traverse all commands from the root.
// This is different from the VisitAll of cobra.Command where only parents
// are checked.
func visitAll(root *cobra.Command, fn func(*cobra.Command)) {
	for _, cmd := range root.Commands() {
		visitAll(cmd, fn)
	}
	fn(root)
}

func noArgs(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return nil
	}
	return fmt.Errorf(
		"che: '%s' is not a che command.\nSee 'che --help'", args[0])
}

func main() {
	// Set terminal emulation based on platform as required.
	stdin, stdout, stderr := term.StdStreams()
	logrus.SetOutput(stderr)

	cheCli := command.NewCli(stdin, stdout, stderr)
	cmd := newCheCommand(cheCli)

	if err := cmd.Execute(); err != nil {
		if sterr, ok := err.(cli.StatusError); ok {
			if sterr.Status != "" {
				fmt.Fprintln(stderr, sterr.Status)
			}
			// StatusError should only be used for errors, and all errors should
			// have a non-zero exit status, so never exit with 0
			if sterr.StatusCode == 0 {
				os.Exit(1)
			}
			os.Exit(sterr.StatusCode)
		}
		fmt.Fprintln(stderr, err)
		os.Exit(1)
	}
}

func showVersion() {
	fmt.Printf("Eclipse Che cli version %s, build %s\n", cli.Version, cli.GitCommit)
}
