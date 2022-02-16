package instal

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/MakeNowJust/heredoc"
	"github.com/abdfnx/instal/cmd/factory"
	"github.com/abdfnx/instal/core/options"
	"github.com/abdfnx/instal/internal/tui"
	"github.com/abdfnx/instal/core/installer"
)

// Execute start the CLI
func Execute(f *factory.Factory, version string, buildDate string) *cobra.Command {
	const desc = `🛰️ Install any binary app from a script URL.`

	opts := options.InstalOptions{
		Shell: "",
		IsHidden: false,
		URL: "",
	}

	// Root command
	var rootCmd = &cobra.Command{
		Use:   "instal <subcommand> [flags]",
		Short:  desc,
		Long: desc,
		// instal args: if there is no args, it will execute tui
		Args: cobra.ArbitraryArgs,
		SilenceErrors: true,
		Example: heredoc.Doc(`
			# Open Resto UI
			instal

			# Install binary app from script URL and run it.
			instal <SCRIPT_URL>
		`),
		Annotations: map[string]string{
			"help:tellus": heredoc.Doc(`
				Open an issue at https://github.com/abdfnx/instal/issues
			`),
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.URL = args[0]
				return installer.RunInstal(&opts, true, "", "", "")
			} else {
				tui.Instal()
			}

			return nil
		},
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Aliases: []string{"ver"},
		Short: "Print the version of your instal binary.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("instal version " + version + " " + buildDate)
		},
	}

	rootCmd.SetOut(f.IOStreams.Out)
	rootCmd.SetErr(f.IOStreams.ErrOut)

	cs := f.IOStreams.ColorScheme()

	helpHelper := func(command *cobra.Command, args []string) {
		rootHelpFunc(cs, command, args)
	}

	rootCmd.PersistentFlags().Bool("help", false, "Help for instal")
	rootCmd.SetHelpFunc(helpHelper)
	rootCmd.SetUsageFunc(rootUsageFunc)
	rootCmd.SetFlagErrorFunc(rootFlagErrorFunc)

	p := "bash"

	if runtime.GOOS == "windows" {
		p = "powershell"
	}

	rootCmd.Flags().StringVarP(&opts.Shell, "shell", "s", "", "shell to use (Default: " + p + ")")
	rootCmd.Flags().BoolVarP(&opts.IsHidden, "hidden", "H", false, "hide the output")

	rootCmd.AddCommand(versionCmd)

	return rootCmd
}
