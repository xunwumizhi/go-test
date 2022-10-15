package main

import (
	"database/sql"
	"log"
)

// AddUser 添加用户
func AddUser(loginName string, pwd string) error {

	stmtIn, err := mysqlDB.Prepare("INSERT INTO users(login_name, pwd) values(?, ?)") //本包中的变量dbConn
	defer stmtIn.Close()
	if err != nil {
		return err
	}
	_, err = stmtIn.Exec(loginName, pwd) //注意err类型已由前面确定了，这里不需要使用`:=`，否则出现编译错误`no new variables on left side of :=`
	if err != nil {
		return err
	}

	return nil
}

// GetUserCredential 获取信息
func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := mysqlDB.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	defer stmtOut.Close()
	if err != nil {
		log.Printf("SELECT error: %s", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	return pwd, nil
}
