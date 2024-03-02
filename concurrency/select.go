package concurrency

import (
	"fmt"
	"math/rand/v2"
	"time"
)

var firstNames = []string{"Anna", "Alexander", "Arnold", "Alexei", "Samuel", "Jonathan"}
var lastNames = []string{"Johnson", "Peterson", "Barren", "Steele"}

func FirstNameProcess(ch chan<- string) {

	index := rand.IntN(len(firstNames) - 1)
	ch <- firstNames[index]
}

func FullNameProcess(readCh <-chan string, writeCh chan<- string) {

	firstName, ok := <-readCh

	index := rand.IntN(len(lastNames) - 1)

	if ok {
		writeCh <- firstName + " " + lastNames[index]

	}

}

func Closer(doneCh chan<- struct{}) {
	time.Sleep(time.Second * 1)

	doneCh <- struct{}{}

}

func Select() {
	firstNameCh := make(chan string)
	fullNameCh := make(chan string)
	doneCh := make(chan struct{})

	for i := 20; i > 0; i-- {
		go FullNameProcess(firstNameCh, fullNameCh)
		go FirstNameProcess(firstNameCh)

	}
	go Closer(doneCh)
	for {
		select {
		case firstName := <-firstNameCh:
			fmt.Printf("reading from channel, firstname is %s\n", firstName)

		case fullName := <-fullNameCh:
			fmt.Printf("reading from channel, fullname is %s\n", fullName)
		case <-doneCh:
			return

		}
	}

}
