package Config

import "github.com/spf13/viper"

const (
	DBHost = "db.host"
	DBPass = "db.password"
	DBPort = "db.port"
	DBUser = "db.user"
	DBName = "db.name"
)

func init() {
	//default
	viper.SetDefault(DBHost, "localhost")
	viper.SetDefault(DBPass, "postgres")
	viper.SetDefault(DBPort, 1234)
	viper.SetDefault(DBUser, "postgres")
	viper.SetDefault(DBName, "postgres")
}
