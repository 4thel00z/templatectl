package main

import (
  "fmt"
  "github.com/4thel00z/templatectl/pkg/templatectl"
  "github.com/4thel00z/templatectl/pkg/cmd"
  "github.com/4thel00z/templatectl/pkg/util/exit"
  "github.com/4thel00z/templatectl/pkg/util/osutil"
)

func main() {
  if exists, err := osutil.DirExists(templatectl.Configuration.TemplateDirPath); ! exists {
    if err := osutil.CreateDirs(templatectl.Configuration.TemplateDirPath); err != nil {
      exit.Error(fmt.Errorf("Tried to initialise your template directory, but it has failed: %s", err))
    }
  } else if err != nil {
    exit.Error(fmt.Errorf("Failed to init: %s", err))
  }

  cmd.Run()
}

