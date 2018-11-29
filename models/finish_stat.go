package models


//每次请求成功，需要传递的结构体
type ResultStat struct {
	Task_id string `json:"task_id"` // 任务id
	Timeline_id string `json:"timeline_id"` // 排期id
	Dtime int `json:"dtime"`  // 时间戳
	Ip string `json:"ip"` // ip
	Day_new bool `json:"day_new"` // IP是否是当天新记录	
}


type FinishStat struct {
	FinishListHour map[string]int
	FinishList5M map[string]int
	FinishPvNum int	  //
	FinishIpNum int   //
}

type MapKey struct {
	TaskId string
	TimelineId string 
	CurDate string //日期
	HourKey string //小时
	CurKey string //五分钟key
}