# Mk

**Mk** is a lightweight task runner that helps you automate common development tasks.
Define your tasks once in a `Mkfile` and run them with a simple command.

## Table of Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [Usage](#usage)
- [Configuration](#configuration)
- [License](#license)

## Installation

### Via Go Toolchain

```bash
go install github.com/kostya-zero/mk@latest
```

### Via GitHub Releases

Download the pre-built binary for your system from the [GitHub Releases](https://github.com/kostya-zero/mk/releases) page.

## Quick Start

1. Create a `Mkfile` in your project root:

```make
default:
    echo "Hello from Mk!"
```

2. Run the task:

```bash
mk
```

## Usage

### Running Tasks

If you run `mk` without arguments, it will execute the `default` task.
You can specify which task to run by providing its name as an argument.

```bash
# Run the default task
mk

# Run a specific task
mk build

# Run a task with arguments
mk run --verbose
```

### Command-Line Flags

Flags must be passed as the first argument.
If placed elsewhere, they will be treated as arguments for the task.

| Flag | Description                     |
| ---- | ------------------------------- |
| `-h` | Print help message              |
| `-l` | List all available tasks        |
| `-e` | Print all environment variables |
| `-v` | Print the version of Mk         |

**Examples:**

```bash
# Print help
mk -h

# List all tasks
mk -l

# Show environment variables
mk -e

# Print version
mk -v
```

## Configuration

Tasks and commands are defined in a `Mkfile`, which uses a syntax similar to `Makefile`.

### Basic Syntax

```make
# Task name followed by colon
task-name:
    command to execute
    another command
```

### Environment Variables

Define environment variables at the top of your `Mkfile`:

```make
$CGO_ENABLED=0
$GOOS=linux
$GOARCH=amd64
```

### Task Dependencies

Run other tasks before executing the current task:

```make
# Run 'clean' and 'tidy' before 'build'
build: clean tidy
    go build -o app
```

### Tasks with Arguments

Add an asterisk (`*`) to the task name and use `$*` to access command-line arguments:

```make
run*:
    go run . $*

test*:
    go test $*
```

**Usage:**

```bash
mk run --debug
mk test ./...  -v
```

## License

This project is licensed under the MIT License.
See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome!
Feel free to open issues or submit pull requests on this repository.
