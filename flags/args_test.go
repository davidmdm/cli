package flags

import "testing"

func TestStringFlag(t *testing.T) {
	tt := []struct {
		name     string
		flag     string
		args     []string
		expected string
	}{
		{
			name:     "flag is nil if no arguments match flag key",
			flag:     "test",
			args:     []string{"some", "positional", "args", "-flag"},
			expected: "",
		},
		{
			name:     "string is returned when flag is provided as double dash",
			flag:     "test",
			args:     []string{"position", "arg", "--test", "value"},
			expected: "value",
		},
		{
			name:     "string is returned when flag is provided as single dash",
			flag:     "test",
			args:     []string{"position", "arg", "-test", "value"},
			expected: "value",
		},
		{
			name:     "string is returned when flag is provided as double dash equal",
			flag:     "test",
			args:     []string{"position", "arg", "--test=value"},
			expected: "value",
		},
		{
			name:     "string is returned when flag is provided as double dash and nested between positional args",
			flag:     "test",
			args:     []string{"position", "--test", "value", "arg"},
			expected: "value",
		},
		{
			name:     "string is returned when flag is provided as single dash and nested between positional args",
			flag:     "test",
			args:     []string{"position", "-test", "value", "arg"},
			expected: "value",
		},
		{
			name:     "string is returned when flag is provided as double dash equal and nested between positional args",
			flag:     "test",
			args:     []string{"position", "--test=value", "arg"},
			expected: "value",
		},
		{
			name:     "nil is returned if flag is used as a boolean",
			flag:     "boolFlag",
			args:     []string{"positional", "arg", "--boolFlag"},
			expected: "",
		},
		{
			name:     "first string value is returned if flag is passed multiple times",
			flag:     "multiFlag",
			args:     []string{"pos", "-multiFlag", "first", "--multiFlag", "second", "pos2", "--multiFlag=third"},
			expected: "first",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			args := parseStrings(tc.args)
			p := args.StringFlag(tc.flag)
			if p == nil && tc.expected == "" {
				return
			}
			if p == nil && tc.expected != "" {
				t.Fatalf("StringFlag is nil and expected non empty string value %s\n", tc.expected)
			}
			if *p != tc.expected {
				t.Fatalf("expected `%s` for flag `%s` but got `%s`\n", tc.expected, tc.flag, *p)
			}
		})
	}
}
