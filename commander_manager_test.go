package gocli_test

import (
	"testing"

	"github.com/LNMMusic/gocli"
	"github.com/stretchr/testify/require"
)

// TestCommandManager_FindCommandManager is the test for the method FindCommandManager.
func TestCommandManager_FindCommandManager(t *testing.T) {
	t.Run("success - case 01: command manager is itself", func(t *testing.T) {
		// arrange
		// - command manager
		cmg := gocli.CommanderManager{
			Name:        "root",
			Description: "root command manager",
		}

		// act
		cm, err := cmg.FindCommandManager()

		// assert
		require.NoError(t, err)
		require.Equal(t, &cmg, cm)
	})

	t.Run("success - case 02: command manager is at level 1", func(t *testing.T) {
		// arrange
		// - command manager
		cmg := gocli.CommanderManager{
			Name:        "root",
			Description: "root command manager",
			CommandManagers: []*gocli.CommanderManager{
				{
					Name:        "level1",
					Description: "level 1 command manager",
				},
			},
		}

		// act
		cm, err := cmg.FindCommandManager("level1")

		// assert
		require.NoError(t, err)
		require.Equal(t, cmg.CommandManagers[0], cm)
	})

	t.Run("success - case 03: command manager is at level 3", func(t *testing.T) {
		// arrange
		// - command manager
		cmg := gocli.CommanderManager{
			Name:        "root",
			Description: "root command manager",
			CommandManagers: []*gocli.CommanderManager{
				{
					Name:        "level1",
					Description: "level 1 command manager",
					CommandManagers: []*gocli.CommanderManager{
						{
							Name:        "level2",
							Description: "level 2 command manager",
							CommandManagers: []*gocli.CommanderManager{
								{
									Name:        "level3",
									Description: "level 3 command manager",
								},
							},
						},
					},
				},
			},
		}

		// act
		cm, err := cmg.FindCommandManager("level1", "level2", "level3")

		// assert
		require.NoError(t, err)
		require.Equal(t, cmg.CommandManagers[0].CommandManagers[0].CommandManagers[0], cm)
	})

	t.Run("failure - case 01: command manager is not found", func(t *testing.T) {
		// arrange
		// - command manager
		cmg := gocli.CommanderManager{
			Name:        "root",
			Description: "root command manager",
			CommandManagers: []*gocli.CommanderManager{
				{
					Name:        "level1",
					Description: "level 1 command manager",
				},
			},
		}

		// act
		cm, err := cmg.FindCommandManager("level2")

		// assert
		require.Error(t, err)
		require.ErrorIs(t, err, gocli.ErrCommandManagerNotFound)
		require.EqualError(t, err, gocli.ErrCommandManagerNotFound.Error())
		require.Nil(t, cm)
	})
}

// TestCommandManager_FindHandler is the test for the method FindHandler.
func TestCommandManager_FindHandler(t *testing.T) {
	t.Run("success - case 01: handler found", func(t *testing.T) {
		// arrange
		// - command manager
		cmg := gocli.CommanderManager{
			Name:        "root",
			Description: "root command manager",
			Cmds: gocli.Commands{
				{
					Name:        "command",
					Description: "command description",
					Handler: func(i gocli.Input) (err error) {
						return
					},
				},
			},
		}

		// act
		h, err := cmg.FindHandler("command")

		// assert
		require.NoError(t, err)
		require.NotNil(t, h)
	})

	t.Run("failure - case 01: handler not found", func(t *testing.T) {
		// arrange
		// - command manager
		cmg := gocli.CommanderManager{
			Name:        "root",
			Description: "root command manager",
			Cmds: gocli.Commands{
				{
					Name:        "command",
					Description: "command description",
					Handler: func(i gocli.Input) (err error) {
						return
					},
				},
			},
		}

		// act
		h, err := cmg.FindHandler("command2")

		// assert
		require.Error(t, err)
		require.ErrorIs(t, err, gocli.ErrCommandHandlerNotFound)
		require.EqualError(t, err, gocli.ErrCommandHandlerNotFound.Error())
		require.Nil(t, h)
	})

	t.Run("failure - case 01: command manager not found", func(t *testing.T) {
		// arrange
		// - command manager
		cmg := gocli.CommanderManager{
			Name:        "root",
			Description: "root command manager",
			CommandManagers: []*gocli.CommanderManager{
				{
					Name:        "level1",
					Description: "level 1 command manager",
				},
			},
		}

		// act
		h, err := cmg.FindHandler("command", "level2")

		// assert
		require.Error(t, err)
		require.ErrorIs(t, err, gocli.ErrCommandManagerNotFound)
		require.EqualError(t, err, gocli.ErrCommandManagerNotFound.Error())
		require.Nil(t, h)
	})
}

// TestCommandManager_AddCommand is the test for the method AddCommand.
func TestCommandManager_AddCommand(t *testing.T) {
	t.Run("success - case 01: add command to root command manager", func(t *testing.T) {
		// arrange
		// - command manager
		cmg := gocli.CommanderManager{
			Name:        "root",
			Description: "root command manager",
		}

		// act
		// - add command
		err := cmg.AddCommand(gocli.Command{
			Name:        "command",
			Description: "command description",
			Handler: func(i gocli.Input) (err error) {
				return
			},
		})

		// assert
		require.NoError(t, err)
		require.Len(t, cmg.Cmds, 1)
	})
}

// TestCommandManager_Group is the test for the method Group.
func TestCommandManager_Group(t *testing.T) {
	t.Run("success - case 01: group command manager to root command manager", func(t *testing.T) {
		// arrange
		// - command manager
		cmg := gocli.CommanderManager{
			Name:        "root",
			Description: "root command manager",
		}

		// act
		// - group command manager
		cm := cmg.Group("group", "group description")

		// assert
		require.NotNil(t, cm)
		require.Len(t, cmg.CommandManagers, 1)
	})
}