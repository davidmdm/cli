package cli

import "github.com/davidmdm/cli/flags"

import "fmt"

import "os"

type App struct {
	Name     string
	Short    string
	Long     string
	Root     *Command
	Commands []Command
}

func (app App) Exec() error {
	if app.Root == nil && len(app.Commands) == 0 {
		return fmt.Errorf("App not implemented")
	}

	args := flags.Parse()
	app.Root.setArgs(args)
	for i := range app.Commands {
		app.Commands[i].setArgs(args)
	}

	positionals := args.Positionals()
	if len(positionals) > 0 && len(app.Commands) > 0 {
		firstPosition := positionals[0]
		for _, cmd := range app.Commands {
			if cmd.Name == firstPosition {
				return cmd.exec()
			}
		}
	}

	if app.Root == nil && len(positionals) > 0 {
		return fmt.Errorf("no command matched %s", positionals[0])
	}

	return app.Root.exec()
}

type Command struct {
	Name          string
	Aliases       []string
	Description   string
	Run           func(args flags.Args) error
	RequiredFlags []string
	SubCommands   []Command

	args flags.Args
}

func (c *Command) setArgs(args flags.Args) {
	c.args = args
}

func (c Command) exec() error {
	// TODO parse flags on App exec and set a flagset on each command so we don't parse every time
	args := flags.Parse()

	// first thing we want to do is delegate to a more specific command (say for more specific help information)
	positionals := args.Positionals()
	if len(positionals) > 0 && len(c.SubCommands) > 0 {
		for _, command := range c.SubCommands {
			if command.Name == positionals[0] {
				return command.exec()
			}
		}
	}

	// If help flag is active print help info
	if args.HasFlag("help") || args.HasFlag("h") {
		fmt.Fprintf(os.Stderr, "%s\n  %s\n\n", c.Name, c.Description)
		return nil
	}

	// If no Runner is implemented (and we know we are at the deepest level of subcommand already) then nothing is implemented
	if c.Run == nil {
		return fmt.Errorf("command defined but not implemented")
	}

	return c.Run(args)
}
