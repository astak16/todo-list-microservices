package global

import "gorm.io/gorm"

type Config struct {
	MySql  MySql `mapstructure:"mysql" `
	Etcd   Etcd  `mapstructure:"etcd"`
	Server ServerInfo
}

type MySql struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Charset  string `mapstructure:"charset"`
}

type Etcd struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerInfo struct {
	Name string `mapstructure:"name" json:"name"`
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

var (
	ServerConfig Config
	MySqlConfig  MySql
	EtcdConfig   Etcd
	Server       ServerInfo
	DB           *gorm.DB
)

const EtcdServerKey = "/services"
