package dcliconfig

import (
	"daikincli/internal/fileutil"
	"fmt"
	"io/ioutil"

	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

//Config is a central config object for the app
type Config struct {
}

var once sync.Once

var (
	instance *Config
)

//GetString returns the key value as a string
func (s *Config) GetString(key ConfigName) string {
	return viper.GetString(string(key))
}

//GetStringDefault returns the key value as a string, or the default value if not found
func (s *Config) GetStringDefault(key ConfigName, defaultValue string) string {
	val := viper.GetString(string(key))
	if val != "" {
		return val
	}
	return defaultValue
}

//GetBool returns the key value as a bool
func (s *Config) GetBool(key ConfigName) bool {
	return viper.GetBool(string(key))
}

//GetConfig returns the current dcli config object. This object is a shared singleton. This will panic if there are problems
func GetConfig() (config *Config) {

	ex, err := os.Executable()

	if err != nil {
		panic(fmt.Errorf("Config failure: %w", err))
	}

	exPath := filepath.Dir(ex)
	exPath, err = fileutil.FileWalk(exPath, "dcli.conf")

	if err != nil {
		//Test will run themselves in the /tmp somewhere, so we need to handle this scenario
		if strings.Contains(filepath.Dir(ex), "/tmp") {
			err = ioutil.WriteFile("./dcli.conf", []byte(testConfig), 0644)
			exPath = "./"
			if err != nil {
				panic(fmt.Errorf("Fatal saving sample config file %w ", err))
			}
			fmt.Println("Warning using the test version of the config file")
		} else {
			panic("Could not find dcli.conf")
		}
	}

	once.Do(func() { // <-- atomic, does not allow repeating

		viper.AutomaticEnv()
		viper.AddConfigPath(exPath)
		viper.SetConfigName("dcli.conf")
		viper.SetConfigType("env")

		err := viper.ReadInConfig() // Find and read the config file
		if err != nil {             // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file %w ", err))
		}

		instance = &Config{}
	})

	if instance == nil {
		panic(fmt.Errorf("Config failure"))
	}

	return instance
}
