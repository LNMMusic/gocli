# GoCLI - Command Line Interface Toolkit

## Introduction
GoCLI is a versatile toolkit for building command-line interfaces (CLI) in Go. It simplifies parsing command-line arguments and executing associated commands. Whether you're building a simple tool or a complex application, GoCLI offers a structured way to handle user inputs efficiently.

## Workflow
1. **Argument Parsing**: GoCLI leverages `os.Args` to receive command-line arguments.
2. **Argument Structure**:
   - **Commands**: One or more actions to be performed.
   - **Flags**: Prefixed with `--` or `-` followed by a key and a value (alphanumeric).
   - **Options**: Prefixed with `-` followed by uppercase words.

   Format: `app.exe command1 command2 command3 --flag1 value1 -flag2 value2 -OPTION1 -OPTION2`

## Fundamental Blocks

```go
// CLI struct wraps the commands, parser, and commander.
type CLI struct {
    Parser    // Parses CLI input.
    Commander // Manages command execution.
}
```

- **Parser**: Validates each part of the args (commands, flags, options) and structures them into an `Input` object.
- **Commander**: Facilitates adding commands with a name and a handler. Allows nested commands or groups.

## Usage

### Basic Example

```go
package main

import (
    "fmt"
    "github.com/LNMMusic/gocli"
    "github.com/LNMMusic/optional"
)

func main() {
    // Initialize CLI with default parser and commander
    cli := gocli.NewCLI(
        gocli.NewParserDefault(optional.None[gocli.ConfigParserDefault]()),
        gocli.NewCommanderManager("basic", "basic example"),
    )

    // Add command 'hello'
    cli.AddCommand(gocli.Command{
        Name: "hello",
        Description: "Prints hello world and shows input info",
        Handler: func(i gocli.Input) error {
            fmt.Println("hello world")
            fmt.Printf("-command: %s\n-chain: %v\n-flags: %v\n-options: %v\n", i.CommandInput.Command, i.CommandInput.Chain, i.Flags, i.Options)
            return nil
        },
    })

    // Execute CLI
    if err := cli.Run(); err != nil {
        fmt.Println(err)
        return
    }
}
```

### Group Commands Example

```go
package main

import (
    "fmt"
    "github.com/LNMMusic/gocli"
    "github.com/LNMMusic/optional"
)

func main() {
    // Initialize CLI
    cli := gocli.NewCLI(
        gocli.NewParserDefault(optional.None[gocli.ConfigParserDefault]()),
        gocli.NewCommanderManager("group", "example with groups"),
    )

    // Create a command group
    group := cli.Group("group", "Represents 'ping' and 'hello' commands")

    // Add 'ping' command to the group
    group.AddCommand(gocli.Command{
        Name: "ping",
        Description: "Prints pong",
        Handler: func(i gocli.Input) error {
            fmt.Println("pong")
            return nil
        },
    })

    // Add 'hello' command to the group
    group.AddCommand(gocli.Command{
        Name: "hello",
        Description: "Prints hello world and shows input info",
        Handler: func(i gocli.Input) error {
            fmt.Println("hello world")
            fmt.Printf("-command: %s\n-chain: %v\n-flags: %v\n-options: %v\n", i.CommandInput.Command, i.CommandInput.Chain, i.Flags, i.Options)
            return nil
        },
    })

    // Execute CLI
    if err := cli.Run(); err != nil {
        fmt.Println(err)
        return
    }
}
```

## Conclusion
GoCLI is designed to make CLI development in Go more intuitive and structured. By abstracting the complexity of argument parsing and command handling, it allows developers to focus on implementing the core logic of their applications.