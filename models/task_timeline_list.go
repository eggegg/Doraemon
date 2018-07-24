package models

import (
	"gopkg.in/mgo.v2/bson"	
)

// task_timeline_list
type TaskTimelineList struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"` 	
	
}