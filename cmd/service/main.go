package main

import (
	"flag"
        "time"
        "os"
        "os/signal"
        "fmt"
        "context"
        configuration "github.com/eggegg/Doraemon/config"
        "github.com/eggegg/Doraemon/handlers"
        "github.com/eggegg/Doraemon/bindings"        
        "github.com/eggegg/Doraemon/middlewares"
        "github.com/eggegg/Doraemon/models"
        "github.com/eggegg/Doraemon/utils"
        
        // "github.com/eggegg/Doraemon/helpers/cronjob"
        // "github.com/eggegg/Doraemon/helpers/env"
        

        "github.com/labstack/echo"
        "github.com/labstack/echo/middleware"        
        "github.com/labstack/gommon/log"

        mgo "gopkg.in/mgo.v2"
        _ "gopkg.in/mgo.v2/bson"

)

func main() {
        //-------------------
        // Load Configuration file
        //-------------------
        configPath := flag.String("config", "../../config/config.json", "path of the config file")

        flag.Parse()
        // Read config
        config, err := configuration.FromFile(*configPath)
        if err != nil {
                log.Fatal(err)
        }

        //-------------------
        // Create a new echo instance
        //-------------------
        e := echo.New()
	e.Validator = new(bindings.Validator)        
      
        e.Pre(middlewares.RequestIDMiddleware)

        e.Logger.SetLevel(configuration.GetLogLvl(config.LogLevel))
        if config.LogLevel == "DEBUG" {
                fmt.Println("DEBUG MODE")
                e.Debug = true
                e.Use(middleware.Logger())  // logger middleware will “wrap” recovery                
        }
        e.HideBanner = false

        e.Use(middleware.Recover()) // as it is enumerated before in the Use calls

        //-------------------
        // HTML page rendering
        //-------------------
        renderer := handlers.Renderer{
		Debug: false,
        }
        e.Renderer = renderer

        //-------------------
	// Custom middleware
        //-------------------
        
	// Stats
	s := middlewares.NewStats()
	e.Use(s.Process)
        e.GET("/stats", s.Handle) // Endpoint to get stats
        

        //-------------------
	// Add Redis to context
        //-------------------
        redisCache := utils.Cache {
                MaxIdle: 100,
                MaxActive: 100,
                IdleTimeoutSecs: 60,
                Address: fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
        }
        redisCache.Pool = redisCache.NewCachePool()
        e.Use(func (next echo.HandlerFunc) echo.HandlerFunc {
                return func (c echo.Context) error {
                        c.Set(models.RedisContextKey, &redisCache)
                        return next(c)
                }
        })

        //-------------------
	// Mongo
        //-------------------
        host := []string{
		fmt.Sprintf("%s:%s",config.Mongodb.Host, config.Mongodb.Port),
	}
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs: host,
		Direct: true,
		Timeout: 5 * time.Second,
	})
	session.SetMode(mgo.Monotonic, true)
	if err != nil {
		e.Logger.Info("cannot connect to :" + fmt.Sprintf("%s:%s",config.Mongodb.Host, config.Mongodb.Port))
		panic(err)
	}
	defer session.Close()

        //-------------------
	// Route
        //-------------------

        e.File("/", "static/index.html")

        // in order to serve static assets
	e.Static("/static", "static")
        
         // Route 
         e.GET("/health-check", handlers.HealthCheck)
         e.GET("/db-check", handlers.DbStatus)
         e.GET("/error", handlers.Error)

        // V1 Routes
        v1 := e.Group("/v1")

        // Static HTML Page
        v1.GET("/html-index", handlers.HtmlIndex)

        
        // Test 
        v1.GET("/test-get-ad", handlers.GetTestAd)
        v1.GET("/test-ad-click", handlers.TestAdClick)
        v1.GET("/test-ad-show", handlers.TestAdShow)

        v1.GET("/get-ad", handlers.GetAd)
        v1.GET("/ad-click", handlers.AdClick)
        v1.GET("/ad-show", handlers.AdShow)

        //-------------------
	// Cron job load redis to cache
        //-------------------
        // dbhandler := env.CreateHandler(&redisCache, session)
        // go cronjob.Start(dbhandler)
        



        //-------------------
	// Start server with GraceShutdown
        //-------------------
        go func() {
                if err := e.Start(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port));err != nil {
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