package network

import (
	"fmt"
	"time"
)

type client struct {
	closed chan struct{}
}

func NewClient() *client {
	return &client{
		closed: make(chan struct{}),
	}
}
func (c *client) Start() {
	for {
		fmt.Println("client doing stuff")
		time.Sleep(time.Millisecond * 1000)

		select {
		case <-c.closed:
			fmt.Println("Shutting down client.")
			return
		default:
			continue
		}
	}
}

func (c *client) Stop() {
	close(c.closed)
}
