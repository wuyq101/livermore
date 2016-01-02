package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Verbose bool

var RootCmd = &cobra.Command{
	Use:   "livermore",
	Short: "Livermore is a command line tool for stock data fetch and display",
	Long: `Livermore is written in Go. 
	       More detail pls visit : https://github.com/wuyq101/livermore
		`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("This is test from cobra")
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}
