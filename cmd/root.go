package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string
var ShowVersion bool

var rootCmd = &cobra.Command{
	Use:              "che [OPTIONS] COMMAND [ARG...]",
	Short:            "Eclipse Che cli",
	SilenceUsage:     true,
	SilenceErrors:    true,
	TraverseChildren: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if ShowVersion {
			showVersion()
			return nil
		}
		cmd.SetOutput(os.Stderr)
		cmd.HelpFunc()(cmd, args)
		return nil

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.AddTemplateFunc("hasSubCommands", hasSubCommands)
	cobra.AddTemplateFunc("operationSubCommands", operationSubCommands)
	cobra.AddTemplateFunc("wrappedFlagUsages", wrappedFlagUsages)

	rootCmd.SetUsageTemplate(usageTemplate)
	rootCmd.SetHelpTemplate(helpTemplate)
	rootCmd.SetFlagErrorFunc(FlagErrorFunc)
	rootCmd.SetHelpCommand(helpCommand)

	rootCmd.PersistentFlags().BoolP("help", "h", false, "Print usage")
	rootCmd.PersistentFlags().MarkShorthandDeprecated("help", "please use --help")
	rootCmd.Flags().BoolVarP(&ShowVersion, "version", "v", false, "Print version information and quit")

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.che/config.yaml)")

}

// FlagErrorFunc prints an error message which matches the format of the
// docker/docker/cli error messages
func FlagErrorFunc(cmd *cobra.Command, err error) error {
	if err == nil {
		return nil
	}

	usage := ""
	if cmd.HasSubCommands() {
		usage = "\n\n" + cmd.UsageString()
	}

	return errors.Errorf("%s\nSee '%s --help'.%s", err, cmd.CommandPath(), usage)
}

func hasSubCommands(cmd *cobra.Command) bool {
	return len(operationSubCommands(cmd)) > 0
}

func operationSubCommands(cmd *cobra.Command) []*cobra.Command {
	cmds := []*cobra.Command{}
	for _, sub := range cmd.Commands() {
		if sub.IsAvailableCommand() && !sub.HasSubCommands() {
			cmds = append(cmds, sub)
		}
	}
	return cmds
}

func wrappedFlagUsages(cmd *cobra.Command) string {
	width := 80
	//TODO detect with
	//if ws, err := term.GetWinsize(0); err == nil {
	//	width = int(ws.Width)
	//}
	return cmd.Flags().FlagUsagesWrapped(width - 1)
}

var usageTemplate = `Usage:
{{- if not .HasSubCommands}}	{{.UseLine}}{{end}}
{{- if .HasSubCommands}}	{{ .CommandPath}} COMMAND{{end}}
{{ .Short | trim }}
{{- if .HasFlags}}
Options:
{{ wrappedFlagUsages . | trimRightSpace}}
{{- end}}
{{- if hasSubCommands .}}
Commands:
{{- range operationSubCommands . }}
  {{rpad .Name .NamePadding }} {{.Short}}
{{- end}}
{{- end}}
{{- if .HasSubCommands }}
Run '{{.CommandPath}} COMMAND --help' for more information on a command.
{{- end}}
`

var helpTemplate = `
{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}`

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName("/.che/config.yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func showVersion() {
	fmt.Printf("Eclipse Che cli version %s, build %s\n", Version, GitCommit)
}
