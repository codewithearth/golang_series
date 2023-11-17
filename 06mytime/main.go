package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdData := time.Date(2020, time.March, 21, 11, 25, 0, 0, time.UTC)

	fmt.Println(createdData)
	fmt.Println(createdData.Format("01-02-2006 Monday  15:04:05"))

}
