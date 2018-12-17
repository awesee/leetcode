package base

import (
	"fmt"
	"os"
	"strings"
)

const CmdName = "leetcode"

var Commands []*Command

type Command struct {
	Run       func(cmd *Command, args []string)
	UsageLine string
	Short     string
	Long      string
}

func (c *Command) Name() string {
	name := c.UsageLine
	if i := strings.Index(name, " "); i > 0 {
		name = name[0:i]
	}
	return name
}

func (c *Command) Usage() {
	fmt.Printf("usage: %s %s\n", CmdName, c.UsageLine)
	fmt.Printf("Run '%s help %s' for details.\n", CmdName, c.Name())
	Exit()
}

func Usage() {
	fmt.Printf("%s is a tool for managing leetcode source code.\n\n", CmdName)
	fmt.Println("Usage:")
	fmt.Printf("\t%s <command> [arguments]\n", CmdName)
	fmt.Println("The commands are:")
	for _, cmd := range Commands {
		fmt.Printf("\t%-11s \t%s\n", cmd.Name(), cmd.Short)
	}
	fmt.Println()
	fmt.Printf(`Use "%s help <command>" for more information about a command.`, CmdName)
	fmt.Println()
	Exit()
}

func Exit() {
	os.Exit(0)
}
