`
The observer pattern is a behavioral pattern that allows an object (the subject) to notify a set of objects (the observers) when its state changes.
`

package main

import "fmt"

// Observer is an interface for observers.
type Observer interface {
	Update(subject *Subject)
}

// ConcreteObserverA is a concrete implementation of Observer.
type ConcreteObserverA struct{}

// Update updates the observer based on the state of the subject.
func (o *ConcreteObserverA) Update(subject *Subject) {
	fmt.Println("ConcreteObserverA:", subject.GetState())
}

// ConcreteObserverB is a concrete implementation of Observer.
type ConcreteObserverB struct{}

// Update updates the observer based on the state of the subject.
func (o *ConcreteObserverB) Update(subject *Subject) {
	fmt.Println("ConcreteObserverB:", subject.GetState())
}

// Subject is an interface for subjects.
type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify()
	GetState() string
	SetState(state string)
}

// ConcreteSubject is a concrete implementation of Subject.
type ConcreteSubject struct {
	observers []Observer
	state     string
}

// Attach attaches an observer to the subject.
func (s *ConcreteSubject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

// Detach detaches an observer from the subject.
func (s *ConcreteSubject) Detach(observer Observer) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		
func main() {
	observerA := &ConcreteObserverA{}
	observerB := &ConcreteObserverB{}
	subject := &ConcreteSubject{}

	subject.Attach(observerA)
	subject.Attach(observerB)
	subject.SetState("Hello")
	subject.Notify()

	subject.Detach(observerA)
	subject.SetState("World")
	subject.Notify()
}

`
In this example, the Observer interface and the ConcreteObserverA and ConcreteObserverB structs form the application's inner layer, while the Subject interface and the ConcreteSubject struct are part of the infrastructure layer. The main function is the outer layer, and it depends on the inner layer (Observer) to update itself based on the state of the subject.

The Subject interface defines an interface for attaching and detaching observers, and for notifying them when the subject's state changes. The ConcreteSubject struct implements this interface and maintains a list of attached observers. The Attach, Detach, and Notify methods allow the subject to add and remove observers, and to notify them when the subject's state changes.

This pattern allows the subject to notify its observers when its state changes, without the subject knowing the concrete implementation of the observers. This separation of concerns allows for more maintainable and testable code, as the inner layer can be tested in isolation from the infrastructure layer.
`
