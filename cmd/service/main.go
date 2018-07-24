package main

import (
        "github.com/eggegg/Doraemon/handlers"
        "github.com/eggegg/Doraemon/middlewares"
        _"github.com/eggegg/Doraemon/models"

        "github.com/labstack/echo"
        "github.com/labstack/echo/middleware"        
        "github.com/labstack/gommon/log"
)

func main() {
        // create a new echo instance
        e := echo.New()
        e.Logger.SetLevel(log.DEBUG)

      

        e.Pre(middlewares.RequestIDMiddleware)

        e.Use(middleware.Logger())  // logger middleware will “wrap” recovery
        e.Use(middleware.Recover()) // as it is enumerated before in the Use calls
        
        e.File("/", "static/index.html")

        // in order to serve static assets
	e.Static("/static", "static")
        
         // Route 
         e.GET("/health-check", handlers.HealthCheck)
         e.GET("/error", handlers.Error)

        // V1 Routes
        v1 := e.Group("/v1")

        // Test 
        v1.GET("/test-get-ad", handlers.GetTestAd)
        v1.GET("/test-ad-click", handlers.TestAdClick)
        v1.GET("/test-ad-show", handlers.TestAdShow)

        v1.GET("/get-ad", handlers.GetAd)
        v1.GET("/ad-click", handlers.AdClick)
        v1.GET("/ad-show", handlers.AdShow)
        

        // start the server, and log if it fails             
	e.Logger.Fatal(e.Start(":8080")) 
}