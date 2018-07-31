package cronjob

import (
	"github.com/eggegg/Doraemon/helpers/env"	
	"github.com/eggegg/Doraemon/models"	
	
	"log"	

	_"gopkg.in/mgo.v2"
	
)


//---------------------
// 加载数据到缓存，定时运行
//---------------------
func AdInitialLoadExecutor(h *env.Dbhandler) error {
	log.Println("conjob:AdInitialLoadExecutor start")

	session := h.Session.Copy()
	defer session.Close()

	conn := h.Cache.Pool.Get()
	defer conn.Close()

	var err error

	var taskTimelineLists []models.TaskTimelineList
	taskTimelineLists, err = models.GetAllTaskTimelineListFromDb(session)
	if err != nil {
		return err
	}

	// log.Printf("num of task_timeline_list: %v", len(taskTimelineLists))

	var finalTaskTimelineLists []models.TaskTimelineList

	if len(taskTimelineLists) > 0 {
		for _, taskTimeLineList := range taskTimelineLists {
			taskId := taskTimeLineList.TaskId
			// log.Println("task_id:", taskId)
			var taskList models.TaskList
			taskList, err = models.GetTaskListById(session,taskId)
			if err != nil {
				continue
			}

			// log.Printf("task_list : %v", taskList)

			// condition check
			if taskList.Stat != 1 || taskList.FinishStat == 1{
				continue
			}

			if taskList.FinishPvNum > taskList.TaskIpNum {
				models.SetTaskListFinish(session, taskId)
			}

			finalTaskTimelineLists = append(finalTaskTimelineLists, taskTimeLineList)
		}

		// write to redis
		err = models.RedisSaveTaskTimelineLists(h.Cache, finalTaskTimelineLists)


	}

	log.Println("conjob:AdInitialLoadExecutor end")

	return nil
}