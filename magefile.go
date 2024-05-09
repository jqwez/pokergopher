//go:build mage
// +build mage

package main

import (
	"os"

	"github.com/magefile/mage/sh"
)

func Build() error {
	if err := Test(); err != nil {
		return err
	}
	return sh.Run("go", "build", "-o", "bin/pokergopher", "main.go")
}

func BuildNoTest() error {
	return sh.Run("go", "build", "-o", "bin/pokergopher", "main.go")
}

func Test() error {
	os.Setenv("MAGEFILE_VERBOSE", "true")
	return sh.Run("go", "test", "./...")
}

func Run() error {
	return sh.Run("./bin/pokergopher")
}

func GoRun() error {
	return sh.Run("go", "run", "main.go")
}
