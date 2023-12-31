package gocli

// NewCommanderManager is the function that creates a new command manager.
func NewCommanderManager(name, description string) (cm *CommanderManager) {
	cm = &CommanderManager{
		Name: name,
		Description: description,
	}
	return
}

// CommanderManager is the struct that represents a command manager.
type CommanderManager struct {
	// Name is the name of the command composite.
	Name string
	// Description is the description of the command composite.
	Description string
	// Commands are the commands of the command composite.
	Cmds Commands
	// CommandManagers is the command composite of the command composite.
	CommandManagers []*CommanderManager
	
	// TODO: implement indexed commands
	// ...
}

// FindHandler is the method that finds a handler by name.
func (c *CommanderManager) FindHandler(commandName string, commandChain ...string) (h CommandHandler, err error) {
	// find command manager
	cmg, err := c.FindCommandManager(commandChain...)
	if err != nil {
		return
	}

	// fetch command
	h, err = cmg.Cmds.FindHandler(commandName)
	if err != nil {
		return
	}

	return
}

// FindCommandManager is the method that finds a command manager by name.
func (c *CommanderManager) FindCommandManager(commandChain ...string) (cm *CommanderManager, err error) {
	size := len(commandChain)
	// check if len of commandChain is 0
	if size == 0 {
		cm = c
		return
	}

	// fetch last command manager
	var cmg []*CommanderManager = c.CommandManagers
	for i:=0; i<size; i++ {
		// exists flag to check if the command manager exists
		var exists bool

		// iterate over command managers to check if the name exists
		for j := range cmg {
			if cmg[j].Name == commandChain[i] {
				// set that command manager is on right track
				exists = true

				// last case
				if i == size-1 {
					cm = cmg[j]
					break
				}

				// set next inner command manager
				cmg = cmg[j].CommandManagers
				continue
			}
		}

		// check if exists
		if !exists {
			cmg = nil
			err = ErrCommandManagerNotFound
			return
		}
	}

	return
}

// AddCommand is the method that adds a command to the command manager.
func (c *CommanderManager) AddCommand(cmd Command) (err error) {
	(*c).Cmds = append((*c).Cmds, cmd)
	return
}

// Group is the method that groups a command manager.
func (c *CommanderManager) Group(name, description string) (cm Commander) {
	cmg := &CommanderManager{
		Name: name,
		Description: description,
	}
	(*c).CommandManagers = append((*c).CommandManagers, cmg)

	cm = cmg
	return

}