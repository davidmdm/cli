package flags

import (
	"github.com/davidmdm/cli/util"
	"strconv"
	"strings"
)

// Args represents the flags and positional arguments parsed from command line
type Args struct {
	flags       map[string][]interface{}
	positionals []string
}

func (args Args) addFlag(key string, value interface{}) {
	sanitizedKey := util.StripDashPrefix(key)
	args.flags[sanitizedKey] = append(args.flags[sanitizedKey], value)
}

// StringFlag returns a pointer to a string for the value of the flag key
func (args Args) StringFlag(key string) *string {
	v, ok := args.flags[key]
	if !ok || len(v) == 0 {
		return nil
	}
	str, ok := v[0].(string)
	if !ok {
		return nil
	}
	return &str
}

// StringSliceFlag returns a pointer a slice of string for the value of the flag key
func (args Args) StringSliceFlag(key string) *[]string {
	v, ok := args.flags[key]
	if !ok {
		return nil
	}

	sl := []string{}
	for _, val := range v {
		str, ok := val.(string)
		if ok {
			sl = append(sl, str)
		}
	}

	if len(sl) == 0 {
		return nil
	}
	return &sl
}

// BoolFlag returns a pointer to a bool for the value of the flag key
func (args Args) BoolFlag(key string) *bool {
	v, ok := args.flags[key]
	if !ok || len(v) == 0 {
		return nil
	}
	switch typedValue := v[0].(type) {
	case string:
		result := strings.ToLower(typedValue) == "true"
		return &result
	case bool:
		return &typedValue
	default:
		return nil
	}
}

// IntFlag returns a pointer to a bool for the value of the flag key
func (args Args) IntFlag(key string) *int {
	v, ok := args.flags[key]
	if !ok || len(v) == 0 {
		return nil
	}
	switch typedValue := v[0].(type) {
	case string:
		result, err := strconv.Atoi(typedValue)
		if err != nil {
			return nil
		}
		return &result
	case int:
		return &typedValue
	default:
		return nil
	}
}

// Positionals returns the parsed positinal arguments
func (args Args) Positionals() []string {
	return args.positionals
}

func newArgs() Args {
	return Args{
		flags:       make(map[string][]interface{}),
		positionals: []string{},
	}
}
