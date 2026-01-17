# Mk

**Mk** is a task runner. You can define your own tasks in `Mkfile` and tell `mk` to run them.

## Installation

```
go install github.com/kostya-zero/mk
```

or download it from releases page for your system.

## Configuration

**Mk** uses `Mkfile` to define tasks and commands in them.
It has syntax similar to `Makefile`.
Here is a quick example:

```make
# The default task. Will be executed if no step
# is provided as argument to mk.
default:
    go run .

# You can define as many tasks as you want!
tidy:
    go mod tidy

# You can run tasks before the specified task.
# For example, run tidy task before build.
build: tidy
    go build . -o mk
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
