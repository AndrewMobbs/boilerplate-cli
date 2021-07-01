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
and usage of using your command.`,
		// By specifying zero positional args, a command with subcommands will throw an error for typos in subcommand names.
		Args: cobra.ExactArgs(0),
		// If the command should just print help listing subcommands, remove the Run clause.
		Run: func(cmd *cobra.Command, args []string) {
			c.App.Example()
		},
	}

	// Add any flags here
	// Add subcommands here. Propagate the App struct to maintain config state.
	exampleCmd.AddCommand((&ExampleSubCommand{App: c.App}).Command())

	return exampleCmd
}
