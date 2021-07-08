package app

const (
	EnumParamFoo = "foo"
	EnumParamBar = "bar"
	EnumParamBaz = "baz"
)

// The logic for the "example" command, separated from CLI plumbing
func (a *App) Example() error {
	a.Logger.Trace("Example()")

	a.Logger.Info("Ran example command")
	return nil
}
func (a *App) SubCommand(fixedParam string, flagParam int, enumParam string) error {
	a.Logger.Tracef("SubCommand(%s,%d)", fixedParam, flagParam, enumParam)
	a.Logger.Warn("Ran SubCommand")
	return nil
}
