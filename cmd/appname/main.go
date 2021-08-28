package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/CodingForFunAndProfit/appname/network"
)

func main() {
	defer func() {
		fmt.Println("Ending main")
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	var routineWaitgroup sync.WaitGroup

	network := network.New()
	routineWaitgroup.Add(1)
	go func() { defer routineWaitgroup.Done(); network.Start() }()

	sig := <-c
	fmt.Printf("Got %s signal. Shutting down...\n", sig)

	network.Stop()
	routineWaitgroup.Wait()
}
