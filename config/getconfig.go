package config

import "github.com/spf13/viper"

func GetConfig() *Config {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
