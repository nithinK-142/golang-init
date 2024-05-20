package main

import (
	"fmt"
	"time"
)

func main() {
	currTime := time.Now()
	fmt.Println(currTime)

	t := time.Now()

	fmt.Println(t.Day(), t.Month(), t.Year(), "", t.Hour(), "h", ":", t.Minute(), "m", ":", t.Second(), "s")

	// go's crazy time formating
	fmt.Println(currTime.Format("01-02-2006 15:04:05 Monday"))

	// create date
	createdDate := time.Date(2020, time.August, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
}
