# manager vending machine

manager vending machine is a vending system that is centrally managed by a Grpc server written in Golang. It offers easy
setup and maintenance, fast and reliable communication, and efficient management of vending operations.

### run vending machine

```bash
make vendingd && ./vendingd machine
```

### test and linter

before create pull request make sure you run this command

```bash
make lint-fix && make lint-get && make lint
make test
make race
```
