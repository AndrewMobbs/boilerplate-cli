/*
Copyright © 2021 Andrew Mobbs <andrew.mobbs@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package main

import (
	"github.com/AndrewMobbs/boilerplate-golang-cli/app"
	"github.com/AndrewMobbs/boilerplate-golang-cli/cli"
	"github.com/sirupsen/logrus"
)

const (
	appName           = "exampleApp"
	defaultConfigName = "config.yaml"
	defaultLogLevel   = "warning"
)

// This is a boilerplate example for a golang CLI application
// It uses Cobra and Viper, but structured for no global variables
// It automagically binds environment variables, config variables and command-line options
// into a Viper configuration set.
// Format for Environment variables is APPNAME_PARAMETER (this does not work on persistent flags)
func main() {
	// execute root command (which cascades to subcommands)
	a := app.NewApp(appName, "", nil, nil)
	rootCmd := cli.RootCommand{App: a, DefaultConfigName: defaultConfigName, DefaultLogLevel: defaultLogLevel}
	err := rootCmd.Command().Execute()
	if err != nil {
		logrus.Fatal(err)
	}
}
