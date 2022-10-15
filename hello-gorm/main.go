package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	// Local --- CST
	// PRC ---CST
	// GMT ---GMT
	dsn := "root:root@tcp(localhost:3306)/pod_group_statistic?charset=utf8&parseTime=True&loc=Local"
	DB, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("open err: ", err)
		return
	}
	rows, err := DB.Raw("SELECT created_at FROM group_statistics_2019_11 WHERE group_id=68").Rows()
	if err != nil {
		log.Fatalln("rows err: ", err)
	}

	for rows.Next() {
		row := TestTimezone{}
		if err := DB.ScanRows(rows, &row); err != nil {
			log.Fatalln("scan err, ", err)
		}

		fmt.Println(row.CreatedAt)
	}

	fmt.Println(time.Now())
	//===================test=============//
	fmt.Printf("fmt print: %v\n", time.Local)
	log.Printf("log print: %v\n", time.Local)
	fmt.Println(time.Now().Day())

	// 当前月1号
	t1 := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Local)
	// 上个月最后一天
	t := t1.AddDate(0, 0, -1)
	fmt.Println(t.String())

	t = t.AddDate(0, 0, -5)
	fmt.Println(t.String())

	dt := time.Date(2008, 06, 10, 07, 10, 11, 100, time.Local)
	// t := time.Date(2018, 01, 10, 07, 10, 11, 100, time.Local)
	fmt.Println(dt)
	fmt.Println(dt.Format("2006_01_02 15:04:05"))
	fmt.Println(dt.Year())
	fmt.Println(fmt.Sprintf("%02d", int(dt.Month())))
	// dt2 := time.Date(2018, 1, 9, 23, 59, 22, 100, time.Local)

	tableName := "`" + "group_statistics-2019-10" + "`"

	sql := `INSERT INTO ` + tableName + `(created_date, created_at) VALUES(?, ?) `
	db.Exec(sql, dt, dt)

	// pod = Pod{Name: "gorm2", KubeIndex: "idc-sz/000000037529000229000006/cluster-front-test-cmf-idc-sz-28436-j5wtt", CPULimit: 2000, CPURequest: 1000}
	// db.Table("pod").Create(&pod)

	// pod = Pod{Name: "gorm3", KubeIndex: "idc-sh/000000037529000229000006/cluster-front-test-cmf-idc-sh-28436-j5wtt", CPULimit: 2000, CPURequest: 500}
	// db.Table("pod").Create(&pod)

	// pod = Pod{Name: "gorm4", KubeIndex: "idc-wx/000000037529000229000006/cluster-front-test-cmf-idc-wx-28436-j5wtt", CPULimit: 1000, CPURequest: 500}
	// db.Table("pod").Create(&pod)

	type Res struct {
		CPURequestSum int64
		CPULimitSum   int64
	}
	res := Res{}
	db.Raw("SELECT SUM(cpu_limit) AS cpu_limit_sum, SUM(cpu_request) AS cpu_request_sum FROM pod").Scan(&res)
	fmt.Println(res)

	res1 := struct {
		CreatedAt time.Time
	}{}
	db.Raw("SELECT created_at FROM pod WHERE name = ?", "gorm2").Scan(&res1)
	if createBytes, err := json.Marshal(res1.CreatedAt); err != nil {
		fmt.Println("marshal err")
	} else {
		fmt.Println("marshal vaule: " + string(createBytes))
	}
	fmt.Println(res1.CreatedAt)

	pod := &Pod{}
	db.Table("pod").Where("name = ?", "gorm2").First(pod)
	fmt.Println(pod.CPULimit)
} // main end
