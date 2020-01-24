package main

import (
	"fmt"
	"os"

	"github.com/davidmdm/cli/flags"
)

func main() {

	args := flags.ParseStrings(os.Args[1:])

	namespace := args.StringFlag("namespace")
	if namespace == nil {
		fmt.Println("namespace is nil")
	} else {
		fmt.Println("value of namespace:", *namespace)
	}

	fmt.Printf("args: %+v\n", args)

}
