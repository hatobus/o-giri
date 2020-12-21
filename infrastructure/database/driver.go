package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/hatobus/o-giri/config"
)

func Connect(conf config.MySQLConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&loc=Asia%%2FTokyo",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DataBase,
	)
	return sql.Open("mysql", dsn)
}
