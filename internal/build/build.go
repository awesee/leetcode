// Package build implements the command build.
package build

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/openset/leetcode/internal/base"
	"github.com/openset/leetcode/internal/version"
)

// CmdBuild - build.CmdBuild
var CmdBuild = &base.Command{
	Run:       runBuild,
	UsageLine: "build",
	Short:     "cross compiler",
	Long:      "compile packages for darwin, windows, linux.",
	Hidden:    true,
}

func runBuild(cmd *base.Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
		return
	}
	target := map[string][]string{
		"darwin":  {"amd64"},
		"windows": {"386", "amd64"},
		"linux":   {"386", "amd64"},
	}
	err := os.Setenv("CGO_ENABLED", "0")
	base.CheckErr(err)
	for goos, goarchList := range target {
		err = os.Setenv("GOOS", goos)
		base.CheckErr(err)
		for _, goarch := range goarchList {
			err = os.Setenv("GOARCH", goarch)
			base.CheckErr(err)
			fileName := fmt.Sprintf("%s%s.%s-%s.zip",
				base.CmdName,
				version.String(),
				goos,
				goarch,
			)
			fmt.Println("OS:", goos, "\tArch:", goarch, "\t", fileName)
			file, err := os.Create(fileName)
			base.CheckErr(err)
			zw := zip.NewWriter(file)
			zf, err := zw.CreateHeader(&zip.FileHeader{
				Name:     binName(),
				Method:   zip.Deflate,
				Modified: time.Now(),
			})
			base.CheckErr(err)
			err = exec.Command("go", "build", "-ldflags", "-s -w", "github.com/openset/leetcode").Run()
			base.CheckErr(err)
			src, err := os.Open(binName())
			base.CheckErr(err)
			_, err = io.Copy(zf, src)
			base.CheckErr(err)
			err = src.Close()
			base.CheckErr(err)
			err = zw.Close()
			base.CheckErr(err)
			err = os.Remove(binName())
			base.CheckErr(err)
		}
	}
}

func binName() string {
	var exeSuffix string
	if os.Getenv("GOOS") == "windows" {
		exeSuffix = ".exe"
	}
	return base.CmdName + exeSuffix
}
