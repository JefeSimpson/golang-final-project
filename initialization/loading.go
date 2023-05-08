package initialization

import "github.com/spf13/viper"

type Config struct {
	Host       string `mapstructure:"host"`
	User       string `mapstructure:"user"`
	Password   string `mapstructure:"password"`
	DBName     string `mapstructure:"dbname"`
	DBPort     string `mapstructure:"dbport"`
	SSLMode    string `mapstructure:"sslmode"`
	ServerPort string `mapstructure:"port"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
