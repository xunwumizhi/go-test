package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("open db fail")
	}
}
