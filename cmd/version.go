package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionInfo = "hoard version 0.0.1 (c) 2016 Thornton Prime"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show the current version",
	Long: `Show the current version and other information about the current
hoard environment.`,
	Run: showVersion,
}

func showVersion(cmd *cobra.Command, args []string) {
	fmt.Println(versionInfo)

}
func init() {
	HoardCmd.AddCommand(versionCmd)
}
