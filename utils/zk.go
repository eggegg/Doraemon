package utils

import (
	"time"
	"strings"
	"fmt"
)


//统一这个项目使用的redis key的前缀
func GetBgstatRedisKey(key string) string {
	headKey := "bgstat_"
	if strings.Contains(key, headKey) {
		return key
	}

	return strings.Join([]string{headKey,key},"")
}

func GetHourKey(ts int) string {
	timeStamp := time.Unix(int64(ts), 0)
	hr, _, _ := timeStamp.Clock()

	return fmt.Sprintf("h_%d", hr)
}

// 返回5m的切割key
func GetCur5MinKey(ts int) string {
	timeStamp := time.Unix(int64(ts), 0)
	hr, min, _ := timeStamp.Clock()

	i5m := (hr*60+min)/5

	return fmt.Sprintf("5m_%d", i5m)
	
}