package main

import (
	"sync"
)

type Component interface {
	Start() error
	Stop() error
}

type DepNode struct {
	Component
	Children []*DepNode
}

func (n DepNode) Start(errhandler func(err error)) {
	for _, c := range n.Children {
		go c.Start(errhandler)
	}

	err := n.Component.Start()
	if err != nil {
		errhandler(err)
	}
}

func (n DepNode) Stop(errhandler func(err error)) {
	n.Component.Stop()

	var wg sync.WaitGroup
	wg.Add(len(n.Children))
	for _, c := range n.Children {
		go func(c *DepNode) {
			defer wg.Done()
			c.Stop(errhandler)
		}(c)
	}
	wg.Wait()
}
