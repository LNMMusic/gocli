package gocli

// CommandHandler is the type that represents a command function.
type CommandHandler func(i Input) (err error)

// Command is an struct that represents a command.
type Command struct {
	// Name is the name of the command.
	Name string
	// Description is the description of the command.
	Description string
	// Handler is the handler of the command.
	Handler CommandHandler
}

// Commands is the type that represents a list of commands.
type Commands []Command

// FindHandler is the method that finds a handler by name.
func (c Commands) FindHandler(commandName string) (h CommandHandler, err error) {
	// check if exists
	var exists bool
	for _, cmd := range c {
		if cmd.Name == commandName {
			h = cmd.Handler
			exists = true
			break
		}
	}
	if !exists {
		err = ErrCommandHandlerNotFound
		return
	}

	return
}