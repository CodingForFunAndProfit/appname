package network

import (
	"fmt"
	"sync"
)

type network struct {
	closed chan struct{}
}

func New() *network {
	return &network{
		closed: make(chan struct{}),
	}
}
func (net *network) Start() {

	var routineWaitgroup sync.WaitGroup
	client := NewClient()
	routineWaitgroup.Add(1)
	go func() { defer routineWaitgroup.Done(); client.Start() }()

	// first possibility when this go routine doesn't have to do
	// something
	<-net.closed

	fmt.Println("Shutting down network.")
	client.Stop()
	routineWaitgroup.Wait()

	/*
		// second possibility when this go routine has to do stuff
		DONE:
		for {
			fmt.Println("I'm doing stuff")
			time.Sleep(time.Millisecond * 1000)

			select {
			case <-net.closed:
				fmt.Println("Shutting down network.")
				client.Stop()
				routineWaitgroup.Wait()
				break DONE
			default:
				continue
			}
		}
	*/
}

func (net *network) Stop() {
	close(net.closed)
}
