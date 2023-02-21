package main

import (
	"fmt"
	"time"
)

func timer() {
	start := time.Now() // contains both the wall and monotonic clocks
	fmt.Println("start time: ", start)
	time.Sleep(1 * time.Second)
	elapsed := time.Until(start)
	fmt.Println("elapsed time: ", elapsed)
}
