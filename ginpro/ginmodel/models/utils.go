package models

import (
	"fmt"
	"time"
)

func Print(str1 string, str2 string) {
	fmt.Printf("%s\n", str1+str2)
}

// TimeStamp2Date 时间戳转换日期
func TimeStamp2Date(timeStamp int64) string {
	t := time.Unix(int64(timeStamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// Date2TimeStamp 日期转换时间戳
func Date2TimeStamp(date string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, date, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// GetTimeStamp 获取当前时间戳
func GetTimeStamp() int64 {
	return time.Now().Unix()
}
