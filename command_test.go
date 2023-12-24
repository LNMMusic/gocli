package gocli

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestCommands_FindHandler tests the method FindHandler of the Commands type.
func TestCommands_FindHandler(t *testing.T) {
	t.Run("success - case 01: command exists", func(t *testing.T) {
		// arrange
		cmds := Commands([]Command{
			{Name: "cmd1", Description: "cmd1 description", Handler: func(i Input) (err error) { return }},
		})

		// act
		h, err := cmds.FindHandler("cmd1")

		// assert
		require.NoError(t, err)
		require.NotNil(t, h)
	})

	t.Run("error - case 01: command does not exist", func(t *testing.T) {
		// arrange
		cmds := Commands([]Command{
			{Name: "cmd1", Description: "cmd1 description", Handler: func(i Input) (err error) { return }},
		})

		// act
		h, err := cmds.FindHandler("cmd2")

		// assert
		require.Error(t, err)
		require.ErrorIs(t, err, ErrCommandHandlerNotFound)
		require.EqualError(t, err, ErrCommandHandlerNotFound.Error())
		require.Nil(t, h)
	})
}