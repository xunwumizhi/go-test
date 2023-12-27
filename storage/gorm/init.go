package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	username = "root"
	passwd   = "root"
	// database = "streaming_media_service"
	database = "pod_group_statistic"

	protocol = "tcp"
	host     = "localhost"
	port     = 3306
)

var dsn string = fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=True", username, passwd, protocol, host, port, database)
var db *gorm.DB

func init() {
	var err error
	conf := mysql.Config{DSN: dsn}
	gormConf := &gorm.Config{}
	db, err = gorm.Open(mysql.New(conf), gormConf)
	if err != nil {
		log.Fatal("open db fail")
	}
}
