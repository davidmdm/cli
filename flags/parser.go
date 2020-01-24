package flags

import (
	"os"
	"regexp"
	"strings"

	"github.com/davidmdm/cli/util"
)

var flagRegex = regexp.MustCompile(`^--?[^-]`)

// Parse parses os.Args
func Parse() Args {
	return parseStrings(os.Args[1:])
}

// parseStrings transforms a slice of strings into a flagMap
func parseStrings(args []string) Args {

	parsedArgs := newArgs()

	flagIndexes := util.GetIndexes(args, func(str string) bool { return flagRegex.MatchString(str) })
	flagCount := len(flagIndexes)
	nonPositionalIndexMap := make(map[int]bool)

	for idx, pos := range flagIndexes {
		nonPositionalIndexMap[pos] = true
		flag := args[pos]

		doubleDash := flag[:2] == "--"
		equalIdx := strings.IndexByte(flag, '=')
		if doubleDash && equalIdx > -1 {
			parsedArgs.addFlag(flag[:equalIdx], flag[equalIdx+1:])
		} else if idx < flagCount-1 && flagIndexes[idx+1] == pos+1 {
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
