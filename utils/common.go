package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"time"
)

// Md5 使用md5加密技术对password进行hash加密
func Md5(password string) string {
	h := md5.New()
	_, err := io.WriteString(h, password)
	if err != nil {
		log.Fatal("md5 failed", err)
	}
	md5Password := string(h.Sum([]byte(nil)))
	md5Password = fmt.Sprintf("%x", md5Password)
	return md5Password
}

// TransformDateToUnix 将2023-09-24T04:45:05+08:00 格式的string类型的时间转化为时间戳int64
func TransformDateToUnix(date string) int64 {
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		log.Fatal(err)
	}
	return t.Unix()
}

// TransformUnixToDate 将时间戳（int64）转化为2023-09-24T04:45:05+08:00（string）
func TransformUnixToDate(date int64) string {
	timeTemplate := "2006-01-02 15:04:05"
	return time.Unix(date, 0).Format(timeTemplate)
}

// CommentTime 将2023-09-24T04:45:05+08:00 格式的string类型的时间转化为09-24(mm-dd)
func CommentTime(date string) string {
	dateInt := TransformDateToUnix(date)
	timeTemplate := "01-02"
	return time.Unix(dateInt, 0).Format(timeTemplate)
}
