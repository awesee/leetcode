package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("leetcode: ")

	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: ")
	flag.PrintDefaults()
}
