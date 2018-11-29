package utils

import (
	"testing"
	"fmt"
)

func TestGetCur5MinKey(t *testing.T)  {
	ts := 1532966369

	cur5mKey := GetCur5MinKey(ts)
	fmt.Printf("5mkey: %s \r\n",cur5mKey)

	hourKey :=GetHourKey(ts)
	fmt.Printf("hour_key: %s \r\n", hourKey)
}