package main

import (
	"app/concurrency/projects"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	// functions.Calc()
	// interfaces.Football()
	// pointers.Pointer1()
	// slicespkg.Slice1()
	// concurrency.FetchLineByLine()
	// concurrency.PromiseAllLikeFunctionWithWG()
	// concurrency.PromiseAllLikeFunctionWithUnbuffCh()
	// concurrency.PromiseAllLikeFunctionWithBuffCh()
	// concurrency.Pipeline()
	// internals.CaptLoopVars()
	// users := concurrency.MakeUsersList(50)
	// semaphore.DeactivateUsersSemaphore(users, 10)
	// workerpool.DeactivateUsersWorkerPool(users, 10)
	// semaphore.Semaphore2()
	// workerpool.WorkerPool2()
	// concurrency.Select()
	// concurrency.WithErrorGroup()
	// concurrency.ContextConcurrencyTimeout()
	// faninfanout.FanInFanOutExample2()
	// faninfanout.FanInFanOutExampleEx()
	// pool.SyncPoolExample()
	// once.OnceExample()
	// datarace.DataRace()
	// datarace.DataRaceEliminated()
	// workerpool.WorkerPool3()
	// projects.IncomeCalculator()
	// projects.ProducerConsumerProblem()
	// projects.DiningPhilosophers()
	// projects.ArmyCommunication()
	// projects.SleepingBarber()
	// projects.ChatServerMain()
	// res, err := projects.HTTPGetBody("https://google.com")

	// response, ok := res.([]byte)
	// if !ok {
	// 	fmt.Println("type assertion error")
	// }

	// fmt.Println(len(response), err)

	urls := []string{
		"https://google.com", "https://vk.ru", "https://ozon.ru", "https://godoc.org",
		"https://google.com", "https://vk.ru", "https://ozon.ru", "https://godoc.org", "https://google.com", "https://google.com",
	}

	m := projects.New(projects.HTTPGetBody)

	wg := sync.WaitGroup{}

	for _, v := range urls {
		wg.Add(1)
		go func(url string) {

			start := time.Now()

			value, err := m.Get(url)

			if err != nil {
				log.Print("error")
			}

			fmt.Printf("%s %s %d bytes\n", url, time.Since(start), len(value.([]byte)))
			wg.Done()
		}(v)

	}

	wg.Wait()

	fmt.Println("-----------------------------------")

	for _, v := range urls {
		wg.Add(1)
		go func(url string) {

			start := time.Now()

			value, err := m.Get(url)

			if err != nil {
				log.Print("error")
			}

			fmt.Printf("%s %s %d bytes\n", url, time.Since(start), len(value.([]byte)))
			wg.Done()
		}(v)

	}

	wg.Wait()
}
