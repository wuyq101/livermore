package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wuyq101/livermore/workflow"
)

var Verbose bool

var RootCmd = &cobra.Command{
	Use:   "livermore",
	Short: "Livermore is a command line tool for stock data fetch and display",
	Long: `Livermore is a command line tool for stock data fetch and display.
		   It was written in Go. 
	       More detail please visit : https://github.com/wuyq101/livermore
		`,
	Run: func(cmd *cobra.Command, args []string) {
		//init work flow
		workflow.Instance()
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}
