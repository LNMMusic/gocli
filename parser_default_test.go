package gocli_test

import (
	"testing"

	"github.com/LNMMusic/gocli"
	"github.com/LNMMusic/optional"
	"github.com/stretchr/testify/require"
)

// TestParserDefault_ParseCommand tests the method ParseCommand of the ParserDefault type.
func TestParserDefault_ParseCommand(t *testing.T) {
	t.Run("success - case 01: command is parsed correctly - one command", func(t *testing.T) {
		// arrange
		// - parser: default
		ps := gocli.NewParserDefault(optional.None[gocli.ConfigParserDefault]())

		// act
		args := "app cmd1 --flag1 value1 -flag2 value2 -O1 -O2"
		ci, err := ps.ParseCommands(args)

		// assert
		require.NoError(t, err)
		require.Equal(t, gocli.CommandInput{
			Name: "app",
			Chain: []string{},
			Command: "cmd1",
		}, ci)
	})

	t.Run("success - case 02: command is parsed correctly - multiple commands", func(t *testing.T) {
		// arrange
		// - parser: default
		ps := gocli.NewParserDefault(optional.None[gocli.ConfigParserDefault]())

		// act
		args := "app cmd1 cmd2 cmd3 --flag1 value1 -flag2 value2 -O1 -O2"
		ci, err := ps.ParseCommands(args)

		// assert
		require.NoError(t, err)
		require.Equal(t, gocli.CommandInput{
			Name: "app",
			Chain: []string{"cmd1", "cmd2"},
			Command: "cmd3",
		}, ci)
	})

	t.Run("failure - case 01: command is empty - has nothing", func(t *testing.T) {
		// arrange
		// - parser: default
		ps := gocli.NewParserDefault(optional.None[gocli.ConfigParserDefault]())
		
		// act
		args := ""
		ci, err := ps.ParseCommands(args)

		// assert
		require.Error(t, err)
		require.ErrorIs(t, err, gocli.ErrInvalidCommands)
		require.EqualError(t, err, gocli.ErrInvalidCommands.Error())
		require.Equal(t, gocli.CommandInput{}, ci)
	})

	t.Run("failure - case 02: command is empty - only has app", func(t *testing.T) {
		// arrange
		// - parser: default
		ps := gocli.NewParserDefault(optional.None[gocli.ConfigParserDefault]())

		// act
		args := "app --flag1 value1 -flag2 value2 -O1 -O2"
		ci, err := ps.ParseCommands(args)

		// assert
		require.Error(t, err)
		require.ErrorIs(t, err, gocli.ErrInvalidCommands)
		require.EqualError(t, err, gocli.ErrInvalidCommands.Error())
		require.Equal(t, gocli.CommandInput{}, ci)
	})
}

// TestParserDefault_ParseFlags tests the method ParseFlags of the ParserDefault type.
func TestParserDefault_ParseFlags(t *testing.T) {
	t.Run("success - case 01: flags are parsed correctly - one flag", func(t *testing.T) {
		// arrange
		// - parser: default
		ps := gocli.NewParserDefault(optional.None[gocli.ConfigParserDefault]())

		// act
		args := "app cmd1 --flag1 value1 -O1 -O2"
		f := ps.ParseFlags(args)

		// assert
		require.Equal(t, map[string]any{
			"flag1": "value1",
		}, f)
	})

	t.Run("success - case 02: flags are parsed correctly - multiple flags", func(t *testing.T) {
		// arrange
		// - parser: default
		ps := gocli.NewParserDefault(optional.None[gocli.ConfigParserDefault]())

		// act
		args := "app cmd1 --flag1 value1 -flag2 value2 -O1 -O2"
		f := ps.ParseFlags(args)

		// assert
		require.Equal(t, map[string]any{
			"flag1": "value1",
			"flag2": "value2",
		}, f)
	})

	t.Run("failure - case 01: flags are empty", func(t *testing.T) {
		// arrange
		// - parser: default
		ps := gocli.NewParserDefault(optional.None[gocli.ConfigParserDefault]())
		
		// act
		args := "app cmd1 -O1 -O2"
		f := ps.ParseFlags(args)

		// assert
		require.Equal(t, map[string]any(nil), f)
	})

	t.Run("failure - case 02: flags are empty - has nothing", func(t *testing.T) {
		// arrange
		// - parser: default
		ps := gocli.NewParserDefault(optional.None[gocli.ConfigParserDefault]())

		// act
		args := ""
		f := ps.ParseFlags(args)

		// assert
		require.Equal(t, map[string]any(nil), f)
	})
}

// TestParserDefault_ParseOptions tests the method ParseOptions of the ParserDefault type.
func TestParserDefault_ParseOptions(t *testing.T) {
	t.Run("success - case 01: options are parsed correctly - one option", func(t *testing.T) {
		// arrange
		// - parser: default
		ps := gocli.NewParserDefault(optional.None[gocli.ConfigParserDefault]())

		// act
		args := "app cmd1 --flag1 value1 -O1"
		o := ps.ParseOptions(args)

		// assert
		require.Equal(t, map[string]int{
			"O1": 1,
		}, o)
	})

	t.Run("success - case 02: options are parsed correctly - multiple options", func(t *testing.T) {
		// arrange
		// - parser: default
		ps := gocli.NewParserDefault(optional.None[gocli.ConfigParserDefault]())

		// act
		args := "app cmd1 --flag1 value1 -O1 -O2"
		o := ps.ParseOptions(args)

		// assert
		require.Equal(t, map[string]int{
			"O1": 1,
			"O2": 1,
		}, o)
	})

	t.Run("failure - case 01: options are empty", func(t *testing.T) {
		// arrange
		// - parser: default
		ps := gocli.NewParserDefault(optional.None[gocli.ConfigParserDefault]())
		
		// act
		args := "app cmd1 --flag1 value1"
		o := ps.ParseOptions(args)

		// assert
		require.Equal(t, map[string]int(nil), o)
	})

	t.Run("failure - case 02: options are empty - has nothing", func(t *testing.T) {
		// arrange
		// - parser: default
		ps := gocli.NewParserDefault(optional.None[gocli.ConfigParserDefault]())

		// act
		args := ""
		o := ps.ParseOptions(args)

		// assert
		require.Equal(t, map[string]int(nil), o)
	})
}

// TestParserDefault_Parse tests the method Parse of the ParserDefault type.
func TestParserDefault_Parse(t *testing.T) {
	t.Run("success - case 01: app + 1 command + 1 flag + 1 option", func(t *testing.T) {
		// arrange
		// - parser: default
		// - args: app cmd1 --flag1 value1 -O1
		ps := gocli.NewParserDefault(optional.None[gocli.ConfigParserDefault]())

		// act
		args := "app cmd1 --flag1 value1 -O1"
		i, err := ps.Parse(args)

		// assert
		require.NoError(t, err)
		require.Equal(t, gocli.Input{
			CommandInput: gocli.CommandInput{
				Name: "app",
				Chain: []string{},
				Command: "cmd1",
			},
			Flags: map[string]any{
				"flag1": "value1",
			},
			Options: map[string]int{
				"O1": 1,
			},
		}, i)
	})

	t.Run("success - case 02: app + 1 command + 2 flags + 2 options", func(t *testing.T) {
		// arrange
		// - parser: default
		// - args: app cmd1 --flag1 value1 -flag2 value2 -O1 -O2
		ps := gocli.NewParserDefault(optional.None[gocli.ConfigParserDefault]())

		// act
		args := "app cmd1 --flag1 value1 -flag2 value2 -O1 -O2"
		i, err := ps.Parse(args)

		// assert
		require.NoError(t, err)
		require.Equal(t, gocli.Input{
			CommandInput: gocli.CommandInput{
				Name: "app",
				Chain: []string{},
				Command: "cmd1",
			},
			Flags: map[string]any{
				"flag1": "value1",
				"flag2": "value2",
			},
			Options: map[string]int{
				"O1": 1,
				"O2": 1,
			},
		}, i)
	})

	t.Run("success - case 03: app + 2 chain + 1 command + 2 flags + 2 options", func(t *testing.T) {
		// arrange
		// - parser: default
		// - args: app cmd1 cmd2 cmd3 --flag1 value1 -flag2 value2 -O1 -O2
		ps := gocli.NewParserDefault(optional.None[gocli.ConfigParserDefault]())

		// act
		args := "app cmd1 cmd2 cmd3 --flag1 value1 -flag2 value2 -O1 -O2"
		i, err := ps.Parse(args)

		// assert
		require.NoError(t, err)
		require.Equal(t, gocli.Input{
			CommandInput: gocli.CommandInput{
				Name: "app",
				Chain: []string{"cmd1", "cmd2"},
				Command: "cmd3",
			},
			Flags: map[string]any{
				"flag1": "value1",
				"flag2": "value2",
			},
			Options: map[string]int{
				"O1": 1,
				"O2": 1,
			},
		}, i)
	})

	t.Run("failure - case 01: app + 1 flag + 1 option | invalid args", func(t *testing.T) {
		// arrange
		// - parser: default
		// - args: app --flag1 value1 -O1
		ps := gocli.NewParserDefault(optional.None[gocli.ConfigParserDefault]())

		// act
		args := "app --flag1 value1 -O1"
		i, err := ps.Parse(args)

		// assert
		require.Error(t, err)
		require.ErrorIs(t, err, gocli.ErrInvalidArgs)
		require.EqualError(t, err, gocli.ErrInvalidArgs.Error())
		require.Equal(t, gocli.Input{}, i)
	})

	t.Run("failure - case 02: 1 flag + 1 option | invalid args", func(t *testing.T) {
		// arrange
		// - parser: default
		// - args: --flag1 value1 -O1
		ps := gocli.NewParserDefault(optional.None[gocli.ConfigParserDefault]())

		// act
		args := "--flag1 value1 -O1"
		i, err := ps.Parse(args)

		// assert
		require.Error(t, err)
		require.ErrorIs(t, err, gocli.ErrInvalidArgs)
		require.EqualError(t, err, gocli.ErrInvalidArgs.Error())
		require.Equal(t, gocli.Input{}, i)
	})

	t.Run("failure - case 03: empty | invalid args", func(t *testing.T) {
		// arrange
		// - parser: default
		// - args: ""
		ps := gocli.NewParserDefault(optional.None[gocli.ConfigParserDefault]())

		// act
		args := ""
		i, err := ps.Parse(args)

		// assert
		require.Error(t, err)
		require.ErrorIs(t, err, gocli.ErrInvalidArgs)
		require.EqualError(t, err, gocli.ErrInvalidArgs.Error())
		require.Equal(t, gocli.Input{}, i)
	})
}