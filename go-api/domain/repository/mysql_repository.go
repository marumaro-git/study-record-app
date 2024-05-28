package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// DB情報取得
func getDBInfo() string {
	viper.SetDefault("DB_USER", "dev_go_api")
	viper.SetDefault("DB_PASSWORD", "dev_go_api")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "3306")
	viper.SetDefault("DB_NAME", "study_record")

	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbName := viper.GetString("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	return dsn
}

// DB接続
func SetUpDb() *sql.DB {

	// DB情報取得
	dsn := getDBInfo()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	// DB接続確認
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	log.Println("DB接続成功" + dsn)
	return db
}
