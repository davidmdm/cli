package flags

import (
	"fmt"
	"github.com/davidmdm/cli/util"
	"regexp"
	"strings"
)

var doubleDashEqualRegex = regexp.MustCompile(`^--\w+=\w*$`)
var dashNoEqual = regexp.MustCompile(`^--?[^=]+$`)

// Parse transforms a slice of strings into a flagMap
func Parse(args []string) FlagMap {

	var flags = FlagMap{}

	fmt.Println("Original args:", args)

	doubleDashEqualFlags := util.FilterStrings(args, func(str string, _ int) bool { return doubleDashEqualRegex.MatchString(str) })
	for _, flag := range doubleDashEqualFlags {
		split := strings.Split(flag, "=")
		flags.add(split[0], split[1])
	}

	args = util.FilterStrings(args, func(str string, _ int) bool { return !doubleDashEqualRegex.MatchString(str) })

	dashNoEqualIndexes := util.GetIndexes(args, func(str string) bool { return dashNoEqual.MatchString(str) })

	fmt.Println("args:", args)
	fmt.Println("idx of dash chars:", dashNoEqualIndexes)

	nonPositionalIndexMap := make(map[int]bool)
	dashCount := len(dashNoEqualIndexes)

	for idx, pos := range dashNoEqualIndexes {
		nonPositionalIndexMap[pos] = true
		if idx < dashCount-1 && dashNoEqualIndexes[idx+1] == pos+1 {
			//the dash elems are adjacent interpret position as bool
			flags.add(args[pos], true)
		} else if pos == len(args)-1 {
			// the dash element is last and therefore a bool
			flags.add(args[pos], true)
		} else {
			flags.add(args[pos], args[pos+1])
			nonPositionalIndexMap[pos+1] = true
		}
	}

	args = util.FilterStrings(args, func(_ string, i int) bool {
		return !nonPositionalIndexMap[i]
	})

	fmt.Printf("Flags: %+v\n", flags)

	fmt.Println("positional args:", args)

	return flags

}
