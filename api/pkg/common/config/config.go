package config

import "github.com/spf13/viper"

type Config struct {
	Port string `mapstructur:"PORT"`
	UserDB string `mapstructur:"USER_DB"`
	PassDB string `mapstructur:"PASSWORD_DB"`
	DbName string `mapstructur:"NAME_DB"`
	HostDb string `mapstructur:"HOST_DB"`
}

func LoadConfig() (c Config, error) {
	viper.AddConfigPath("./pkg/common/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)
	return
}