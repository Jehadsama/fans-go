package db

import (
	"os"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "fans-go/app/config"
	"fans-go/app/libs/utils"
)

func ConnectToMysql() *gorm.DB {

	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	net := os.Getenv("mysql_net")
	host := os.Getenv("MYSQL_HOST")
	database := os.Getenv("MYSQL_DATABASE")
	parsetime, _ := strconv.ParseBool(os.Getenv("MYSQL_PARSETIME"))

	mysqlConfig := &mysql.Config{
		User:      user,
		Passwd:    password,
		Net:       net,
		Addr:      host,
		DBName:    database,
		ParseTime: parsetime,
	}

	db, err := gorm.Open(gormMysql.New(gormMysql.Config{
		DriverName: "mysql",
		DSN:        mysqlConfig.FormatDSN(),
	}), &gorm.Config{})

	utils.CheckError("ConnectToMysql,db", err)

	sqlDB, err := db.DB()
	utils.CheckError("ConnectToMysql,sqlDB", err)

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

var DB = ConnectToMysql()
