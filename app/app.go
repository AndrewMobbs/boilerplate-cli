package app

import (
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
	a := App{
		AppName:      appName,
		DatabasePath: databasePath,
		ViperCfg:     viperCfg,
		Logger:       logger,
	}
	if a.DatabasePath != "" {
		a.db = NewAppDB(a.DatabasePath, appName)
	}
	return &a
}

// Init() initializes the application state
// Other methods are best structured in their own files
func (a *App) Init() error {
	if a.db == nil {
		a.db = NewAppDB(a.DatabasePath, a.AppName)
	}
	filestat, err := os.Stat(a.DatabasePath)

	exists := true
	if err != nil {
		if os.IsNotExist(err) {
			exists = false
		} else {
			log.Fatal("Error statting Database file: ", err)
		}
	} else {
		if !filestat.Mode().IsRegular() {
			log.Fatal("Database file exists but isn't regular file.")
		}
	}

	if exists {
		err = a.db.Open()
		a.Logger.Info("Using database : ", a.DatabasePath)
	} else {
		err = a.db.Initialize()
		a.Logger.Info("Created database at : ", a.DatabasePath)
	}

	return err
}
