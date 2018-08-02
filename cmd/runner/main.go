package main

import (
	"fmt"
	"flag"
	"log"
	"os/signal"
	"time"
	"os"
	"syscall"
	"sync"
	configuration "github.com/eggegg/Doraemon/config"


	"github.com/eggegg/Doraemon/models"
	"github.com/eggegg/Doraemon/utils"

	"github.com/eggegg/Doraemon/helpers/cronjob"
	"github.com/eggegg/Doraemon/helpers/env"
	
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
	// Add Redis to context
	//-------------------
	redisCache := utils.Cache {
		MaxIdle: 100,
		MaxActive: 100,
		IdleTimeoutSecs: 60,
		Address: fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
	}
	redisCache.Pool = redisCache.NewCachePool()

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
		fmt.Printf("cannot connect to :" + fmt.Sprintf("%s:%s",config.Mongodb.Host, config.Mongodb.Port))
		panic(err)
	}
	defer session.Close()


	//-------------------
	// Cron job load redis to cache
	//-------------------
	dbhandler := env.CreateHandler(&redisCache, session)
	go cronjob.Start(dbhandler)


	//-------------------
	// Cron job load redis to cache
	//-------------------
	var counts map[string]int
	var countsLock sync.Mutex


	go func ()  {
		for {
			err := models.ProcessQueue(&countsLock, &counts, dbhandler)
			if err != nil {
				log.Println("DoCalculate err")
				break
			}
		}
		log.Println("End ProcessQueue...")
	}()		
			
		

	log.Println("Waiting for votes on nsq...")
	updateDuration := time.Duration(config.Runner.Duration) * time.Second
	ticker := time.NewTicker(updateDuration)
	 

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	
	//-------------------
	// Shut Down
	//-------------------
	for {
		select {
		case <-ticker.C:
			models.DoCount(&countsLock, &counts, dbhandler)
		case <-termChan:
			ticker.Stop()
			// finished
			log.Println("Stopping....")			
			return
		}
	}

}