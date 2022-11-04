package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// db settings
const (
	USERNAME = "root"
	PASSWORD = "root"
	NETWORK  = "tcp"
	SERVER   = "localhost"
	PORT     = 3306
	DATABASE = "streaming_media_service"
)

// mysqlDB  generated by sql.Open()
var mysqlDB *sql.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	// dsn = "root:root@/streaming_media_service?charset=utf8"
	var err error
	mysqlDB, err = sql.Open("mysql", dsn)

	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return
	}

	mysqlDB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	mysqlDB.SetMaxOpenConns(100)                  //设置最大连接数
	mysqlDB.SetMaxIdleConns(16)                   //设置闲置连接数
}