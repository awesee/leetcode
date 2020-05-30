package base

import (
	"fmt"
	"strings"
)

// Commands - base.Commands
var Commands []*Command

// Command - base.Command
type Command struct {
	Run       func(cmd *Command, args []string)
	UsageLine string
	Short     string
	Long      string
	Hidden    bool
}

// Name - base.Command.Name
func (c *Command) Name() string {
	name := c.UsageLine
	if i := strings.Index(name, " "); i > 0 {
		name = name[0:i]
	}
	return name
}

// Usage - base.Command.Usage
func (c *Command) Usage() {
	fmt.Printf("usage: %s %s\n\n", CmdName, c.UsageLine)
	fmt.Printf("Run '%s help %s' for details.\n", CmdName, c.Name())
}

// UsageHelp - base.Command.UsageHelp
func (c *Command) UsageHelp() {
	fmt.Printf("usage: %s %s\n\n", CmdName, c.UsageLine)
	fmt.Println(c.Long)
}

// Register - base.Register
func Register(cmds ...*Command) {
	Commands = append(Commands, cmds...)
}
