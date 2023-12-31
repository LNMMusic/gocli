package gocli

import "github.com/stretchr/testify/mock"

// NewCommanderMock is the function that returns a new CommanderMock.
func NewCommanderMock() (r *CommanderMock) {
	r = &CommanderMock{}
	return
}

// CommanderMock is the mock that implements the Commander interface.
type CommanderMock struct {
	mock.Mock
}

// FindHandler is the method that finds a handler by name.
func (m *CommanderMock) FindHandler(commandName string, commandChain ...string) (h CommandHandler, err error) {
	// args
	args := m.Called(commandName, commandChain)

	// return
	h = args.Get(0).(CommandHandler)
	err = args.Error(1)
	return
}

// AddCommand is the method that adds a command to the registry.
func (m *CommanderMock) AddCommand(command Command) (err error) {
	// args
	args := m.Called(command)

	// return
	err = args.Error(0)
	return
}

// Group is the method that groups a list of commands.
func (m *CommanderMock) Group(name string, description string) (cm Commander) {
	// args
	args := m.Called(name, description)

	// return
	cm = args.Get(0).(Commander)
	return
}
