// The main file of the project

package main

import (
	"os"
	"pseudo-lang/plc"
)

func main() {
	exitCode := plc.Compiler(os.Args)
	os.Exit(exitCode)
}
