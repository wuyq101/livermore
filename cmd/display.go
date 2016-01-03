package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wuyq101/livermore/workflow"
)

func init() {
	RootCmd.AddCommand(displayCmd)
}

var displayCmd = &cobra.Command{
	Use:   "display",
	Short: "display stock by term ui",
	Long:  "Display stock info in terminal ui",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start to display stock info ...")
		w := workflow.Instance()
		if args == nil || len(args) <= 0 {
			fmt.Println("please input stock code, like sh600036 sz000166")
			return
		}
		w.DisplayStockInfo(args)
	},
}
