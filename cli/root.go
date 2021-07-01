package cli

import (
	"fmt"
	"path/filepath"

	"github.com/AndrewMobbs/boilerplate-golang-cli/app"
	"github.com/adrg/xdg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type RootCommand struct {
	App               *app.App
	DefaultConfigName string
	DefaultLogLevel   string
}

func (c *RootCommand) Command() *cobra.Command {
	c.App.Logger = log.New()
	c.App.ViperCfg = viper.New()
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
			c.App.DatabasePath = c.App.ViperCfg.GetString("database")
			// loglevel might have changed from environment/config file. Try again.
			if c.App.ViperCfg.GetString("loglevel") != "" && !logFlagSet {
				setLogLevel(c.App.Logger, c.App.ViperCfg.GetString("loglevel"))
			}
			return err
		},
		//	Run: func(cmd *cobra.Command, args []string) {
		// Working with OutOrStdout/OutOrStderr allows us to unit test our command easier
		//	cmd.OutOrStdout()
		//	},
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
