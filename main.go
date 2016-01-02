package main

import (
	"fmt"
	"github.com/wuyq101/livermore/cmd"
	"os"
	"runtime"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
