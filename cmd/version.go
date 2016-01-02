package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version numbre of Livermore",
	Long:  "All software has versions. This is Livermore's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Livermore v0.1 --head")
	},
}
