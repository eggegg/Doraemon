package models

import (
	"github.com/eggegg/Doraemon/utils"
	"encoding/json"

	// mgo "gopkg.in/mgo.v2"	
	// "gopkg.in/mgo.v2/bson"	
	"github.com/pkg/errors"

	"strings"
	"fmt"
)


const (
	TASK_TIMELINE_LIST_SET = "task_timeline_list_set"
	TASK_TIMELINE_LIST_AD = "task_timeline_list_"


	TASK_AD_EXPIRE = 60*60*24
)

func MarkResetCacheTime(cache utils.Cache, key string) (int, error){
	conn := cache.Pool.Get()
	defer conn.Close()

	ts := utils.GetCurrentTimestampSec()
	redisKey := utils.GetBgstatRedisKey("resetcache")
	_, err := conn.Do("HSET", redisKey, ts)
	if err != nil {
		return ts, err
	}

	conn.Do("PERSIST", redisKey)

	return ts, nil
}

func RedisSaveTaskTimelineLists(cache *utils.Cache, list []TaskTimelineList) error {
	conn := cache.Pool.Get()
	defer conn.Close()

	if len(list) == 0 {
		return errors.New("task list empty")
	}

	var taskListIds []string
	for _, taskList := range list {
		taskListId := taskList.Id.Hex()
		taskListIds = append(taskListIds, taskListId)

		oneTaskList, err := json.Marshal(taskList)
		if err != nil {
			return errors.Wrap(err, "jsonencode err")
		}

		redisKey := RedisKeyTaskTimeLine(taskListId)
		conn.Send("SET", redisKey, oneTaskList)
		conn.Send("EXPIRE", redisKey, TASK_AD_EXPIRE)
	}

	if len(taskListIds) > 0 {
		conn.Send("DEL", TASK_TIMELINE_LIST_SET)
		for _, id := range taskListIds {
			conn.Send("SADD", TASK_TIMELINE_LIST_SET, id)
		}
		conn.Send("EXPIRE", TASK_TIMELINE_LIST_SET, TASK_AD_EXPIRE)
	}

	fmt.Printf("success load num:(%v) \r\n", len(taskListIds))

	return nil
}

func RedisKeyTaskTimeLine(id string) string  {
	return strings.Join([]string{TASK_TIMELINE_LIST_AD, id}, "")	
}