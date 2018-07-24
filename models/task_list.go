package models

import (
	"gopkg.in/mgo.v2/bson"	
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