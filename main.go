package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type kv struct {
	key   string
	value interface{}
}

var flags []kv

var doubleDashEqualRegex = regexp.MustCompile(`^--\w+=\w*$`)
var dashNoEqual = regexp.MustCompile(`^--?\w+$`)

func main() {

	args := os.Args[1:]

	doubleDashEqualFlags := filterStrings(args, func(str string) bool { return doubleDashEqualRegex.MatchString(str) })
	for _, flag := range doubleDashEqualFlags {
		split := strings.Split(flag, "=")
		flags = append(flags, kv{key: split[0][2:], value: split[1]})
	}

	args = filterStrings(args, func(str string) bool { return !doubleDashEqualRegex.MatchString(str) })

	dashNoEqualIndexes := getIndexes(args, func(str string) bool { return dashNoEqual.MatchString(str) })

	fmt.Println(args)
	fmt.Println("idx of dash chars", dashNoEqualIndexes)

	for idx, pos := range dashNoEqualIndexes {

		if pos+1 < len(args)-1 && dashNoEqualIndexes[idx+1] != pos+1 {
			flags = append(flags, kv{key: args[pos], value: args[pos+1]})
		} else {
			//This implies the next argument is also a dash arg, therefore interpret this as a bool flag
			flags = append(flags, kv{key: args[pos], value: true})
		}
	}

	fmt.Println(flags)

}

func filterStrings(sl []string, pred func(string) bool) []string {
	result := []string{}
	for _, e := range sl {
		if pred(e) {
			result = append(result, e)
		}
	}
	return result
}

func getIndexes(sl []string, pred func(string) bool) []int {
	result := []int{}
	for i := range sl {
		if pred(sl[i]) {
			result = append(result, i)
		}
	}
	return result
}
