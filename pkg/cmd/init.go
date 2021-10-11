package cmd

import (
	"fmt"

	cli "github.com/spf13/cobra"
	"github.com/4thel00z/templatectl/pkg/templatectl"
	"github.com/4thel00z/templatectl/pkg/util/exit"
	"github.com/4thel00z/templatectl/pkg/util/osutil"
)

// Init contains the cli-command for initializing the local template
// registry in case it's not initialized.
var Init = &cli.Command{
	Use:   "init",
	Short: "Initialize directories required by templatectl (By default done by installation script)",
	Run: func(c *cli.Command, _ []string) {
		// Check if .config/templatectl exists
		if exists, err := osutil.DirExists(templatectl.Configuration.TemplateDirPath); exists {
			if shouldRecreate := GetBoolFlag(c, "force"); !shouldRecreate {
				exit.GoodEnough("template registry is already initialized use -f to reinitialize")
			}
		} else if err != nil {
			exit.Error(fmt.Errorf("init: %s", err))
		}

		if err := osutil.CreateDirs(templatectl.Configuration.TemplateDirPath); err != nil {
			exit.Error(err)
		}

		exit.OK("Initialization complete")
	},
}
