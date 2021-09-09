package cli

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/AndrewMobbs/boilerplate-golang-cli/app"
	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// initConfig reads in config file and ENV variables if set.
// This code largely taken from https://github.com/carolynvs/stingoftheviper
/*
MIT License

Copyright (c) 2020 Carolyn Van Slyck

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
func InitConfig(cmd *cobra.Command, a *app.App, defaultConfigName string, cfgFile string) error {
	createConfigFile := false
	configFilePath := ""
	defaultConfigPath := filepath.Join(xdg.ConfigHome, a.AppName, defaultConfigName)
	// User has specified a non-default location
	if cfgFile != defaultConfigPath {
		configFilePath = cfgFile
		filestat, err := os.Stat(cfgFile)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				createConfigFile = true
			} else {
				a.Logger.Fatal("Error statting config file : ", err)
			}
		} else {
			if !filestat.Mode().IsRegular() {
				a.Logger.Fatalf("Config file %s must be a regular file", filestat.Name())
			}
		}
	} else { // no user-specified config file, look for an existing one
		var err error
		configFilePath, err = xdg.SearchConfigFile(filepath.Join(a.AppName, defaultConfigName))
		// If that fails, try creating one in the default XDG location
		if err != nil {
			configFilePath = defaultConfigPath
			createConfigFile = true
		} else {
			a.Logger.Info("Found config file at : ", configFilePath)
		}
	}

	if createConfigFile {
		if err := os.MkdirAll(filepath.Dir(configFilePath), os.ModeDir|0700); err != nil {
			a.Logger.Fatal("Error Creating config file path : ", err)
		}
		_, err := os.Create(configFilePath)
		if err != nil {
			a.Logger.Fatal(err)
		}
		a.Logger.Info("Created empty config file at : ", configFilePath)
	}

	a.ViperCfg.SetConfigFile(configFilePath)
	// If a config file is found, read it in.
	if err := a.ViperCfg.ReadInConfig(); err == nil {
		a.Logger.Info("Using config file : ", a.ViperCfg.ConfigFileUsed())
	} else {
		return err
	}
	a.ViperCfg.SetEnvPrefix(a.AppName) // read in environment variables that match
	a.ViperCfg.AutomaticEnv()
	err := bindFlags(cmd, a.ViperCfg, a.AppName)

	return err
}

func bindFlags(cmd *cobra.Command, v *viper.Viper, appName string) error {
	var e error
	e = nil
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			v.BindEnv(f.Name, fmt.Sprintf("%s_%s", appName, envVarSuffix))
		}
		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			err := cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
			if err != nil {
				e = err
			}
		} else {
			err := v.BindPFlag(f.Name, f)
			if err != nil {
				e = err
			}
		}
	})
	return e
}
