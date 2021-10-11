package cmd

import (
	"fmt"

	cli "github.com/spf13/cobra"

	"github.com/4thel00z/templatectl/pkg/templatectl"
	"github.com/4thel00z/templatectl/pkg/util/tlog"
	"github.com/4thel00z/templatectl/pkg/util/validate"
)

// Version contains the cli-command for printing the current version of the tool.
var Version = &cli.Command{
	Use:   "version",
	Short: "Show the templatectl version information",
	Run: func(c *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{})

		shouldntPrettify := GetBoolFlag(c, "dont-prettify")
		if shouldntPrettify {
			fmt.Println(templatectl.Version)
		} else {
			tlog.Info(fmt.Sprint("Current version is ", templatectl.Version))
		}
	},
}
