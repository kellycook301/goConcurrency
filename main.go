package main

import "fmt"

func main() {
	// first goRoutine example
	// c := make(chan string)
	// go count("sheep", c)

	// for msg := range c {
	// 	fmt.Println(msg)
	// }

	// second goRoutine example
	// c := make(chan string, 3)
	// c <- "hello"
	// c <- "world"
	// c <- "huh?"

	// msg := <-c
	// fmt.Println(msg)

	// msg = <-c
	// fmt.Println(msg)

	// msg = <-c
	// fmt.Println(msg)

	// third goRoutine example
	// c1 := make(chan string)
	// c2 := make(chan string)

	// go func() {
	// 	for {
	// 		c1 <- "Every 500ms"
	// 		time.Sleep(time.Millisecond * 500)
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		c2 <- "Every two seconds"
	// 		time.Sleep(time.Millisecond * 2000)
	// 	}
	// }()

	// for {
	// 	select {
	// 	case msg1 := <-c1:
	// 		fmt.Println(msg1)
	// 	case msg2 := <-c2:
	// 		fmt.Println(msg2)
	// 	}
	// }

	// fourth goRoutine example that takes up a LOT of CPU hahahaha
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

// in relation to first goRoutine
// func count(thing string, c chan string) {
// 	for i := 1; i <= 5; i++ {
// 		c <- thing
// 		time.Sleep(time.Millisecond * 10)
// 	}

// 	close(c)
// }
