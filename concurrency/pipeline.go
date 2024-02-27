package concurrency

import "fmt"

func counter(out chan<- int) { // write-only channel
	for x := 0; x < 20; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) { // write-only and read-only channel
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) { // write-only channel
	for v := range in {
		fmt.Println(v)
	}
}

func Pipeline() {

	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)

}
