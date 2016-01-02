package main

import (
	"fmt"
	"github.com/wuyq101/livermore/cmd"
	"github.com/wuyq101/livermore/config"
	"github.com/wuyq101/livermore/model"
	"github.com/wuyq101/livermore/workflow"
	"os"
	"runtime"
)

func main() {
	fmt.Println("Hi, what's up ?")
	fmt.Println("Start to init ...")

	fmt.Println("Start to init config ...")
	config.Instance()

	fmt.Println("Start to init model ...")
	m := model.NewModel()

	fmt.Println("Start to init workflow ...")
	w := workflow.NewWorkFlow(m)

	w.Model()
	fmt.Println("Init finished ...")

	fmt.Println("Start to insert a stock")
	w.Model().InsertStock(&model.Stock{Name: "test"})

	fmt.Println("Start to init cobra ...")
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
