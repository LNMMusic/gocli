package gocli

import (
	"os"
	"strings"
)

// NewCLI is the function that returns a new CLI.
func NewCLI(p Parser, c Commander) (r CLI) {
	r = CLI{
		Parser: p,
		Commander: c,
	}
	return
}

// CLI is the struct that wraps the commands.
type CLI struct {
	// Parser is the parser of the CLI.
	// it returns the Input of the command handler.
	Parser
	// Commander is the commander of the CLI.
	// helps finding the command handler from the input.
	Commander
}

// Run is the method that runs the CLI.
func (c CLI) Run() (err error) {
	// fetch args
	// - os.Stdin
	args := strings.Join(os.Args[1:], " ")
	
	// parse the input
	input, err := c.Parser.Parse(args)
	if err != nil {
		return
	}
	
	// find the command handler
	handler, err := c.Commander.FindHandler(input.CommandInput.Command, input.CommandInput.Chain...)
	if err != nil {
		return
	}
	
	// run the command handler
	err = handler(input)
	if err != nil {
		return err
	}
	
	return
}