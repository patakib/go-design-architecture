# Design Patterns and Architectural Patterns in Go

## Software Architecture Use Cases

### Layered architecture
Use a layered architecture when you want to separate the different concerns of the system into different layers, and you want to allow each layer to depend only on the layers below it. This pattern is useful for systems with complex business logic that needs to be separated from the infrastructure.
### Microservices architecture
Use a microservices architecture when you want to decompose a large system into a set of small, independently deployable services. This pattern is useful for systems that need to be flexible and scalable, and that need to be able to evolve quickly.
### Event-driven architecture
Use an event-driven architecture when you want to decouple the different components of the system, and you want to allow them to communicate asynchronously through events. This pattern is useful for systems that need to be highly responsive and that need to be able to handle a large volume of requests.
### Pipe and filter architecture
Use a pipe and filter architecture when you want to decompose a complex process into a series of simple, reusable components that are connected by pipes. This pattern is useful for systems that need to process large amounts of data in a parallel and scalable way.

## Design Pattern Use Cases
### Factory
Use the factory pattern when you want to create objects of different types, but you want to abstract the creation process from the client.
### Singleton
Use the singleton pattern when you want to ensure that a class has only one instance, and you want to provide a global access point to that instance.
### Observer
Use the observer pattern when you want to allow an object (the subject) to notify a set of objects (the observers) when its state changes.
### Adapter
Use the adapter pattern when you have an existing class that you cannot modify, but you need to use it in a new context that requires a different interface.
### Decorator
Use the decorator pattern when you want to add new behavior to an object dynamically, without modifying its implementation.
### Command
Use the command pattern when you want to encapsulate a request as an object, and you want to be able to execute the request and undo it.
### Facade
Use the facade pattern when you want to provide a simplified interface to a complex system.
### Proxy
Use the proxy pattern when you want to control access to an object, or when you want to add additional functionality to an object.
