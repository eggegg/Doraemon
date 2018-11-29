package handlers

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo"

	"github.com/eggegg/Doraemon/bindings"
	"github.com/eggegg/Doraemon/renderings"
	"github.com/eggegg/Doraemon/middlewares"
	"github.com/eggegg/Doraemon/models"
	"github.com/eggegg/Doraemon/utils"
	
	uuid "github.com/satori/go.uuid"

	redigo "github.com/garyburd/redigo/redis"

	"encoding/json"

	
)


//-------------------
// Check Redis status
//-------------------
func DbStatus(c echo.Context) error {
	resp := renderings.NormalResponse{}
	// get redis from context
	db := c.Get(models.RedisContextKey).(*utils.Cache)
	
	// get redis conn from pool
	conn := db.Pool.Get()
	defer conn.Close()

	testkey := "test_key"
	conn.Send("SET", testkey, "zaker")
	conn.Send("EXPIRE", testkey, 10)

	values, err := redigo.String(conn.Do("GET", testkey))
	if err != nil{
		c.Logger().Error("Redis get err:", err)
		resp.Success = false
		resp.Message = "Unable to bind request for get ad"
		return c.JSON(http.StatusBadRequest, resp)
	}
	c.Logger().Errorf("Redis test get value:%s", values)


	resp.Success = true
	resp.Message = fmt.Sprintf("Redis OK:%s", values)
	return c.JSON(http.StatusOK, resp)
}

func GetAd(c echo.Context) error {
	c.Logger().Debugf("RequestID: %s", c.Get(middlewares.RequestIDContextKey).(uuid.UUID))	
	c.Logger().Info("Get Ad:")
	resp := renderings.NormalResponse{}

	// Get Parameters
	ar := new(bindings.AdRequest)
	if err := c.Bind(ar); err != nil {
		resp.Success = false
		resp.Message = "Unable to bind request for get ad"
		return c.JSON(http.StatusBadRequest, resp)
	}

	if err := c.Validate(ar); err != nil {
		resp.Success = false
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	} 
	c.Logger().Errorf("request: %v",ar)


	// get redis from context
	db := c.Get(models.RedisContextKey).(*utils.Cache)

	//发送成功队列
	finishStat := models.FinishStat {
		Task_id: "5b398f436df0907bce3cffe4",
		Timeline_id : "5b5eeadbbfa809f8e1fa7845",
		Ip: "192.168.10.10",
		Dtime: utils.GetCurrentTimestampSec(),
		Day_new: true,
	}
	oneStatJson, err := json.Marshal(finishStat)
	if err != nil {
		c.Logger().Errorf("finish stat json encode err: %v", err) 
	}

	c.Logger().Printf("FinishStat: %v", oneStatJson)

	// Send to redis
	db.EnqueueValue(models.TASK_FINISH_QUEUE, string(oneStatJson) )

	
	
	// search ad from task timeline list model
	// user, err := models.GetUserByUsername(db, lr.Username)
	// if err != nil {
	// 	resp.Success = false
	// 	resp.Message = "Username or Password incorrect"
	// 	return c.JSON(http.StatusUnauthorized, resp)
	// }

	// return
	resp.Success = true
	resp.Message = "Successfully get ad!"
	return c.JSON(http.StatusOK, resp)
}

func AdShow(c echo.Context) error {
	c.Logger().Debugf("RequestID: %s", c.Get(middlewares.RequestIDContextKey).(uuid.UUID))	
	c.Logger().Info("Ad show")
	resp := renderings.NormalResponse{}
	

	resp.Success = true
	resp.Message = "Ad show!"
	return c.JSON(http.StatusOK, resp)
}

func AdClick(c echo.Context) error {
	c.Logger().Debugf("RequestID: %s", c.Get(middlewares.RequestIDContextKey).(uuid.UUID))	
	c.Logger().Info("Ad click")
	resp := renderings.NormalResponse{}
	

	resp.Success = true
	resp.Message = "Ad click!"
	return c.JSON(http.StatusOK, resp)
}

/*
	Test ad function
*/
func GetTestAd(c echo.Context) error {
	c.Logger().Debugf("RequestID: %s", c.Get(middlewares.RequestIDContextKey).(uuid.UUID))	
	c.Logger().Info("Get Test Ad:")
	resp := renderings.NormalResponse{}
	

	resp.Success = true
	resp.Message = "Successfully get test ad!"
	return c.JSON(http.StatusOK, resp)
}

func TestAdShow(c echo.Context) error {
	c.Logger().Debugf("RequestID: %s", c.Get(middlewares.RequestIDContextKey).(uuid.UUID))	
	c.Logger().Info("Test Ad show")
	resp := renderings.NormalResponse{}
	

	resp.Success = true
	resp.Message = "Test ad show!"
	return c.JSON(http.StatusOK, resp)
}

func TestAdClick(c echo.Context) error {
	c.Logger().Debugf("RequestID: %s", c.Get(middlewares.RequestIDContextKey).(uuid.UUID))	
	c.Logger().Info("Test Ad click")
	resp := renderings.NormalResponse{}
	

	resp.Success = true
	resp.Message = "Test ad click!"
	return c.JSON(http.StatusOK, resp)
}