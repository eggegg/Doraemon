package models

import (
	"gopkg.in/mgo.v2/bson"	
)

type TaskFinishStat struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"` 	
}