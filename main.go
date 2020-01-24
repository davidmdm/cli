package main

import (
	"fmt"
	"os"

	"github.com/davidmdm/cli/flags"
)

func main() {

	args := os.Args[1:]

	fm := flags.Parse(args)

	fmt.Println("orginal:", args)
	fmt.Printf("args: %+v\n", fm)
}
