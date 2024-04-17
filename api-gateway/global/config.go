package global

type Config struct {
	Server ServerInfo `mapstructure:"server"`
	Etcd   Etcd       `mapstructure:"etcd"`
	// UserService UserServiceConfig `mapstructure:"user"`
}

type ServerInfo struct {
	Name string `mapstructure:"name" json:"name"`
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type Etcd struct {
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	Prefix string `mapstructure:"prefix"`
}

// type UserServiceConfig struct {
// 	Host string `mapstructure:"host"`
// 	Port int    `mapstructure:"port"`
// }

var (
	ServerConfig Config
	// UserService  UserServiceConfig
	Server     ServerInfo
	EtcdConfig Etcd
)

const EtcdServerKey = "/services"
