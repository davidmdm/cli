package flags

import (
	"regexp"
	"os"
	"strings"

	"github.com/davidmdm/cli/util"
)

var doubleDashEqualRegex = regexp.MustCompile(`^--\w+=\w*$`)
var dashNoEqual = regexp.MustCompile(`^--?[^=]+$`)

// Parse parses os.Args
func Parse() Args {
	return parseStrings(os.Args[1:])
}

// parseStrings transforms a slice of strings into a flagMap
func parseStrings(args []string) Args {

	parsedArgs := newArgs()

	doubleDashEqualFlags := util.FilterStrings(args, func(str string, _ int) bool { return doubleDashEqualRegex.MatchString(str) })
	for _, flag := range doubleDashEqualFlags {
		split := strings.Split(flag, "=")
		parsedArgs.addFlag(split[0], split[1])
	}

	args = util.FilterStrings(args, func(str string, _ int) bool { return !doubleDashEqualRegex.MatchString(str) })

	dashNoEqualIndexes := util.GetIndexes(args, func(str string) bool { return dashNoEqual.MatchString(str) })
	dashCount := len(dashNoEqualIndexes)
	nonPositionalIndexMap := make(map[int]bool)

	for idx, pos := range dashNoEqualIndexes {
		nonPositionalIndexMap[pos] = true
		if idx < dashCount-1 && dashNoEqualIndexes[idx+1] == pos+1 {
			//the dash elems are adjacent interpret position as bool
			parsedArgs.addFlag(args[pos], true)
		} else if pos == len(args)-1 {
			// the dash element is last and therefore a bool
			parsedArgs.addFlag(args[pos], true)
		} else {
			parsedArgs.addFlag(args[pos], args[pos+1])
			nonPositionalIndexMap[pos+1] = true
		}
	}

	parsedArgs.positionals = util.FilterStrings(args, func(_ string, i int) bool { return !nonPositionalIndexMap[i] })

	return parsedArgs
}
