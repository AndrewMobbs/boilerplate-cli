package cli

import (
	"github.com/AndrewMobbs/boilerplate-golang-cli/app"
	"github.com/spf13/cobra"
)

type ExampleCommand struct {
	App *app.App
}

// an example command that takes a subcommand
func (c *ExampleCommand) Command() *cobra.Command {
	exampleCmd := &cobra.Command{
		Use:   "example",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. .`,
		Args: cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			c.App.Example()
		},
	}

	// Add the flags here
	// Add subcommands here.
	exampleCmd.AddCommand((&SubCommand{App: c.App}).Command())

	return exampleCmd
}
