package projects

import (
	"bytes"
	"fmt"
	"sync"
	"time"

	"github.com/fatih/color"
)

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

type FinishOrder struct {
	list []string
	mu   *sync.Mutex
}

var (
	philosophers = []Philosopher{
		{name: "Plato", leftFork: 4, rightFork: 0},
		{name: "Socrates", leftFork: 0, rightFork: 1},
		{name: "Aristotle", leftFork: 1, rightFork: 2},
		{name: "Democritus", leftFork: 2, rightFork: 3},
		{name: "Epicurus", leftFork: 3, rightFork: 4},
	}
	hunger           = 3
	chewingTime      = time.Second * 1
	philosophingTime = time.Second * 1
	finishOrder      = FinishOrder{
		list: make([]string, 0, len(philosophers)),
		mu:   &sync.Mutex{},
	}
)

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	forks := make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < len(philosophers); i++ {

		go diningProblem(philosophers[i], wg, forks, seated, &finishOrder)

	}

	wg.Wait()

	var resultOrderBytes bytes.Buffer

	for _, v := range finishOrder.list {
		resultOrderBytes.WriteString(v + " ")
	}

	color.Green("---------------------------------\nfinish dinner order:")
	color.Green(string(resultOrderBytes.String()))
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup, finishOrder *FinishOrder) {
	defer wg.Done()

	color.Blue("%s sat at the table\n", philosopher.name)

	seated.Done()

	seated.Wait()

	for i := hunger; i > 0; i-- {

		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork\n", philosopher.name)

		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork\n", philosopher.name)

		}

		fmt.Printf("\t%s has both forks and eating\n", philosopher.name)
		time.Sleep(chewingTime)

		fmt.Printf("\t%s is thinking...\n", philosopher.name)
		time.Sleep(philosophingTime)

		fmt.Printf("\t%s put down both forks\n", philosopher.name)

		forks[philosopher.rightFork].Unlock()
		forks[philosopher.leftFork].Unlock()
	}

	finishOrder.mu.Lock()
	finishOrder.list = append(finishOrder.list, philosopher.name)
	finishOrder.mu.Unlock()

	color.Green("%s is satisfied\n", philosopher.name)
	color.Green("%s left the table\n", philosopher.name)
}

func DiningPhilosophers() {
	color.Cyan("Start of the solution\n----------------------")
	color.Cyan("The table is empty\n\n")

	dine()

	color.Cyan("\nThe table is empty")
	color.Cyan("End of the solution")
}
