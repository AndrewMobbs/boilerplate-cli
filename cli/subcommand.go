package cli

import (
	"github.com/AndrewMobbs/boilerplate-golang-cli/app"
	"github.com/spf13/cobra"
)

type ExampleSubCommand struct {
	App        *app.App
	flagParam  int
	fixedParam string
}

// An example subcommand with positional and flag arguments
func (c *ExampleSubCommand) Command() *cobra.Command {
	subCmd := &cobra.Command{
		Use:   "subcommand <string>",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. `,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c.fixedParam = args[0]
			c.App.SubCommand(c.fixedParam, c.flagParam)
		},
	}

	// Add the flags here
	subCmd.Flags().IntVarP(&c.flagParam, "flagparam", "f", 5, "Example parameter supplied by flag")

	// Add subcommands here.

	return subCmd
}
