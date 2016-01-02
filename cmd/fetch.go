package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wuyq101/livermore/workflow"
)

var FetchAll bool

func init() {
	RootCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().BoolVarP(&FetchAll, "all", "a", false, "Fetch all stock info current in database.")
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch stock data by code",
	Long: `Fetch stock data by code
          stock code: sh600036, sz000166
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start to fetch stock data ...")
		w := workflow.Instance()
		if FetchAll {
			w.FetchAllStockInfo()
			return
		}
		fmt.Printf("args %q \n", args)
		if args == nil || len(args) <= 0 {
			fmt.Println("please input stock code, like sh600036 sz000166")
		}
		w.FetchStockInfo(args)
	},
}
