package utils

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBSource   string `mapstructure:"DB_SOURCE"`
	HostJagoan string `mapstructure:"HOSTJAGOAN"`
	HostLocal  string `mapstructure:"HOSTLOCAL"`
	Database   string `mapstructure:"DATABASE"`
	Dbuser     string `mapstructure:"DBUSER"`
	Dbpassword string `mapstructure:"DBPASSWORD"`

	AppEnv  string `mapstructure:"APP_ENV"`
	AppHost string `mapstructure:"APP_HOST"`

	MysqlDBHost string `mapstructure:"MYSQL_DB_HOST"`
	MysqlDBPort string `mapstructure:"MYSQL_DB_PORT"`
	MysqlDBName string `mapstructure:"MYSQL_DB_NAME"`
	MysqlDBUser string `mapstructure:"MYSQL_DB_USER"`
	MysqlDBPass string `mapstructure:"MYSQL_DB_PASS"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.SetDefault("APP_ENV", "dev")
	viper.SetDefault("APP_HOST", "127.0.0.1:3000")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
