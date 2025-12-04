package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Lets play with time")
	presentTime := time.Now()
	fmt.Println("Present Time: ", presentTime)
	fmt.Println(presentTime.Format("02-01-2006 -- Monday -- 3:04:05 MST"))//month (1), day (2), hour (15), for 12 hours(3), minute (4), second (5), and year (2006).
	/// 01 inddicates the month
	// 02 indicates the day
	// 2006 indicates the year
	// 15 indicates the hour in 24-hour format
	// 04 indicates the minutes
	// 05 indicates the seconds
	
	

	createAt := time.Date(2025,time.December,5,2,39,0,0,time.Local);
	fmt.Println("Create At: ", createAt)
	fmt.Println(createAt.Format("02/01/2006 -- Monday -- 03:04:05 PM"))

	/// delay := time.Second * 5
	fmt.Println("Quit after 5 sec")
	for i:=0; i < 5; i++ {
		if(i == 4) {
			fmt.Println("Quit now!")
	
		}else{
		fmt.Println(5-i, " seconds remaining")}
		time.Sleep(1 * time.Second)
	}

}