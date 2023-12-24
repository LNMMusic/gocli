package gocli

import (
	"regexp"
	"strings"

	"github.com/LNMMusic/optional"
)

// ConfigParserDefault is the struct that wraps the configuration of the default parser.
type ConfigParserDefault struct {
	// PatternCLI is the regexp pattern of the full command line.
	PatternCLI string
	// PatternChain is the regexp pattern of the chain.
	PatternChain string
	// PatternFlags is the regexp pattern of the flag.
	PatternFlag string
	// PatternOptions is the regexp pattern of the option.
	PatternOption string
	// Trimmer is a white space trimmer in between.
	Trimmer string
}

// NewParserDefault is the function that creates a new default parser.
func NewParserDefault(cfg optional.Option[ConfigParserDefault]) (p *ParserDefault) {
	// default configuration
	defaultCfg := ConfigParserDefault{
		PatternCLI: `^(\w+(?:\s+\w+)+)(\s+-{1,2}\w+\s+\w+)*(\s+-[A-Z0-9]+)*$`,
		PatternChain: `^(\w+(?:\s+\w+)+)`,
		PatternFlag: `(\s+-{1,2}\w+\s+\w+)+`,
		PatternOption: `(\s+-[A-Z0-9]+)+`,
		Trimmer: `\s{2,}`,
	}
	if cfg.IsSome() {
		config := cfg.Unwrap()
		if config.PatternCLI != "" {
			defaultCfg.PatternCLI = config.PatternCLI
		}
		if config.PatternChain != "" {
			defaultCfg.PatternChain = config.PatternChain
		}
		if config.PatternFlag != "" {
			defaultCfg.PatternFlag = config.PatternFlag
		}
		if config.PatternOption != "" {
			defaultCfg.PatternOption = config.PatternOption
		}
		if config.Trimmer != "" {
			defaultCfg.Trimmer = config.Trimmer
		}
	}

	p = &ParserDefault{
		patternCLI: regexp.MustCompile(defaultCfg.PatternCLI),
		patternCommand: regexp.MustCompile(defaultCfg.PatternChain),
		patternFlag: regexp.MustCompile(defaultCfg.PatternFlag),
		patternOption: regexp.MustCompile(defaultCfg.PatternOption),
		Trimmer: regexp.MustCompile(defaultCfg.Trimmer),
	}
	return
}

// ParserDefault is the struct that wraps the default parser.
type ParserDefault struct {
	// patternCLI is the regexp pattern of the full command line.
	// - default: `^(\w+(?:\s+\w+)+)(\s+-{1,2}\w+\s+\w+)*(\s+-[A-Z0-9]+)*$`
	patternCLI *regexp.Regexp
	// patternCommand is the regexp pattern of the command.
	// - default: `^(\w+(?:\s+\w+)+)`
	patternCommand *regexp.Regexp
	// patternFlags is the regexp pattern of the flag.
	// - default: `(\s+-{1,2}\w+\s+\w+)+`
	patternFlag *regexp.Regexp
	// patternOptions is the regexp pattern of the option.
	// - default: `(\s+-[A-Z0-9]+)+`
	patternOption *regexp.Regexp
	// Trimmer is a white space trimmer in between.
	// - default: `\s{2,}`
	Trimmer *regexp.Regexp
}

// Parse is the method that parses the input.
func (p *ParserDefault) Parse(args string) (i Input, err error) {
	// check matching between args and patternCLI
	if !p.patternCLI.MatchString(args) {
		err = ErrInvalidArgs
		return
	}

	// commands
	commands, _ := p.ParseCommands(args)
	// flags
	flags := p.ParseFlags(args)
	// options
	options := p.ParseOptions(args)

	// input
	i = Input{
		CommandInput: commands,
		Flags: flags,
		Options: options,
	}
	return
}

// ParseCommands is the method that parses the commands.
func (p *ParserDefault) ParseCommands(args string) (c CommandInput, err error) {
	// check matching between args and patternCommand
	// - at least 2 commands: app and command
	if !p.patternCommand.MatchString(args) {
		err = ErrInvalidCommands
		return
	}
	match := p.patternCommand.FindStringSubmatch(args)[0]

	// parsing
	// trim white spaces
	// - start and end
	match = strings.TrimSpace(match)
	// - in between
	match = p.Trimmer.ReplaceAllString(match, " ")

	// commands (split match)
	commands := strings.Split(match, " ")
	// - check if valid
	size := len(commands)
	// command input
	c = CommandInput{
		Name: commands[0],
		Chain: commands[1:size-1],  // if there are is no chain, it will be empty
		Command: commands[size-1],
	}
	return
}

// ParseFlags is the method that parses the flags.
func (p *ParserDefault) ParseFlags(args string) (f map[string]any) {
	// check matching between args and patternFlag
	if !p.patternFlag.MatchString(args) {
		return
	}
	match := p.patternFlag.FindStringSubmatch(args)[0]

	// parsing
	// trim white spaces
	// - start and end
	match = strings.TrimSpace(match)
	// - in between
	match = p.Trimmer.ReplaceAllString(match, " ")

	// flags (split match)
	flags := strings.Split(match, " ")
	size := len(flags)
	f = make(map[string]any)
	// parse flags
	for i := 0; i < size; i+=2 {
		// split
		key := strings.TrimLeft(flags[i], "-")
		value := flags[i+1]

		// add to map
		f[key] = value
	}

	return
}

// ParseOptions is the method that parses the options.
func (p *ParserDefault) ParseOptions(args string) (o map[string]int) {
	// check matching between args and patternOption
	if !p.patternOption.MatchString(args) {
		return
	}
	match := p.patternOption.FindStringSubmatch(args)[0]

	// parsing
	// trim white spaces
	// - start and end
	match = strings.TrimSpace(match)
	// - in between
	match = p.Trimmer.ReplaceAllString(match, " ")

	// options (split match)
	options := strings.Split(match, " ")
	size := len(options)
	o = make(map[string]int)
	for i := 0; i < size; i++ {
		// parse key
		key := strings.TrimLeft(options[i], "-")
		// add to map
		o[key] = 1
	}

	return
}