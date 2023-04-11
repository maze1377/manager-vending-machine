# Manager Vending Machines

The Manager Vending Machine is a centrally managed vending system that uses a GRPC server written in Golang. It offers
easy setup and maintenance, fast and reliable communication, and efficient vending operations management.

### Running Vending Machines

To test the vending machine, follow these steps:

1- Compile the code:

```bash
make dependencies
make vendingd
```

2- Run the vending machine (you can change the config-file and run multiple machines):

```bash
./vendingd machine -c  local.vending.1.yaml
./vendingd machine -c  local.vending.2.yaml
```

3- Run the manager to attach to the machine:

```bash
./vendingd manager -c local.sample.yaml
```

4- Send a command by example (uncomment the command type and change it):

```bash
go run ./examples/client.go -b localhost:10000
go run ./examples/client.go -b localhost:10001
```

### Testing and Linting

Before creating a pull request, make sure you run these commands:

```bash
make lint-fix && make lint-get && make lint
make test
make race
```

### Design Patterns and Architecture

The following design patterns and architecture used for managing a distributed vending machine:

1-Microservices Architecture: A microservices architecture can be used to decompose the vending machine system into
small,
independent services that can communicate with each other using APIs. This architecture can improve scalability, fault
tolerance, and maintainability.

2-Command Design Pattern: The Command design pattern can be used to encapsulate actions taken by the vending machine
system in a command object. Each command can represent a specific action, such as adding an item to the inventory,
processing a payment, or dispensing an item. This pattern can provide a flexible and extensible way to manage the
vending machine's behavior.

3-Observer Design Pattern: The Observer design pattern can be used to notify the vending machine system when an event
occurs, such as a payment being processed or an item being dispensed. This pattern can provide a decoupled way to manage
the vending machine's behavior and improve maintainability.

4-State Design Pattern: The state pattern can be used to model the different states that a vending machine can be in,
such
as "idle," "dispensing," and "payment." Each state would be represented by a separate class, and the vending machine
object would transition between states as it performs different operations.

5-Domain-Driven Design: Domain-Driven Design (DDD) can be used to model the vending machine system based on its domain
concepts, such as items, payments, and dispensing. This approach can help to create a clear understanding of the
system's behavior and requirements and can guide the design of the system's architecture and APIs.

6-CQRS and Event Sourcing: Command Query Responsibility Segregation (CQRS) and Event Sourcing can be used to separate
the
write and read operations in the vending machine system. CQRS can be used to handle the write operations, such as adding
items to the inventory, while Event Sourcing can be used to store a log of events that represent changes to the system's
state. This approach can provide a scalable and resilient way to manage the vending machine's state and improve
performance.

7-Load Balancing: Load balancing can be used to distribute the workload across multiple instances of the vending machine
system. This can improve scalability and fault tolerance by ensuring that the system can handle a large number of
requests and can recover from failures without affecting the user experience.

8-GRPC Service: GRPC is a communication framework that allows vending machines to communicate with a central server or
set
of servers. The GRPC service would define the messages and operations that the server can perform on the vending
devices, such as dispensing products and updating inventory.

These design patterns and architecture can provide a flexible and maintainable way to manage the complex stateful
systems of distributed vending machines.
