package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/openset/leetcode/internal/base"
	"github.com/openset/leetcode/internal/version"
)

func init() {
	base.Commands = []*base.Command{
		version.CmdVersion,
	}
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(fmt.Sprintf("%s: ", base.CmdName))

	flag.Usage = base.Usage
	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
	}
}
