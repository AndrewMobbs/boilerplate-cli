package cli

import (
	"os"

	"github.com/AndrewMobbs/boilerplate-golang-cli/app"
	"github.com/spf13/cobra"
)

type ExampleSubCommand struct {
	App        *app.App
	flagParam  int
	fixedParam string
	enumParam  string
}

type EnumParam struct {
	action string
}

func (a *EnumParam) String() string {
	return a.action
}

func (a *EnumParam) Type() string {
	return "string"
}

func (a *EnumParam) Set(s string) error {
	switch s {
	case app.EnumParamFoo,
		app.EnumParamBar,
		app.EnumParamBaz:
		a.action = s
		return nil
	case "": // Default value
		a.action = app.EnumParamFoo
		return nil
	}
	return os.ErrInvalid
}

// An example subcommand with positional and flag arguments
func (c *ExampleSubCommand) Command() *cobra.Command {
	var enumParam EnumParam
	subCmd := &cobra.Command{
		Use:   "subcommand <string>",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. `,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c.fixedParam = args[0]
			c.enumParam = enumParam.action
			c.App.SubCommand(c.fixedParam, c.flagParam, c.enumParam)
		},
	}

	// Add the flags here
	subCmd.Flags().IntVarP(&c.flagParam, "flag-param", "f", 5, "Example parameter supplied by flag")
	subCmd.Flags().VarP(&enumParam, "enum-param", "e", "Parameter with constrained valid values; foo|bar|baz (default foo)")

	// Add subcommands here.

	return subCmd
}
