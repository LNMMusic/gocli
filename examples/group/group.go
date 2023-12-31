package main

import (
	"fmt"

	"github.com/LNMMusic/gocli"
	"github.com/LNMMusic/optional"
)

func main() {
	// env
	// ...

	// cli
	// - config
	cli := gocli.NewCLI(
		gocli.NewParserDefault(optional.None[gocli.ConfigParserDefault]()),
		gocli.NewCommanderManager("basic", "basic example"),
	)
	// - group
	group := cli.Group("group", "this groups represents 2 commands, ping and hello")
	// - add commands
	group.AddCommand(gocli.Command{
		Name: "ping",
		Description: "this command prints pong",
		Handler: func(i gocli.Input) error {
			// print pong
			fmt.Println("pong")
			return nil
		},
	})
	group.AddCommand(gocli.Command{
		Name: "hello",
		Description: "this command prints hello world and show info about the input",
		Handler: func(i gocli.Input) error {
			// print hello world
			fmt.Println("hello world")
			// show info
			fmt.Printf("-command: %s\n-chain: %v\n-flags: %v\n-options: %v\n", i.CommandInput.Command, i.CommandInput.Chain, i.Flags, i.Options)
			return nil
		},
	})
	// - run
	if err := cli.Run(); err != nil {
		fmt.Println(err)
		return
	}
}