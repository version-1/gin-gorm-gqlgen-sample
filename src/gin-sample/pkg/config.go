package configs

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Database Database
}

type Database struct {
	User     string `toml:"user"`
	Password string `toml:"password"`
	Address  string `toml:"address"`
	Name     string `toml:"name"`
}

func GetConfig(env string) Config {
	var config Config
	_, err := toml.DecodeFile("configs/"+env+".toml", &config)
	if err != nil {
		panic(err)
	}

	return config
}
