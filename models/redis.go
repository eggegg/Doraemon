package models

import (
	"github.com/eggegg/Doraemon/utils"
	"encoding/json"

	redigo "github.com/garyburd/redigo/redis"

	
	// mgo "gopkg.in/mgo.v2"	
	// "gopkg.in/mgo.v2/bson"	
	"github.com/pkg/errors"

	"strings"
	"fmt"
)


const (
	//广告缓存数据
	TASK_TIMELINE_LIST_SET = "task_timeline_list_set"
	TASK_TIMELINE_LIST_AD = "task_timeline_list_"

	TASK_AD_EXPIRE = 60*60*24


	//每日总ip
	TASK_DAILY_IP_SET = "a_"
	//每小时总ip
	TASK_HOURLY_IP_SET = "hourip_"
	
	//日志队列
	TASK_FINISH_QUEUE = "task_finish_queue"
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


//将task_timeline_list数据写入缓存
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


/**

=========   redis ops:   ====================

	hget: a_(task_id)_(cur_s_date)
	changeto:
	hset: a_(task_id)_(cur_s_date), ip, count


	hset: hourip_(cur_hour)_(task_id), ip, time()
	hset: hour_allip_(cur_hour)_(task_id), ip, time()


*/




//--------------------------
// 	 每日ip出现次数：  hset: a_(task_id)_(cur_s_date), ip, count
//--------------------------
func TaskDailyHashKey(taskId string, date string) string {
	return fmt.Sprintf("%s%s%s", TASK_DAILY_IP_SET, taskId, date)
}

//获取每天的ip总数
func TaskDailyHashIpLength(cache *utils.Cache, taskId string, date string) (int, error) {
	conn := cache.Pool.Get()
	defer conn.Close()

	redisKey := TaskDailyHashKey(taskId, date)

	count , err := redigo.Int(conn.Do("HLEN", redisKey))
	if err != nil {
		return 0, err
	}

	return  count, nil
}

//每天的ip增加计数
func TaskDailyHashIpIncr(cache *utils.Cache, taskId string, date string, ip string) error {
	conn := cache.Pool.Get()
	defer conn.Close()

	redisKey := TaskDailyHashKey(taskId, date)

	if _, err := conn.Do("HINCRBY", redisKey, ip, 1); err != nil {
		return errors.Wrap(err, "hincrby err: ")
	}

	return nil
}

//每天的ip出现次数
func TaskDailyHashIpCount(cache *utils.Cache, taskId string, date string, ip string) (int, error){
	conn := cache.Pool.Get()
	defer conn.Close()

	redisKey := TaskDailyHashKey(taskId, date)

	count, err := redigo.Int(conn.Do("HGET", redisKey, ip)); 
	if err != nil {
		return 0, errors.Wrap(err, "hget err: ")
	}
	return count, nil
}


//--------------------------
//  当前小时使用的新ip：  hset: hourip_(cur_hour)_(task_id), ip, time()
//
//  @time 第一次ip命中才更新为当前时间  
//
//--------------------------
func TaskHourlyHashKey(taskId string, hour string) string {
	return fmt.Sprintf("%s%s%s", TASK_HOURLY_IP_SET, hour, taskId)
}

//当前小时的ip总数
func TaskHourlyHashIpLength(cache *utils.Cache, taskId string, date string) (int, error) {
	conn := cache.Pool.Get()
	defer conn.Close()

	redisKey := TaskHourlyHashKey(taskId, date)

	count , err := redigo.Int(conn.Do("HLEN", redisKey))
	if err != nil {
		return 0, err
	}

	return  count, nil
}

//每小时ip出现时间
func TaskHourlyHashIpHit(cache *utils.Cache, taskId string, hour string, ,ip string, ts int) error {
	conn := cache.Pool.Get()
	defer conn.Close()

	redisKey := TaskHourlyHashKey(taskId, hour)

	if _, err := conn.Do("HSET", redisKey, ip, ts); err != nil {
		return errors.Wrap(err, "hset err: ")
	}

	return nil
}


//--------------------------
// 当前小时使用的ip： hset: hour_allip_(cur_hour)_(task_id), ip, time() 
//
// @time 每次ip命中都更新为当前时间
//
//--------------------------



//--------------------------
// 
//--------------------------