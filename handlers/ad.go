package handlers

import (
	"net/http"
	"github.com/labstack/echo"

	"github.com/eggegg/Doraemon/bindings"
	"github.com/eggegg/Doraemon/renderings"
	"github.com/eggegg/Doraemon/middlewares"
	
	uuid "github.com/satori/go.uuid"
	
)

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
	//db := c.Get(models.DBContextKey).(*sql.DB)
	
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