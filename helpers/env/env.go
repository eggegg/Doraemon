package env


import (
	"github.com/eggegg/Doraemon/utils"	
	mgo "gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
	
)

type Dbhandler struct {
	Cache *utils.Cache
	Session *mgo.Session
}

func CreateHandler(db *utils.Cache, session *mgo.Session) *Dbhandler {
	return &Dbhandler{db, session}
}