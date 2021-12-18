package common

import (
	"fmt"
	"github.com/spf13/cobra"
	goVersion "go.hein.dev/go-version"
)

func AddVersionCmd(cmd *cobra.Command, versionInfoProvider func() *goVersion.Info) {
	shortened := false
	output := "yaml"
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Version",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			versionInfo := versionInfoProvider()
			resp := goVersion.FuncWithOutput(shortened, versionInfo.Version, versionInfo.Commit, versionInfo.Date, output)
			fmt.Print(resp)
			return nil
		},
	}
	versionCmd.Flags().BoolVarP(&shortened, "short", "s", false, "Print just the version number.")
	versionCmd.Flags().StringVarP(&output, "output", "o", "yaml", "Output format. One of 'yaml' or 'json'.")
	cmd.AddCommand(versionCmd)
}
