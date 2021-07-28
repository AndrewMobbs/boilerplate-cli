package app

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/AndrewMobbs/appdb"
)

type appDB struct {
	db      *sql.DB
	Path    string
	AppName string
}

const schemaVersion uint8 = 1

// schema returns the database schema. Golang doesn't provide const arrays/slices.
func schema() []string {
	return []string{
		`CREATE TABLE id(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		`,
	}
}

func NewAppDB(path string, appName string) *appDB {
	return &appDB{
		Path:    path,
		AppName: appName,
	}
}

// Open() opens the app database.
func (s *appDB) Open() error {
	var err error
	if s.Path == "" {
		return fmt.Errorf("database not configured, run init command")
	}
	if s.db == nil {
		s.db, err = appdb.Open(s.Path, s.AppName, schemaVersion)
	}
	return err
}

// Initialize the database, deploying the schema
func (s *appDB) Initialize() error {
	db, err := appdb.InitAppDB(s.Path, s.AppName, schemaVersion, schema())
	if err == nil {
		db.Close()
	}
	return err
}

func (s *appDB) Close() error {
	if s.db != nil {
		return s.db.Close()
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
