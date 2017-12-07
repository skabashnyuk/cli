package main

import (
	"github.com/skabashnyuk/cli/cmd"
)

//func newCheCommand() *cobra.Command {
//	opts := flags.NewCheCliOptions()
//	var flags *pflag.FlagSet
//
//	cmd := &cobra.Command{
//		Use:              "che [OPTIONS] COMMAND [ARG...]",
//		Short:            "Eclipse Che cli",
//		SilenceUsage:     true,
//		SilenceErrors:    true,
//		TraverseChildren: true,
//		Args:             noArgs,
//		Run: func(cmd *cobra.Command, args []string) {
//			if opts.Version {
//				showVersion()
//			}
//		},
//	}
//	cli.SetupRootCommand(cmd)
//
//	flags = cmd.Flags()
//	flags.BoolVarP(&opts.Version, "version", "v", false, "Print version information and quit")
//	commands.AddCommands(cmd)
//
//	return cmd
//}
//
//func noArgs(cmd *cobra.Command, args []string) error {
//	if len(args) == 0 {
//		return nil
//	}
//	return fmt.Errorf(
//		"che: '%s' is not a che command.\nSee 'che --help'", args[0])
//}

func main() {
	cmd.Execute()
	//cmd := newCheCommand()
	//
	//if err := cmd.Execute(); err != nil {
	//	if sterr, ok := err.(cli.StatusError); ok {
	//		if sterr.Status != "" {
	//			fmt.Fprintln(os.Stderr, sterr.Status)
	//		}
	//		// StatusError should only be used for errors, and all errors should
	//		// have a non-zero exit status, so never exit with 0
	//		if sterr.StatusCode == 0 {
	//			os.Exit(1)
	//		}
	//		os.Exit(sterr.StatusCode)
	//	}
	//	fmt.Fprintln(os.Stderr, err)
	//	os.Exit(1)
	//}
}

//func showVersion() {
//	fmt.Printf("Eclipse Che cli version %s, build %s\n", cli.Version, cli.GitCommit)
//}
