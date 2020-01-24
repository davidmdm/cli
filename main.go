package main

import (
	"fmt"

	"github.com/davidmdm/cli/flags"
)

func main() {

	args := flags.Parse()

	namespace := args.StringFlag("namespace")
	if namespace == nil {
		fmt.Println("namespace is nil")
	} else {
		fmt.Println("value of namespace:", *namespace)
	}

	fmt.Printf("args: %+v\n", args)

}
