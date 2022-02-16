package checker

import (
	"fmt"
	"strings"

	"github.com/mgutz/ansi"
	"github.com/abdfnx/looker"
	"github.com/abdfnx/instal/api"
	"github.com/abdfnx/instal/cmd/factory"
)

func Check(buildVersion string) {
	cmdFactory := factory.New()
	stderr := cmdFactory.IOStreams.ErrOut

	latestVersion := api.GetLatest()
	isFromHomebrewTap := isUnderHomebrew()
	isFromUsrBinDir := isUnderUsr()
	isFromGHCLI := isUnderGHCLI()
	isFromAppData := isUnderAppData()

	var command = func() string {
		if isFromHomebrewTap {
			return "brew upgrade instal"
		} else if isFromUsrBinDir || isFromAppData {
			return "instal https://bit.ly/instal-cli"
		} else if isFromGHCLI {
			return "gh extention upgrade instal"
		}

		return ""
	}

	if buildVersion != latestVersion {
		fmt.Fprintf(stderr, "%s %s → %s\n",
		ansi.Color("There's a new version of ", "yellow") + ansi.Color("instal", "cyan") + ansi.Color(" is avalaible:", "yellow"),
		ansi.Color(buildVersion, "cyan"),
		ansi.Color(latestVersion, "cyan"))

		if command() != "" {
			fmt.Fprintf(stderr, ansi.Color("To upgrade, run: %s\n", "yellow"), ansi.Color(command(), "black:white"))
		}
	}
}

var instalExe, _ = looker.LookPath("instal")

func isUnderHomebrew() bool {
	return strings.Contains(instalExe, "brew")
}

func isUnderUsr() bool {
	return strings.Contains(instalExe, "usr")
}

func isUnderAppData() bool {
	return strings.Contains(instalExe, "AppData")
}

func isUnderGHCLI() bool {
	return strings.Contains(instalExe, "gh")
}
