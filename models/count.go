package models

import (
	"time"
	"sync"
	"log"
	"github.com/eggegg/Doraemon/helpers/env"
	"github.com/eggegg/Doraemon/utils"
	
	
	"encoding/json"
	"bytes"
	
	"math/rand"
	
)


func MockQueueData(h *env.Dbhandler) error {
	//读取redis队列
	conn := h.Cache.Pool.Get()
	defer conn.Close()


	for i:=0; i<100; i++ {

		tempStat := ResultStat {
			Task_id: "12345678",
			Timeline_id: "38384747",
			Dtime: utils.GetCurrentTimestampSec() + rand.Intn(1000),
			Ip: "192.168.10.34",
			Day_new: true,
		}

		tempJson , err := json.Marshal(tempStat)
		if err != nil {
			log.Println("[MockQueue] json encode err", err)
		}

		_, err = conn.Do("rpush", TASK_FINISH_QUEUE, tempJson)
		if err != nil {
			log.Println("[MockQueue] rpush err", err)
		}

	}
	log.Printf("[MockQueue] 100 success")
	

	return nil
}



// 处理队列，并记录到全局变量
func ProcessQueue(countsLock *sync.Mutex, counts *map[MapKey]*FinishStat, h *env.Dbhandler) error {

	//读取redis队列
	conn := h.Cache.Pool.Get()
	defer conn.Close()

	var err error

	reply, err := conn.Do("LPOP", TASK_FINISH_QUEUE)
	if err != nil {
		return err
	}

	//队列已经清空，停留一秒再返回继续循环
	if reply == nil {
		time.Sleep(1 * time.Second)		
		return nil
	}

	var resultStat ResultStat	
	decoder := json.NewDecoder(bytes.NewReader(reply.([]byte)))
	if err := decoder.Decode(&resultStat); err != nil {
		return err
	}

	log.Printf("[GetJob] %v, task_id:%v \r\n", resultStat, resultStat.Task_id)
	
	countsLock.Lock()
	// defer countsLock.Unlock()
	defer func ()  {
		log.Println("unlock")
		countsLock.Unlock()
	}()


	//是否需要初始化
	if *counts == nil {
		*counts = make(map[MapKey]*FinishStat)
		log.Printf("init: %v", counts)
		
	}

	//获取当前的keymap
	dayKey := MapKey{
		TaskId: resultStat.Task_id,
		TimelineId: resultStat.Timeline_id,
		CurDate: utils.GetDateTimeFromTimeStamp(resultStat.Dtime, "day"),
		HourKey: utils.GetHourKey(resultStat.Dtime),
		CurKey: utils.GetCur5MinKey(resultStat.Dtime),
	}
	log.Printf("daykey: %v", dayKey)
	_, ok := (*counts)[dayKey]
	if !ok {
		(*counts)[dayKey] = &FinishStat{
			FinishList5M: make(map[string]int),
			FinishListHour: make(map[string]int),
		}
	}

	//------------------------------
	//   修改Redis数据
	//------------------------------
	

	//------------------------------
	//   修改全局统计变量
	//------------------------------

	(*counts)[dayKey].FinishPvNum++
	if resultStat.Day_new {
		(*counts)[dayKey].FinishIpNum++		
	}

	cur5Key := utils.GetCur5MinKey(resultStat.Dtime)
	hourKey := utils.GetHourKey(resultStat.Dtime)
	(*counts)[dayKey].FinishList5M[cur5Key]++
	(*counts)[dayKey].FinishListHour[hourKey]++

	// log.Printf("stat: %v", *counts)

	return nil
}


//批量处理全局变量，保存至数据库
func DoCount(countsLock *sync.Mutex, counts *map[MapKey]*FinishStat, h *env.Dbhandler)  error {
	countsLock.Lock()
	defer countsLock.Unlock()
	if len(*counts) == 0 {
		log.Println("No new count, skipping database update")
		return nil
	}
	log.Println("Updating database...")
	log.Println(*counts)
	ok := true

	for mapKey, finishStat := range *counts {
		log.Printf("key:%v, value:%v", mapKey, finishStat)

		//更新mongo表
		err := updateTaskFinishStat(mapKey, finishStat, h)		
		if err != nil {
			log.Println("update task_finish_stat err:", err)
		}

	}

	if ok {
		log.Println("Finished updating database...")
		*counts = nil // reset counts
	}

	return nil
}