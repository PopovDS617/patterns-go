package projects

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Receiver an order number %d\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("Making pizza #%d. It will take %d seconds....\n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The oven broke during cooking pizza #%d!", pizzaNumber)
		} else {
			success = true

			msg = fmt.Sprintf("Pizza order #%d is ready", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p
	}

	return &PizzaOrder{pizzaNumber: pizzaNumber}

}

func pizzeria(pizzaMaker *Producer) {
	i := 0

	for {
		currentPizza := makePizza(i)

		if currentPizza != nil {
			i = currentPizza.pizzaNumber

			select {
			case pizzaMaker.data <- *currentPizza:

			case quitCh := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitCh)
				return
			}
		}
	}
}

func ProducerConsumerProblem() {

	color.Cyan("The Pizzeria is open\n---------------------------------------------")

	pizzaJob := &Producer{data: make(chan PizzaOrder), quit: make(chan chan error)}

	go pizzeria(pizzaJob)

	// listen to producer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery!", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer would be mad")
			}
		} else {
			color.Cyan("Done making pizzas!")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error closing channel!", err)
			}
		}
	}
	color.Cyan("Done for the day\n---------------------------------")
	color.Cyan("Succesfully made %d pizzas. Failed to make %d. Total attempts %d.", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		color.Red("It was and awful day...")
	case pizzasFailed >= 6:
		color.Yellow("It was not a good day...")
	case pizzasFailed >= 3:
		color.Yellow("We could have done better...")
	case pizzasFailed >= 0:
		color.Green("Nicely done!")

	}
}
