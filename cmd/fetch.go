package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	fetchCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
	RootCmd.AddCommand(fetchCmd)
}

var Source string

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch stock data by code",
	Long: `Fetch stock data by code
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start to fetch stock data ...")
		fmt.Printf("args %q \n", args)
		fmt.Println("Version ", Verbose)
		fmt.Println("Source", Source)
	},
}
