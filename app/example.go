package app

// The logic for the "example" command, separated from CLI plumbing
func (a *App) Example() error {
	a.Logger.Trace("Example()")

	a.Logger.Info("Ran example command")
	return nil
}
func (a *App) SubCommand(fixedParam string, flagParam int) error {
	a.Logger.Tracef("SubCommand(%s,%d)", fixedParam, flagParam)
	a.Logger.Warn("Ran SubCommand")
	return nil
}
