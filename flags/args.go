package flags

import "github.com/davidmdm/cli/util"

// Args represents the flags and positional arguments parsed from command line
type Args struct {
	flags       map[string][]interface{}
	positionals []string
}

func (args Args) addFlag(key string, value interface{}) {
	sanitizedKey := util.StripDashPrefix(key)
	args.flags[sanitizedKey] = append(args.flags[sanitizedKey], value)
}

func newArgs() Args {
	return Args{
		flags:       make(map[string][]interface{}),
		positionals: []string{},
	}
}
