package initialize

import (
	"api-gateway/global"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}

	global.Server = global.ServerConfig.Server
	global.EtcdConfig = global.ServerConfig.Etcd
}
