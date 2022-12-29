package main

import (
	"fmt"
	"sync"
)

// Filter is an interface for filters.
type Filter interface {
	Process(data chan interface{}, output chan interface{})
}

// ConcreteFilterA is a concrete implementation of Filter.
type ConcreteFilterA struct{}

// Process processes the data and sends the result to the output channel.
func (f *ConcreteFilterA) Process(data chan interface{}, output chan interface{}) {
	for d := range data {
		output <- d.(int) * 2
	}
	close(output)
}

// ConcreteFilterB is a concrete implementation of Filter.
type ConcreteFilterB struct{}

// Process processes the data and sends the result to the output channel.
func (f *ConcreteFilterB) Process(data chan interface{}, output chan interface{}) {
	for d := range data {
		output <- d.(int) + 1
	}
	close(output)
}

func main() {
	data := make(chan interface{})
	outputA := make(chan interface{})
	outputB := make(chan interface{})

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			data <- i
		}
		close(data)
	}()

	go func() {
		defer wg.Done()
		filterA := &ConcreteFilterA{}
		filterA.Process(data, outputA)
	}()

	go func() {
		defer wg.Done()
		filterB := &ConcreteFilterB{}
		filterB.Process(outputA, outputB)
	}()

	go func() {
		defer wg.Done()
		for result := range outputB {
			fmt.Println(result)
		}
	}()

	wg.Wait()
}


`
In this example, the Filter interface defines an interface for filters, and the ConcreteFilterA and ConcreteFilterB structs are concrete implementations of this interface. The Process method of each filter receives data from a channel and sends the result to another channel.

The main function creates three goroutines: one that generates data, one that processes the data with ConcreteFilterA, and one that processes the data with ConcreteFilterB. The goroutines communicate through channels, which act as pipes that connect the filters.

This pattern allows the filters to be reused and composed in different ways, and it allows the data to be processed in parallel. This can improve the performance and scalability of the system.
`
