package models

import (
	mgo "gopkg.in/mgo.v2"		
	"gopkg.in/mgo.v2/bson"	

	"github.com/pkg/errors"
	
	"github.com/eggegg/Doraemon/utils"
)

// task_timeline_list
type TaskTimelineList struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"` 	
	TaskId string `json:"task_id" bson:"task_id"`
	TaskBeginDtime string `json:"task_begin_dtime" bson:"task_begin_dtime"`
	TaskEndDtime string `json:"task_end_dtime" bson:"task_end_dtime"`
	TaskIpNum int `json:"task_ip_num" bson:"task_ip_num"`
	PlanHour map[string]int `json:"plan_hour" bson:"plan_hour"`
	PlanUa map[string]int `json:"plan_ua" bson:"plan_ua"`
	PlanIpNum5M map[string]int `json:"plan_ip_num_5m" bson:"plan_ip_num_5m"`
	MaxOnedayPvNum int `json:"max_oneday_pv_num" bson:"max_oneday_pv_num"`
	PlanIpNumHour map[string]int `json:"plan_ip_num_hour" bson:"plan_ip_num_hour"`
	FinishStat int `json:"finish_stat" bson:"finish_stat"`  
	Stat int `json:"stat" bson:"stat"`  
	AddTime int `json:"add_time" bson:"add_time"`  
	UpdateTime int `json:"update_time" bson:"update_time"` 
	FinishPvNum int `json:"finish_pv_num" bson:"finish_pv_num"`  
	FinishIpNum int `json:"finish_ip_num" bson:"finish_ip_num"` 
}


func GetAllTaskTimelineListFromDb(session *mgo.Session) ([]TaskTimelineList, error)  {
	var taskTimelineList []TaskTimelineList

	db := session.Copy()
	defer db.Close()

	curTime := utils.GetFormattedDateTime("min")

	c := db.DB("liuliang").C("task_timeline_list")
	err := c.Find(bson.M{
		"stat" : 1,
		"task_begin_dtime" : bson.M{"$lt": curTime},
		"task_end_dtime" : bson.M{"$gt": curTime},
	}).All(&taskTimelineList)
	
	if err != nil {
		return taskTimelineList, errors.Wrap(err, "mongodb query error: ")
	}
	return taskTimelineList, nil
}