package configs

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var Config *Configuration

func init() {
	//if (len(os.Args) > 1 && os.Args[1] == "-test.v") || (len(os.Args) > 1 && strings.Contains(strings.ToLower(os.Args[1]), "testlog")) {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, "../.."+string(filepath.Separator))))
	err := os.Chdir(apppath)
	if err != nil {
		panic(err)
	}
	//}
	Config = New()
}

func ReloadConfig() (err error) {
	tempConfig, err := Reload()
	if err != nil {
		return
	}
	Config = tempConfig
	return
}

func getEnvironment() string {
	if len(os.Args) > 1 && os.Args[1] != "-test.v" && !strings.Contains(strings.ToLower(os.Args[1]), "testlog") {
		return os.Args[1]
	}
	return "local"
}

func getConfigFilePath(env string) string {
	return fmt.Sprintf("./resource/conf/config.%s.json", env)
}

func validateAppMode(appMode, env string) {
	if appMode != env {
		panic(errors.New(fmt.Sprintf("Please change 'apps.mode' to '%v'", env)))
	}
}

func New() *Configuration {
	env := getEnvironment()
	path := getConfigFilePath(env)

	viper.SetConfigFile(path)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		//if strings.Contains(strings.ToLower(env), "testlog") {
		//	return nil
		//} else {
		panic(err)
		//}

	}

	defaultConfig := Configuration{}
	err := viper.Unmarshal(&defaultConfig)
	if err != nil {
		panic(err)
	}

	validateAppMode(defaultConfig.Apps.Mode, env)

	viper.OnConfigChange(func(e fsnotify.Event) {
		err := ReloadConfig()
		if err != nil {
			fmt.Println("error reload config: ", err.Error())
		} else {
			fmt.Println("Config file changed:", time.Now().Format(time.RFC1123Z))
		}
	})
	viper.WatchConfig()

	return &defaultConfig
}

func Reload() (*Configuration, error) {
	env := Config.Apps.Mode
	defaultConfig := Configuration{}
	path := getConfigFilePath(env)

	viper.SetConfigFile(path)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		return &defaultConfig, err
	}

	err := viper.Unmarshal(&defaultConfig)
	if err != nil {
		return &defaultConfig, err
	}

	if env != defaultConfig.Apps.Mode {
		return &defaultConfig, errors.New("apps.mode is different from the previous configuration!")
	}

	return &defaultConfig, nil
}
