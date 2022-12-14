package main

import (
	"fmt"
	"sync"
	"time"
)

// ch chan<- int,
func oddNoGenerator(len int, wg *sync.WaitGroup, oddCh chan int, evenCh chan int) {
	fmt.Println("odd one")
	defer wg.Done()

	<-oddCh
	for i := 0; i < len; i++ {

		if (i % 2) != 0 {
			<-oddCh
			fmt.Printf("odd number printed = %d\n", i)
			continue

		}
		evenCh <- i

	}
}

func evenNoGenerator(len int, wg *sync.WaitGroup, oddCh chan int, evenCh chan int) {
	fmt.Println("even one")
	defer wg.Done()

	for i := 0; i < len; i++ {
		if (i % 2) == 0 {
			<-evenCh
			fmt.Printf("even number printed = %d\n", i)
			continue

		}
		oddCh <- i

	}
}

func sameOpGen(len int) {
	for i := 0; i < len; i++ {

		if (i % 2) != 0 {

			fmt.Printf("odd number printed = %d\n", i)
			continue

		}
		fmt.Printf("even number printed = %d\n", i)

	}
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	//var mu sync.Mutex
	evenCh := make(chan int, 1)
	oddCh := make(chan int, 1)

	wg.Add(2)
	evenCh <- 1

	fmt.Println(<-evenCh, "checking channel")
	go oddNoGenerator(1000000, &wg, evenCh, oddCh)
	go evenNoGenerator(1000000, &wg, evenCh, oddCh)

	sameOpGen(1000000)

	//wg.Wait()
	fmt.Println(time.Since(start))

	fmt.Println("main function finished executing")

}
