package models

import (
	mgo "gopkg.in/mgo.v2"	
	"gopkg.in/mgo.v2/bson"	

	"github.com/pkg/errors"

	"github.com/eggegg/Doraemon/utils"
)

// task_list 
type TaskList struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"` 
	TaskName string `json:"task_name" bson:"task_name"`
	TaskIpNum int `json:"task_ip_num" bson:"task_ip_num"`
	TaskUrl string `json:"task_url" bson:"task_url"`
	TaskSpeHandler string `json:"task_spe_handle" bson:"task_spe_handle"`
	StatReadUrl string `json:"stat_read_url" bson:"stat_read_url"`  
	StatReadExtUrl string `json:"stat_read_ext_url" bson:"stat_read_ext_url"`  
	TaskOpenType string `json:"task_open_type" bson:"task_open_type"`  
	LocalCacheFile string `json:"local_cache_file" bson:"local_cache_file"`  
	StayMinTime int `json:"stay_min_time" bson:"stay_min_time"` 
	StayMaxTime int `json:"stay_max_time" bson:"stay_max_time"`  
	WebPreprareTime int `json:"web_preprare_time" bson:"web_preprare_time"`  
	JsScriptList string `json:"js_script_list" bson:"js_script_list"`  
	IpReuseTime int `json:"ip_reuse_time" bson:"ip_reuse_time"`  
	PvIpBili string `json:"pv_ip_bili" bson:"pv_ip_bili"`  
	SpeFromList string `json:"spe_from_list" bson:"spe_from_list"`  
	SpeType string `json:"spe_type" bson:"spe_type"`  
	SpePhoneGroup string `json:"spe_phone_group" bson:"spe_phone_group"`  
	AddTime int `json:"add_time" bson:"add_time"`  
	UpdateTime int `json:"update_time" bson:"update_time"`  
	Stat int `json:"stat" bson:"stat"`  
	FinishStat int `json:"finish_stat" bson:"finish_stat"`  
	Areas []string `json:"areas" bson:"areas"`  
	FinishPvNum int `json:"finish_pv_num" bson:"finish_pv_num"`  
	FinishIpNum int `json:"finish_ip_num" bson:"finish_ip_num"`  	
	
}

func GetAllTaskListFromDb(session *mgo.Session) ([]TaskList, error)  {
	var taskList []TaskList

	db := session.Copy()
	defer db.Close()

	curTime := utils.GetFormattedDateTime("min")

	c := db.DB("liuliang").C("task_list")
	err := c.Find(bson.M{
		"stat":1,
		"task_begin_time" : bson.M{"$lt": curTime},
		"task_end_time" : bson.M{"$gt": curTime},
	}).All(&taskList)
	if err != nil {
		return taskList, errors.Wrap(err, "mongodb query error: ")
	}

	return taskList, nil
}

// get task_list by id
func GetTaskListById(session *mgo.Session, id string) (TaskList, error) {
	db := session.Copy()
	defer db.Close()

	var taskList TaskList
	c := db.DB("liuliang").C("task_list")
	err := c.FindId(bson.ObjectIdHex(id)).One(&taskList)
	if err != nil {
		return taskList, err
	}

	return taskList, nil
}

// set task_list finish
func SetTaskListFinish(session *mgo.Session, id string) error {
	db := session.Copy()
	defer db.Close()

	c := db.DB("liuliang").C("task_list")
	err := c.Update(bson.M{
		"_id": id,
	}, bson.M{
		"$set": bson.M{
			"finish_stat" : 1,
		},
	})

	if err != nil {
		return err
	}
	return nil

}