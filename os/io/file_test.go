package main

import (
	"encoding/csv"
	"log"
	"os"
	"testing"
	"time"
)

func TestWriteFile(t *testing.T) {
	fileName := "./TestWriteFile.csv"
	file, err := os.Create(fileName)
	if err != nil {
		log.Panicf("create file[%s] error: %v", fileName, err)
		return
	}
	defer file.Close()

	str := "time now: " + time.Now().String()
	_, err = file.WriteString(str)
	if err != nil {
		log.Panicf("write file[%s] error: %v", fileName, err)
	}
}

func TestCsvWriteAll(t *testing.T) {
	fileName := "TestCsvWriteAll.csv"
	data := [][]string{
		{"line1", "Tom", "111"},
		{"line2", "Jerry", "123"},
		{"line3", "Jefff", "457"},
	}

	if err := CsvWriteAll(fileName, data); err != nil {
		log.Printf("error: %v", err)
	}
}

// WriteCsvAll 一次性将内存数据写到文件
func CsvWriteAll(fileName string, data [][]string) (err error) {
	file, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = '|'
	err = writer.WriteAll(data)
	return
}
