package main

import (
	"os"

	"github.com/davidmdm/cli/flags"
)

func main() {

	args := os.Args[1:]

	flags.Parse(args)

}
