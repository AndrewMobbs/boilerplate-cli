package cli

import (
	"path/filepath"

	"github.com/AndrewMobbs/boilerplate-golang-cli/app"
	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
)

type InitCommand struct {
	App *app.App
}

// A command to initialize all configuration, database and eny other similar things for the application.
// If not relevant doesn't need to exist
func (c *InitCommand) Command() *cobra.Command {
	dataPath := ""
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize application data",
		Long:  `This command initialises whatever setup, such as a database, is needed for the application.`,
		Run: func(cmd *cobra.Command, args []string) {
			c.App.DatabasePath = dataPath
			err := c.App.Init()
			if err != nil {
				c.App.Logger.Fatal("Error initializing app database: ", err)
			}
			c.App.ViperCfg.Set("database", dataPath)
			err = c.App.ViperCfg.WriteConfig()
			if err != nil {
				c.App.Logger.Fatal("Error writing config: ", err)
			}
		},
	}
	initCmd.Flags().StringVarP(&dataPath, "database", "d", filepath.Join(xdg.DataHome, c.App.AppName, c.App.AppName+".db"), "Database file")

	return initCmd
}
