package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println(time.Now().Format(time.RFC3339))

	fmt.Println(time.Now().UTC().Format(time.RFC3339))

	var nilTime time.Time
	var nil2 time.Time
	fmt.Println(nilTime)
	fmt.Println("v1==v2? : ", nilTime == nil2)
	fmt.Println(nilTime.UTC())

	fmt.Printf("1E9=%f, 1E9 == 1e9? : %t\n", 1E9, 1E9 == 1e9)
	now := time.Now()
	next := now.Add(15 * time.Second)
	nanoSecond := int64(next.Sub(now))
	s := float64(nanoSecond) / 1E9
	fmt.Println(int64(now.Sub(next)))
	fmt.Println(s)

	timestamp := now.Unix()
	timestampUTC := now.UTC().Unix()
	fmt.Println(now, ", ", now.UTC())
	fmt.Println(timestamp, timestampUTC)
	fmt.Println(int32(timestamp))

	timestampNano := time.Now().UnixNano()
	timestampNano32 := int32(time.Now().UnixNano())
	fmt.Println(timestampNano)
	fmt.Println(timestampNano32)

	tr := time.NewTimer(-3 * time.Second)
	<-tr.C
	fmt.Println("-3s over")

	dbZone, _ := time.LoadLocation("Asia/Shanghai")
	now = time.Now().In(dbZone)

	// 当月最后一天23:59:00进行创表，所以不要在月末最后一天23:59:00-24:00:00时间内发版
	tmp := now.Add(1 * time.Minute)
	tmp = time.Date(tmp.Year(), tmp.Month(), 1, 0, 0, 0, 0, dbZone).AddDate(0, 1, 0)
	monthEndDay11 := tmp.Add(-1 * time.Minute)
	fmt.Println(monthEndDay11)

	// now := time.Now().In(dbZone)

	// tmp := now.Add(1 * time.Minute)
	// // now-12-31 23:59:00
	// runTime := time.Date(tmp.Year(), 12, 31, 23, 59, 0, 0, dbZone)
	// fmt.Println(runTime)

	for i := 0; i < 3; i++ {
		now := time.Now().Add(20 * time.Second)
		nextMonth1st := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), 0, 0, time.Local).Add(time.Minute)
		monthEndDay11 := nextMonth1st.Add(-20 * time.Second)
		fmt.Println(monthEndDay11.String())
		timer := time.NewTimer(monthEndDay11.Sub(time.Now()))
		<-timer.C
		log.Println("hello")
	}

	stop := time.NewTimer(15 * time.Second)
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("hh")
		case <-stop.C:
			fmt.Println("timeout")
			return
		}
	}

	st := time.Now()
	time.Sleep(2 * time.Second)

	use := time.Now().Sub(st).Seconds()
	fmt.Println("spent time/s: ", use)
}

type Time time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

// UnmarshalJSON 只用指针类型可以调用此方法
func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormart)
}
