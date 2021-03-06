package utils

import (
	"time"
	"strconv"
)

// Get Current Timestamp 
func GetCurrentTimestampSec() int {
	ts,_ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	return ts
}


//当前日期转换成字符串
func GetFormattedDateTime(timeType string) string {
	var item string
	switch timeType {
	case "day":
		item = "2006-01-02"
		break
	case "hour":
		item = "2006-01-02 15"
		break
	case "min":
		item = "2006-01-02 15:04"
		break		
	}
	t:= time.Now().Format(item)
	return t
}

func GetDateTimeFromTimeStamp(ts int, timeType string) string {
	var item string
	switch timeType {
	case "day":
		item = "2006-01-02"
		break
	case "hour":
		item = "2006-01-02 15"
		break
	case "min":
		item = "2006-01-02 15:04"
		break
	}
	t:= time.Unix(int64(ts), 0).Format(item)
	return t
}