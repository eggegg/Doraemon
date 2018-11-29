package main

import (
	"testing"
	"fmt"
)


// func TestSimple(t *testing.T)  {

// 	map1 := map[string]int{
//         "one": 1,
//         "two": 2,
//     }

// 	fmt.Printf("before: %v", map1)
	
// 	map1["one"]++

// 	fmt.Printf("after: %v", map1)
// }

func TestMapOps(t *testing.T)  {
	fmt.Println("test...")

	type countData struct {
		name string
		value int
	}

	type FinishStat struct {
		FinishPvNum int
		FinishListHour map[string]int
		FinishList5M map[string]int
		FinishIpNum int
		FinishListIpHour map[string]int
		FinishListAllIpHour map[string]int
		TimelinePv int
		TimelineIp int
	}

	type mapKey struct {
		TaskId string
		CurDate string
	}

	

	curFinishStat := map[mapKey]*FinishStat{}
	dayKey := mapKey{"12345678", "2018-08-03"}
	_, ok := curFinishStat[dayKey]
	if !ok {
		curFinishStat[dayKey] = &FinishStat{
			FinishList5M: make(map[string]int),
			FinishListHour: make(map[string]int),
			FinishListIpHour: make(map[string]int),
			FinishListAllIpHour: make(map[string]int),
		}
	}
	fmt.Printf("init: %v", curFinishStat)
	

	fmt.Printf("Before Finish: %T \r\n", curFinishStat)
fmt.Println("---")


	// if ok == true{
		for i:=0;i<5;i++ {

			hourKey := "test"

			mm, ok := curFinishStat[dayKey].FinishList5M[hourKey];
			fmt.Printf("value in 5m is mm:%T, :%T, isok:%v\r\n", mm, curFinishStat[dayKey].FinishList5M["test"], ok)
			// if !ok {
			// 	curFinishStat[dayKey].FinishList5M["test"] = 0
			// }
			// if ok == true{
				curFinishStat[dayKey].FinishList5M["test"]++
			// }
			fmt.Printf("after ok: %v", curFinishStat[dayKey].FinishList5M)
			fmt.Printf("address: %p", &curFinishStat[dayKey].FinishList5M)
			fmt.Println("---")
		}
	// } 

	fmt.Printf("Final Finish: %v", curFinishStat)
	
}