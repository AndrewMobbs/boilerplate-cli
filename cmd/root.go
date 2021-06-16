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
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/spf13/viper"
)

var cfgFile string
var viperCfg *viper.Viper

const (
	appName           = "myAppName"
	defaultConfigName = "config.yaml"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   appName,
	Short: "blah blah blah",
	Long:  `blah blah blah .`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return initConfig(cmd)
	},
	//	Run: func(cmd *cobra.Command, args []string) {
	// Working with OutOrStdout/OutOrStderr allows us to unit test our command easier
	//	cmd.OutOrStdout()
	//	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", filepath.Join(xdg.ConfigHome, appName, defaultConfigName), "config file")

}

// initConfig reads in config file and ENV variables if set.
func initConfig(cmd *cobra.Command) error {
	v := viper.New()
	viperCfg = v
	createConfigFile := false
	configFilePath := ""
	defaultConfigPath := filepath.Join(xdg.ConfigHome, appName, defaultConfigName)
	// User has specified a non-default location
	if cfgFile != defaultConfigPath {
		configFilePath = cfgFile
		filestat, err := os.Stat(cfgFile)
		if err != nil {
			if os.IsNotExist(err) {
				createConfigFile = true
			} else {
				log.Fatal("Error statting config file: ", err)
			}
		} else {
			if !filestat.Mode().IsRegular() {
				log.Fatalf("Config file %s must be a regular file", filestat.Name())
			}
		}
	} else { // no user-specified config file, look for an existing one
		var err error
		configFilePath, err = xdg.SearchConfigFile(filepath.Join(appName, defaultConfigName))
		// If that fails, try creating one in the default XDG location
		if err != nil {
			configFilePath = defaultConfigPath
			createConfigFile = true
		} else {
			log.Println("Found config file at ", configFilePath)
		}
	}

	if createConfigFile {

		if err := os.MkdirAll(filepath.Dir(configFilePath), os.ModeDir|0700); err != nil {
			log.Fatal("Error Creating config file path: ", err)
		}
		_, err := os.Create(configFilePath)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Created empty config file at", configFilePath)
	}

	log.Println("Config file at ", configFilePath)
	v.SetConfigFile(configFilePath)
	// If a config file is found, read it in.
	if err := v.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", v.ConfigFileUsed())
	}

	v.SetEnvPrefix(appName) // read in environment variables that match
	v.AutomaticEnv()
	bindFlags(cmd, v)

	return nil
}

func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			v.BindEnv(f.Name, fmt.Sprintf("%s_%s", appName, envVarSuffix))
		}
		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}
