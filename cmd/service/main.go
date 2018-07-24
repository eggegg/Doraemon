package main

import (
        "time"
        "os"
        "os/signal"
        "context"
        "github.com/eggegg/Doraemon/handlers"
        "github.com/eggegg/Doraemon/bindings"        
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
	e.Validator = new(bindings.Validator)        
      
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
        

        // Start server with GraceShutdown
        go func() {
                if err := e.Start(":8080");err != nil {
                        e.Logger.Info("shutting down the server.")
                }
        }()    

        // Wait for interrupt signal to gracefuly shutdown the server with 
        // a timeout of 10 seconds.
        quit := make(chan os.Signal)
        signal.Notify(quit, os.Interrupt, os.Kill)

        <- quit

        ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
        defer cancel()
        if err := e.Shutdown(ctx); err != nil {
                e.Logger.Fatal(err)
        }
	
}