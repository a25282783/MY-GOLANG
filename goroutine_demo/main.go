package main

import (
	"fmt"
	"strconv"
	"time"
)

func test(i int) {
	fmt.Println("test..." + strconv.Itoa(i))
}

func main() {
	for i := 0; i < 5; i++ {
		go test(i)
	}

	time.Sleep(time.Second)
}
