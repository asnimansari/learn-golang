package main

import (
	"fmt"
	"time"
)

func main() {
	newYork, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		panic(err)
	}
	timein := time.Date(1980, 1, 1, 0, 0, 0, 0, newYork)
	timein = timein.Add(1241795629 * time.Second)
	fmt.Println(timein)
}
