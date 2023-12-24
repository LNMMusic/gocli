package gocli

import "errors"

var (
	// ErrCommandHandlerNotFound is the error that returns when the command is not found.
	ErrCommandHandlerNotFound = errors.New("command not found")
	// ErrCommandManagerNotFound is the error that returns when the command manager is not found.
	ErrCommandManagerNotFound = errors.New("command manager not found")
)

// Commander is the interface that wraps the methods that a command registry must implement.
type Commander interface {
	// FindHandler is the method that finds a handler by name.
	FindHandler(commandName string, commandChain ...string) (h CommandHandler, err error)

	// AddCommand is the method that adds a command to the registry.
	AddCommand(command Command) (err error)

	// Group is the method that groups a list of commands.
	Group(name string, description string) (cm Commander)
}