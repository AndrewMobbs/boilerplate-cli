/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"database/sql"
	"log"
	"path/filepath"

	"github.com/AndrewMobbs/appdb"
	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
)

var dataPath string
var db *sql.DB

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize application data",
	Long:  `This command initialises whatever setup, such as a database, is needed for the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		db, err = appdb.InitAppDB(dataPath, appName, 1, schema())
		if err != nil {
			log.Fatal("Error initialising database: ", err)
		}
		viperCfg.Set("database", dataPath)
		err = viperCfg.WriteConfig()
		if err != nil {
			log.Fatal("Error writing config: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&dataPath, "database", "d", filepath.Join(xdg.DataHome, appName, appName+".db"), "Database file")

}

func schema() []string {
	return []string{
		`CREATE TABLE id(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		`,
	}
}
