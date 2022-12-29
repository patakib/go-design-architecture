`
DI is a pattern that allows a component to receive its dependencies (i.e., the objects it depends on) as constructor parameters, rather than creating them directly. This makes it easier to test the component in isolation and to swap out dependencies for different implementations.
`

package main

import (
	"fmt"
	"log"
)

// Logger is an interface for logging messages.
type Logger interface {
	Log(message string)
}

// StdoutLogger is a concrete implementation of Logger that logs messages to stdout.
type StdoutLogger struct{}

// Log logs a message to stdout.
func (l *StdoutLogger) Log(message string) {
	fmt.Println(message)
}

// Service is a component that depends on a logger.
type Service struct {
	logger Logger
}

// NewService creates a new Service with the given logger.
func NewService(logger Logger) *Service {
	return &Service{logger: logger}
}

// DoSomething does something and logs a message.
func (s *Service) DoSomething() {
	s.logger.Log("Doing something...")
}

func main() {
	logger := &StdoutLogger{}
	service := NewService(logger)

	if err := service.DoSomething(); err != nil {
		log.Fatal(err)
	}
}


`
In this example, the Logger interface and the StdoutLogger struct belong to the infrastructure layer, while the Service struct is part of the application's inner layer. The main function is the outer layer, and it depends on the inner layer (Service) to perform some action and log a message.

The Service struct receives its Logger dependency as a constructor parameter, and this dependency is injected into the Service struct by the NewService function. This separation of concerns allows for more maintainable and testable code, as the Service struct can be tested in isolation from the Logger dependency. It also makes it easier to swap out the Logger dependency for a different implementation (e.g., a logger that logs to a file instead of stdout) without affecting the Service struct. `
`
