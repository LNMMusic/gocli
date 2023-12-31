package gocli

import "errors"

var (
	// ErrInvalidArgs is the error that occurs when the command is invalid.
	ErrInvalidArgs = errors.New("invalid args")
	// ErrInvalidCommands is the error that occurs when the command chain is invalid.
	ErrInvalidCommands = errors.New("invalid commands")
)

type CommandInput struct {
	// Chain is the chain of the command.
	Chain []string
	// Command is the command.
	Command string
}

// Input is the struct that wraps the input of the command.
type Input struct {
	// CommandInput is the input of the command.
	CommandInput CommandInput
	// Flags are the arguments of the command.
	Flags map[string]any
	// Options are the options of the command.
	Options map[string]int
}

// Parser is the interface that wraps the basic Parse method.
type Parser interface {
	Parse(args string) (i Input, err error)
}