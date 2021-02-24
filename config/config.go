package config

type MongoConf struct {
	Host string `mapstructure:"host" json:"host"`
	PORT int `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type MainConfig struct {
	Mongo MongoConf `mapstructure:"mongo" json:"mongo"`
}
