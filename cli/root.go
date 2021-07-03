package cli

import (
	"fmt"
	"path/filepath"

	"github.com/AndrewMobbs/boilerplate-golang-cli/app"
	"github.com/adrg/xdg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	App               *app.App
	DefaultConfigName string
	DefaultLogLevel   string
}

func (c *RootCommand) Command() *cobra.Command {
	cfgFile := ""
	loglevel := ""

	rootCmd := &cobra.Command{
		Use:   c.App.AppName,
		Short: c.App.AppName,
		Long:  c.App.AppName + ` is an example of a golang CLI driven application, with config and state database support.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// set initial log level from command params.
			logFlagSet := cmd.Flags().Changed("loglevel")
			setLogLevel(c.App.Logger, loglevel)
			err := InitConfig(cmd, c.App, c.DefaultConfigName, cfgFile)
			// Open database connection if db has been configured
			if err == nil {
				c.App.DatabasePath = c.App.ViperCfg.GetString("database")
				if c.App.DatabasePath != "" {
					err = c.App.OpenAppDB()
				}
			}
			// loglevel might have changed from environment/config file. Try again.
			if c.App.ViperCfg.GetString("loglevel") != "" && !logFlagSet {
				setLogLevel(c.App.Logger, c.App.ViperCfg.GetString("loglevel"))
			}
			return err
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			err := c.App.Close()
			if err != nil {
				c.App.Logger.Warn("error in tidy-up : ", err)
			}
			return err
		},
	}
	// Add app-wide persistent flags
	doPersistentFlags(rootCmd, c.App, c.DefaultConfigName, c.DefaultLogLevel, &cfgFile, &loglevel)

	// Create app structure

	// Add top level commands
	rootCmd.AddCommand((&ExampleCommand{App: c.App}).Command())
	rootCmd.AddCommand((&InitCommand{App: c.App}).Command())

	return rootCmd
}

func setLogLevel(logger *log.Logger, level string) {
	l, err := log.ParseLevel(level)
	if err != nil {
		logger.Fatal("Invalid log level : ", level)
	}
	logger.SetLevel(l)
}

// doPersistentFlags adds persistent flags (i.e. valid for all commands) to the root command
// This could be tider, especially if more persistent flags are needed with defaults!
func doPersistentFlags(cmd *cobra.Command, a *app.App, defaultConfigName string, defaultLogLevel string, cfgFile *string, loglevel *string) {
	cmd.PersistentFlags().StringVar(cfgFile, "config", filepath.Join(xdg.ConfigHome, a.AppName, defaultConfigName), "config file")
	levels := ""
	for _, l := range log.AllLevels {
		levels = levels + ", " + l.String()
	}
	levels = levels[2:]
	defaultLevel := a.ViperCfg.GetString("loglevel")
	if defaultLevel == "" {
		defaultLevel = defaultLogLevel
	}

	cmd.PersistentFlags().StringVar(loglevel, "loglevel", defaultLevel, fmt.Sprintf("Log Level. Valid levels: [%s]", levels))
}
