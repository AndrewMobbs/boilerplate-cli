package app

// The logic for the "example" command, separated from CLI plumbing
func (a *App) Example() error {
	a.Logger.Info("Ran example command")
	return nil
}
func (a *App) SubCommand(fixedParam string, flagParam int) error {
	a.Logger.Warnf("Fixed param %s, flag param %d\n", fixedParam, flagParam)
	return nil
}
