package cmd

import (
	"fmt"

	"github.com/pantheonproject/gaia/buildvars"
	"github.com/spf13/cobra"
)

func (r *rootCmd) setupVersionCmd() error {
	cmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Display the version of the software",
		Long:    "Reads and displays the build-time variable Version from the buildvars package, which should contain the Semantic Version of this executable.",
		Run: func(cmd *cobra.Command, args []string) {
			r.getLogger().Info(fmt.Sprintf("Software Version: %s", buildvars.Version))
		},
	}
	r.cmd.AddCommand(cmd)
	return nil
}
