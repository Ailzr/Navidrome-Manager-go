package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func init() {
	//使用viper读取配置文件信息
	workDir, _ := os.Getwd()
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("conf load success ...")
}
