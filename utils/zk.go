package utils

import (
	"strings"
)


//统一这个项目使用的redis key的前缀
func GetBgstatRedisKey(key string) string {
	headKey := "bgstat_"
	if strings.Contains(key, headKey) {
		return key
	}

	return strings.Join([]string{headKey,key},"")
}
