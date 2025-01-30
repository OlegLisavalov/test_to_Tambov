package config

import(
	"github.com/spf13/viper"
)

type Config struct{
	DBHost string
	DBPort string
	DBUser string
	DBPassword string
	DBName string
}

func LoadConfig() *Config {
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5433")
	viper.SetDefault("DB_USER", "oleglisavalov")
	viper.SetDefault("DB_PASSWORD", "123123")
	viper.SetDefault("DB_NAME", "db_for_task")

	viper.AutomaticEnv()

	return &Config{
		DBHost: viper.GetString("DB_HOST"),
		DBPort: viper.GetString("DB_PORT"),
		DBUser: viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName: viper.GetString("DB_NAME"),
	}
}