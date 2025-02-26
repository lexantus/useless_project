package configs

import "github.com/spf13/viper"

type DBConfig struct {
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func GetConfig() *DBConfig {
	var configInstance *DBConfig

	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&configInstance); err != nil {
		panic(err)
	}
	return configInstance
}
