package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wuyq101/livermore/workflow"
)

func init() {
	RootCmd.AddCommand(showCmd)
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show stock infomation by order",
	Long:  "show stock infomation by order...",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start to show stock info ...")
		w := workflow.Instance()
		if args == nil || len(args) <= 0 {
			fmt.Println("please input stack order, like mf (money flow)")
			return
		}
		w.ShowByMoneyFlow(args)
	},
}
