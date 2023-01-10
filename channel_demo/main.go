package main

import (
	"fmt"
	"sync"
)

func waitTest(wg *sync.WaitGroup, c chan int) {
	// 減counter
	defer wg.Done()
	num := <-c
	num++
	c <- num
}
func main() {
	// // 遍歷前須關閉
	// intChan := make(chan int, 3)
	// intChan <- 10
	// intChan <- 20
	// intChan <- 30
	// close(intChan)
	// for v := range intChan {
	// 	// num := <-intChan
	// 	fmt.Println(v)
	// }
	// fmt.Println(len(intChan))

	// 讀不到也要close
	// intChan2 := make(chan int, 3)
	// if v, ok := <-intChan2; ok {
	// 	fmt.Println(v, ok)
	// } else {
	// 	fmt.Println(v, ok)
	// }

	// 同步攜程
	var wg sync.WaitGroup
	c := make(chan int)
	// 加counter
	wg.Add(1)
	go waitTest(&wg, c)
	wg.Add(1)
	go waitTest(&wg, c)
	c <- 5
	// 等counter
	wg.Wait()
	close(c)
	fmt.Println(<-c)

}
