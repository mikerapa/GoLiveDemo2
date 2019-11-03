package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getData(inChan chan int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < rand.Intn(25); i++ {
		inChan <- rand.Intn(10)
	}
	close(inChan)
}

func main() {
	inChan := make(chan int)

	go getData(inChan)
	for newInt := range inChan {
		fmt.Println("new int:", newInt)
	}
}
