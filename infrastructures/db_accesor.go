package infrastructures

import (
	_ "github.com/go-sql-driver/mysql" // for using mysql
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"log"
)

const (
	// ConfigDir directory path for searching file name
	ConfigDir = "."
	// ConfigName configuration file name
	ConfigName = "db_credential"
)

var (
	// DB pointer
	DB *sqlx.DB
)

// loadConfig loads configuration file
func loadConfig() error {
	viper.SetConfigName(ConfigName)
	viper.AddConfigPath(ConfigDir)
	return viper.ReadInConfig()
}

// OpenDB Connecting to DB
func OpenDB() error {
	if err := loadConfig(); err != nil {
		return err
	}

	user := viper.GetString("user")
	password := viper.GetString("password")
	host := viper.GetString("host")
	dbName := viper.GetString("name")
	log.Println(user + ":" + password + "@tcp(" + host + ")/" + dbName + "?charset=utf8mb4&parseTime=true&locl=Asia%2FTokyo")
	DB = sqlx.MustConnect(
		"mysql",
		user+":"+password+"@tcp("+host+")/"+dbName+"?charset=utf8mb4",
	)
	res := DB.MustExec("CREATE DATABASE IF NOT EXISTS hiname;")
	_, err := res.RowsAffected()
	return err
}

// CloseDB Closing DB connection
func CloseDB() {
	DB.Close()
}
