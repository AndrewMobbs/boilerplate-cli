package app

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// App gives the global persistent state for the application
// Example fields
type App struct {
	db           *appDB
	AppName      string
	DatabasePath string
	ViperCfg     *viper.Viper
	Logger       *log.Logger
}

// NewApp() creates a new App object
func NewApp(appName string, databasePath string, viperCfg *viper.Viper, logger *log.Logger) *App {
	logger.Trace("App.NewApp(%s,%s,%T,%T)", appName, databasePath, viperCfg, logger)
	a := App{
		AppName:      appName,
		DatabasePath: databasePath,
		ViperCfg:     viperCfg,
		Logger:       logger,
	}
	if a.DatabasePath != "" {
		err := a.OpenAppDB()
		if err != nil {
			a.Logger.Fatal("error opening database : ", err)
		}
	}
	return &a
}

// Init() initializes the application state
// Other methods are best structured in their own files
func (a *App) Init() error {
	a.Logger.Trace("App.Init()")
	if a.db == nil {
		a.db = NewAppDB(a.DatabasePath, a.AppName)
	}
	exists := a.checkDBExists()
	var err error
	if exists {
		err = a.db.Open()
		a.Logger.Info("Using database at : ", a.DatabasePath)
	} else {
		err = a.db.Initialize()
		a.Logger.Info("Created database at : ", a.DatabasePath)
	}
	return err
}

// OpenAppDB opens the application database if it exists
func (a *App) OpenAppDB() error {
	a.Logger.Trace("App.OpenAppDB()")
	if a.DatabasePath == "" {
		return fmt.Errorf("database path not set, run init")
	}
	if a.db != nil {
		a.db.Close()
	}
	exists := a.checkDBExists()
	if exists {
		a.db = NewAppDB(a.DatabasePath, a.AppName)
		return a.db.Open()
	}
	return fmt.Errorf("no database at %s", a.DatabasePath)
}

// Close tidies up any data structures, such as open database connection
func (a *App) Close() error {
	a.Logger.Trace("App.Close()")
	if a.db != nil {
		return a.db.Close()
	}
	return nil
}

func (a *App) checkDBExists() bool {
	a.Logger.Trace("App.checkDBExists()")
	filestat, err := os.Stat(a.DatabasePath)
	exists := true
	if err != nil {
		if os.IsNotExist(err) {
			exists = false
		} else {
			a.Logger.Fatal("Error statting Database file: ", err)
		}
	} else {
		if !filestat.Mode().IsRegular() {
			a.Logger.Fatal("Database file exists but isn't regular file.")
		}
	}
	return exists
}
