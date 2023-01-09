package main

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

}
