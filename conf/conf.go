package conf

import (
	"errors"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

var (
	Conf              config
	defaultConfigFile = "./conf/conf.toml"
)

type config struct {
	App app      `toml:"app"`
	DB  database `toml:"database"`
}

type app struct {
	GrpcPort string `toml:"grpcport"`
	HttpPort string `toml:"httpport"`
}

type database struct {
	Host     string
	Database string
	Port     string
	User     string
	Password string
}

func InitConfig() error {
	configBytes, err := ioutil.ReadFile(defaultConfigFile)
	if err != nil {
		return errors.New("config load err:" + err.Error())
	}
	_, err = toml.Decode(string(configBytes), &Conf)
	if err != nil {
		return errors.New("config decode err:" + err.Error())
	}
	return nil
}
