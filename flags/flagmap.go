package flags

import "github.com/davidmdm/cli/util"

type FlagMap map[string][]interface{}

func (fm FlagMap) add(key string, value interface{}) {
	sanitizedKey := util.StripDashPrefix(key)
	fm[sanitizedKey] = append(fm[sanitizedKey], value)
}
