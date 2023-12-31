package gocli_test

import (
	"errors"
	"os"
	"testing"

	"github.com/LNMMusic/gocli"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// TestCLI_Run is the test for the method Run.
func TestCLI_Run(t *testing.T) {
	t.Run("success - case 01: command is executed successfully", func(t *testing.T) {
		// arrange
		// - std-in
		os.Args = []string{"app.exe", "cmd1", "--flag1", "value1", "-flag2", "value2", "-O1", "-O2"}
		// - parser: mock
		pr := gocli.NewParserMock()
		pr.On("Parse", "cmd1 --flag1 value1 -flag2 value2 -O1 -O2").Return(gocli.Input{
			CommandInput: gocli.CommandInput{
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
		}, nil)
		// - commander: mock
		cm := gocli.NewCommanderMock()
		cm.On("FindHandler", "cmd1", mock.Anything).Return(
			gocli.CommandHandler(func(i gocli.Input) (err error) {
				return
			}),
			nil,
		)
		// - cli
		cli := gocli.NewCLI(pr, cm)

		// act
		err := cli.Run()

		// assert
		require.NoError(t, err)
		pr.AssertExpectations(t)
		cm.AssertExpectations(t)
	})

	t.Run("failure - case 01: parser fails", func(t *testing.T) {
		// arrange
		// - std-in
		os.Args = []string{"app.exe", "cmd1", "--flag1", "value1", "-flag2", "value2", "-O1", "-O2"}
		// - parser: mock
		pr := gocli.NewParserMock()
		pr.On("Parse", "cmd1 --flag1 value1 -flag2 value2 -O1 -O2").Return(gocli.Input{}, gocli.ErrInvalidArgs)
		// - commander: mock
		// ...
		// - cli
		cli := gocli.NewCLI(pr, nil)

		// act
		err := cli.Run()

		// assert
		require.Error(t, err)
		require.ErrorIs(t, err, gocli.ErrInvalidArgs)
		require.EqualError(t, err, gocli.ErrInvalidArgs.Error())
		pr.AssertExpectations(t)
	})

	t.Run("failure - case 02: command handler not found", func(t *testing.T) {
		// arrange
		// - std-in
		os.Args = []string{"app.exe", "cmd1", "--flag1", "value1", "-flag2", "value2", "-O1", "-O2"}
		// - parser: mock
		pr := gocli.NewParserMock()
		pr.On("Parse", "cmd1 --flag1 value1 -flag2 value2 -O1 -O2").Return(gocli.Input{
			CommandInput: gocli.CommandInput{
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
		}, nil)
		// - commander: mock
		cm := gocli.NewCommanderMock()
		cm.On("FindHandler", "cmd1", mock.Anything).Return(
			gocli.CommandHandler(func(i gocli.Input) (err error) {
				return
			}),
			gocli.ErrCommandHandlerNotFound,
		)
		// - cli
		cli := gocli.NewCLI(pr, cm)

		// act
		err := cli.Run()
		
		// assert
		require.Error(t, err)
		require.ErrorIs(t, err, gocli.ErrCommandHandlerNotFound)
		require.EqualError(t, err, gocli.ErrCommandHandlerNotFound.Error())
		pr.AssertExpectations(t)
		cm.AssertExpectations(t)
	})

	t.Run("failure - case 03: command handler fails", func(t *testing.T) {
		// arrange
		// - std-in
		os.Args = []string{"app.exe", "cmd1", "--flag1", "value1", "-flag2", "value2", "-O1", "-O2"}
		// - parser: mock
		pr := gocli.NewParserMock()
		pr.On("Parse", "cmd1 --flag1 value1 -flag2 value2 -O1 -O2").Return(gocli.Input{
			CommandInput: gocli.CommandInput{
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
		}, nil)
		// - commander: mock
		cm := gocli.NewCommanderMock(); errCmHandler := errors.New("command handler fails")
		cm.On("FindHandler", "cmd1", mock.Anything).Return(
			gocli.CommandHandler(func(i gocli.Input) (err error) {
				return errCmHandler
			}),
			nil,
		)
		// - cli
		cli := gocli.NewCLI(pr, cm)

		// act
		err := cli.Run()

		// assert
		require.Error(t, err)
		require.ErrorIs(t, err, errCmHandler)
		require.EqualError(t, err, errCmHandler.Error())
	})
}