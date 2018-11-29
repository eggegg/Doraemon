package models

import (

	mgo "gopkg.in/mgo.v2"	
	"gopkg.in/mgo.v2/bson"	

	"github.com/pkg/errors"
	"github.com/eggegg/Doraemon/utils"
	"github.com/eggegg/Doraemon/helper/env"
	
	"fmt"
	
)

type TaskFinishStat struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"` 	
	TaskId string `json:"task_id" bson:"task_id"`
	Date string `json:"date" bson:"date"`
	AddDtime string `json:"add_dtime" bson:"add_dtime"`
	FinishListHour map[string]int `json:"finish_list_hour" bson:"finish_list_hour"`
	FinishListIpHour map[string]int `json:"finish_list_ip_hour" bson:"finish_list_ip_hour"`
	FinishList5M map[string]int `json:"finish_list_5m" bson:"finish_list_5m"`
	FinishIpNum int `json:"finish_ip_num" bson:"finish_ip_num"`
	FinishPvNum int `json:"finish_pv_num" bson:"finish_pv_num"`
	LastClickDtime string `json:"last_click_dtime" bson:"last_click_dtime"`
	FinishListAllIpHour map[string]int `json:"finish_list_all_ip_hour" bson:"finish_list_all_ip_hour"`
	UpdateTime int `json:"update_time" bson:"update_time"`
}


func updateTaskFinishStat(mapKey MapKey, finishStat *FinishStat, h *env.Dbhandler) error {
	
	// Mongodb
	db := h.session.Copy()
	defer db.Close()


	taskId := mapKey.TaskId
	date := mapKey.CurDate
	hour := mapKey.HourKey

	//是否有记录
	c := db.DB("liuliang").C("task_finish_stat")
	var oldStat TaskFinishStat
	err := c.Find(bson.M{
		"task_id" : taskId,
		"date" : date,
	}).One(&oldStat)
	fmt.Printf("value:%v, err:%v \r\n", oldStat, err)
	if err != nil {
		if err.Error() != "not found" {
			return errors.Wrap(err, "cant find stat ")			
		} else {
			// 如果不存在新建
			oldStat.AddDtime = utils.GetFormattedDateTime("min")
			oldStat.TaskId = taskId
			oldStat.Date = date
			oldStat.UpdateTime = utils.GetCurrentTimestampSec()

			err = c.Insert(&oldStat)
			if err != nil {
				return errors.Wrap(err, "create new stat error: ")
			}
		}
	}
	
	//更新数据
	for key, count := range finishStat.FinishList5M {
		 _, ok := oldStat.FinishList5M[key]
		 if !ok {
			 oldStat.FinishList5M[key] = count
		 } else {
			 oldStat.FinishList5M[key] += count
		 }
	}
	for key, count := range finishStat.FinishListHour {
		_, ok := oldStat.FinishListHour[key]
		if !ok {
			oldStat.FinishListHour[key] = count
		} else {
			oldStat.FinishListHour[key] += count
		}
	}

	oldStat.LastClickDtime = utils.GetFormattedDateTime("min")

	oldStat.FinishPvNum += finishStat.FinishPvNum 

	// 每天总ip数
	dailyLength, err  := TaskDailyHashIpLength(h.Cache, taskId, date)
	if err != nil {
		return errors.Wrap(err, "dailyLength get err: ")
	}
	oldStat.FinishIpNum =  dailyLength 

	// 每小时新ip数
	hourlyLength, err := TaskHourlyHashIpLength(h.Cache, taskId, hour)
	if err != nil {
		return errors.Wrap(err, "hourlyLength get err: ")
	}
	oldStat.FinishListIpHour[hour]=hourlyLength 



	return nil
}