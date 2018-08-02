package models

import (
	"time"
	"sync"
	"log"
	"github.com/eggegg/Doraemon/helpers/env"
	
)

// func DoCalculate(countsLock *sync.Mutex, counts *map[string]int, h *env.Dbhandler) error {
	
	
// 	for {

// 		if counts == nil {
// 			counts = make(map[string]int)
// 		}

// 		err := processQueue(countsLock, counts, h)
// 		if err != nil {
// 			log.Println("DoCalculate err")
// 			break
// 		}
// 	}
	

// 	return nil
// }

func ProcessQueue(countsLock *sync.Mutex, counts *map[string]int, h *env.Dbhandler) error {
	countsLock.Lock()
	// defer countsLock.Unlock()
	defer func ()  {
		log.Println("unlock")
		countsLock.Unlock()
	}()

	if *counts == nil {
		*counts = make(map[string]int)
	}

	// _, ok := (*counts)["test"]
	// if ok {
	(*counts)["test"]++
	// } 
	
	for key, value := range *counts {
		log.Printf("key:%v, value:%v", key, value)
	}
	
	time.Sleep(1 * time.Second)

	return nil
}

func DoCount(countsLock *sync.Mutex, counts *map[string]int, h *env.Dbhandler)  error {
	countsLock.Lock()
	defer countsLock.Unlock()
	if len(*counts) == 0 {
		log.Println("No new votes, skipping database update")
		return nil
	}
	log.Println("Updating database...")
	log.Println(*counts)
	ok := true

	for key, value := range *counts {
		log.Printf("key:%v, value:%v", key, value)
	}

	// *counts=nil

	if ok {
		log.Println("Finished updating database...")
		*counts = nil // reset counts
	}

	return nil
}