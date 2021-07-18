package app

import (
	"fmt"
)

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
func (a *App) SubCommand(fixedParam string, enumParam string) error {
	a.Logger.Tracef("SubCommand(%s,%s)", fixedParam, enumParam)
	fmt.Printf("flag-param is %d\nenum-param is %s\n", a.ViperCfg.GetInt("flag-param"), enumParam)
	a.Logger.Warn("Ran SubCommand")
	return nil
}
