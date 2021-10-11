package cmd

import (
	"fmt"
	"github.com/4thel00z/templatectl/pkg/util/exit"
	"github.com/4thel00z/templatectl/pkg/util/validate"
	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	cli "github.com/spf13/cobra"
)

const version = "0.2.1"

func upgrade(cmd *cli.Command, verbose bool, previous semver.Version) error {
	if verbose {
		selfupdate.EnableLog()
	}
	latest, err := selfupdate.UpdateSelf(previous, "4thel00z/templatectl")
	if err != nil {
		return err
	}

	if previous.Equals(latest.Version) {
		cmd.Println("Current binary is the latest version", version)
	} else {
		cmd.Println("Update successfully done to version", latest.Version)
		cmd.Println("Release note:\n", latest.ReleaseNotes)
	}
	return nil
}

var Upgrade = &cli.Command{
	Use:   "upgrade",
	Short: "Upgrade templatectl if there is a new version",
	Run: func(cmd *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{})
		verbose := GetBoolFlag(cmd, "verbose")
		previous, err := semver.Parse(version)
		if err != nil {
			exit.Fatal(fmt.Errorf("upgrade: %s\nThis is a developer bug, please open an issue here: https://github.com/4thel00z/templatectl", err))
		}
		err = upgrade(cmd, verbose, previous)
		if err != nil {
			exit.Fatal(fmt.Errorf("upgrade: could not update due to %s", err))
		}
	},
}
