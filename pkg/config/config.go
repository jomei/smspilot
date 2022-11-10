package config

import "github.com/spf13/viper"

//https://dev.to/techschoolguru/load-config-from-file-environment-variables-in-golang-with-viper-2j2d

type Config struct {
	DNS        string `mapstructure:"DNS"`
	LogLevel   int    `mapstructure:"LOGLEVEL"`
	LOGFILE    string `mapstructure:"LOGFILE"`
	SIGNINGKEY string `mapstructure:"SIGNINGKEY"`
}

var Conf, _ = LoadConfig(".")

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
