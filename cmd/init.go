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
	"fmt"
        "path/filepath"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
)

var dataPath string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize application data",
	Long:  `This command initialises whatever setup, such as a database, is needed for the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//dbPath, err := initializeStructures(args)
		/* 		if err != nil {
			log.Fatal(err)
		} */
		fmt.Println("Initialized DB at ", dataPath)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&dataPath, "database", "d", filepath.Join(xdg.DataHome,appName,appName+".db"), "Database file")

}
