package main

import (
	"fmt"
)

// Concurreny in Go

/*
	a. Sequential vs Concurrent vs Parallel task
	b. go-routine program, main go-routine, anonymous go-routine
	c. Go runtime schuduler
	d. Wait Groups and examples
	e. Channels
	f. buffered, unbuffered, Reading and writing channels
	g. Select statement
	h. Mutex lock
	i. Concurreny practise


	Home Work -
	1. ping, pong -> 2 goroutine - > ping, pong, ping , pong -> infinite , only 5 time
	2. Producer consumer problem- > buff channels - > 1 producer and 2 consuer.
*/

/*
func main() {
	fmt.Println("Starting main...")

	var wg sync.WaitGroup

	wg.Add(2)
	go routineTask1(&wg)
	go routineTask2(&wg)
	go routineTask1(&wg)

	fmt.Println("before wait...")
	wg.Wait()

	// for i:= 0; i<1000; i++{
	// 	go routineTask()
	// }

	// fmt.Println("inside main....")

	// // anonymous go-routine
	// go func() {
	// 	fmt.Println("anynomos goroutine")
	// }()

	// time.Sleep(10 * time.Millisecond)

	fmt.Println("ending main...")
}

func routineTask1(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("routine task1")

}

func routineTask2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("routine task2")

}
*/

// ch -> queue  blocking queue //

/*
func main() {

	// slie := []{2,3,5}

	var lock sync.Mutex

	lock.Lock()
	defer lock.Unlock()
	// name := "veltris"
	ch := make(chan string) //  1 in and one out

	// ch_buf := make(chan string, 3) // 3 goies in an

	go routinetask(ch, lock)
	fmt.Println("before the channel reciver")
	msg := <-ch // get the value out of channel
	fmt.Println("message from rotune : ", msg)
	fmt.Println("ending main...")

	close(ch)
}

func routinetask(ch chan<- string, lock sync.Mutex) {
	// time.Sleep(2 * time.Second)
	lock.Lock()
	defer lock.Unlock()

	fmt.Println("in ananoyms channel")
	ch <- "Hello form routine channel to veltris" // Put the value to channel
}
*/

func main() {
	fmt.Println("starting main...")

	shared_Ch := make(chan int, 3)

	// producer
	go producer(shared_Ch)

	// consumer
	for val := range shared_Ch {

		fmt.Println("from consumer", val)
	}

}

func producer(shared_Ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("in producer", i)
		shared_Ch <- i
	}

	// Close the channel to indicate that no more values will be sent
	close(shared_Ch)
}
