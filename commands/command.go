package commands

import "fmt"

type Command struct {
	// The command name, usually the executable's name.
	Name string

	// The one-line usage message.
	ShortDesc string

	// The long message shown in the 'help <cmd>' output.
	LongDesc string

	// Examples of how to use the command
	Example string

	// List of all valid non-flag arguments.
	ValidArgs []string

	// List of commands supported.
	commands []*Command

	// Actual args parsed from flags.
	args []string

	// Parent command for this command
	parent *Command
}

func (c *Command) SetArgs(args []string) {
	c.args = args
}

func (c *Command) Execute() error {
	fmt.Println("Command.Execute()")
	return nil
}

// Adds one command to this parent command.
func (c *Command) AddCommand(cmd *Command) {
	if cmd == c {
		panic("Command can't be a child of itself")
	}
	if cmd.parent != nil {
		if cmd.parent != c {
			panic("Command belongs to anothor command")
		} else {
			return
		}
	}
	cmd.parent = c
	c.commands = append(c.commands, cmd)
}