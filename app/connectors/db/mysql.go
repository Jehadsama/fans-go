package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "fans-go/app/config"
	"fans-go/app/libs/utils"
)

func getMySQLUri() string {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	database := os.Getenv("MYSQL_DATABASE")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		user,
		password,
		host,
		port,
		database)
}

func ConnectToMysql() *gorm.DB {
	db, err := gorm.Open(mysql.Open(getMySQLUri()), &gorm.Config{})
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
