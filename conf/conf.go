package conf

import (
	"io/ioutil"
	"github.com/BurntSushi/toml"
	"errors"
)

var (
	Conf              config
	defaultConfigFile = "conf/conf.toml"
)

type config struct {
	// 数据库配置
	DB database `toml:"database"`
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