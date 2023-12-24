package gocli

// NewRegistry is the function that returns a new Registry.
func NewRegistry(p Parser, c Commander) (r Registry) {
	r = Registry{
		Parser: p,
		Commander: c,
	}
	return
}

// Registry is the struct that wraps the commands.
type Registry struct {
	// Parser is the parser of the registry.
	// it returns the Input of the command handler.
	Parser
	
	// Commander is the commander of the registry.
	// helps finding the command handler from the input.
	Commander
}