package models


//每次请求成功，需要传递的结构体
type FinishStat struct {
	Task_id string `json:"task_id"` // 任务id
	Timeline_id string `json:"timeline_id"` // 排期id
	Dtime int `json:"dtime"`  // 时间戳
	Ip string `json:"ip"` // ip
	Day_new bool `json:"day_new"` // IP是否是当天新记录	
}