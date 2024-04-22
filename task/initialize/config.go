package initialize

import (
	"task/global"

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
	global.MySqlConfig = global.ServerConfig.MySql
	global.EtcdConfig = global.ServerConfig.Etcd
	global.Server = global.ServerConfig.Server
}
