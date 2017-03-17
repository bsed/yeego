package yeego

import (
	"Mogu/utils"
	"fmt"

	"github.com/spf13/viper"
)

type configType struct {
	*viper.Viper
}

var Config configType

func MustInitConfig(filePath string, fileName string) {
	Config = configType{viper.New()}
	Config.WatchConfig()
	Config.SetConfigName(fileName)
	//filePath支持相对路径和绝对路径 etc:"/a/b" "b" "./b"
	if filePath[:1] != "/" {
		Config.AddConfigPath(utils.GetCurrentPath() + "/" + filePath + "/")
	} else {
		Config.AddConfigPath(filePath + "/")
	}
	fmt.Println(utils.GetCurrentPath() + "/" + filePath + "/" + fileName)
	if err := Config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err).Error())
	}
}

func (config *configType) GetStringWithoutPrefix(key string) string {
	return Config.GetString(string(RunMode) + "." + key)
}

