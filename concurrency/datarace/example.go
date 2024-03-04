package datarace

import (
	"fmt"
	"sync"
)

func DataRace() {
	var counter int
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(wg *sync.WaitGroup) {
			counter++
			wg.Done()
		}(&wg)
	}
	wg.Wait()
	fmt.Println(counter)
}

func DataRaceEliminated() {
	var counter int
	var wg sync.WaitGroup

	m := sync.Mutex{}

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(wg *sync.WaitGroup) {
			m.Lock()
			defer m.Unlock()
			counter++
			wg.Done()
		}(&wg)
	}
	wg.Wait()
	fmt.Println(counter)
}
