// Package base provides base support.
package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// URL - base.URL
const URL = "https://github.com/awesee/leetcode/tree/main"

// CmdName - base.CmdName
var CmdName = filepath.Base(os.Args[0])

// Usage - base.Usage
func Usage() {
	fmt.Printf("%s is a tool for managing leetcode source code.\n\n", CmdName)
	fmt.Println("Usage:")
	fmt.Printf("\t%s <command> [arguments]\n", CmdName)
	fmt.Println("The commands are:")
	for _, cmd := range Commands {
		if !cmd.Hidden {
			fmt.Printf("\t%-11s \t%s\n", cmd.Name(), cmd.Short)
		}
	}
	fmt.Printf("\nUse \"%s help <command>\" for more information about a command.\n", CmdName)
}

// FilePutContents - base.FilePutContents
func FilePutContents(filename string, data []byte) []byte {
	ext := filepath.Ext(filename)
	if strings.EqualFold(ext, ".json") {
		data = JSONIndent(data)
	}
	if len(data) > 0 {
		filename = getFilePath(filename)
		err := ioutil.WriteFile(filename, data, 0644)
		CheckErr(err)
	}
	return data
}

// JSONIndent - base.JSONIndent
func JSONIndent(src []byte) []byte {
	if !json.Valid(src) {
		return nil
	}
	var buf bytes.Buffer
	err := json.Indent(&buf, src, "", "\t")
	CheckErr(err)
	return buf.Bytes()
}

// CheckErr - base.CheckErr
func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// AuthInfo - base.AuthInfo
func AuthInfo(cmd string) string {
	format := "<!--|This file generated by command(leetcode %s); DO NOT EDIT.%s|-->\n"
	format += "<!--+----------------------------------------------------------------------+-->\n"
	format += "<!--|@author    awesee <openset.wang@gmail.com>                           |-->\n"
	format += "<!--|@link      https://github.com/awesee                                 |-->\n"
	format += "<!--|@home      https://github.com/awesee/leetcode                        |-->\n"
	format += "<!--+----------------------------------------------------------------------+-->\n"
	return fmt.Sprintf(format, cmd, strings.Repeat(" ", 15-len(cmd)))
}

func getFilePath(filename string) string {
	if dir := filepath.Dir(filename); dir != "" {
		err := os.MkdirAll(dir, 0755)
		CheckErr(err)
	}
	return filename
}
