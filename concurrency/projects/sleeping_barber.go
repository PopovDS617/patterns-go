package projects

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/fatih/color"
)

var barbershopCapacity = 6
var clientArrivalRate = 200
var workDurationPerClient = 1000 * time.Millisecond
var barbershopWorkingHours = 10 * time.Second
var numberOfBarbers = 0

type Barbershop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	isOpen          bool
	servedClients   int
	mu              sync.Mutex
}

func (b *Barbershop) addServedClient() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.servedClients++

}

func (b *Barbershop) served() int {
	return b.servedClients
}

func (b *Barbershop) Open() {
	b.isOpen = true
}
func (b *Barbershop) Close() {
	color.Yellow("Closing barbershop")
	b.isOpen = false
	close(b.ClientsChan)

	for a := 1; a <= b.NumberOfBarbers; a++ {
		<-b.BarbersDoneChan
	}

	close(b.BarbersDoneChan)

	color.Yellow("--- Barbershop is closed")
}

func (b *Barbershop) IsOpen() bool {
	return b.isOpen
}

func NewBarbershop(shopCapacity int, hairCutDuration time.Duration, numberOfBarbers int, barbersDoneChan chan bool, clientsChan chan string) *Barbershop {
	return &Barbershop{
		ShopCapacity:    shopCapacity,
		HairCutDuration: hairCutDuration,
		NumberOfBarbers: numberOfBarbers,
		BarbersDoneChan: barbersDoneChan,
		ClientsChan:     clientsChan,
	}
}

func (b *Barbershop) startShift(barber string) {
	b.NumberOfBarbers++

	go func() {
		isSleeping := false

		for {
			color.Cyan("%s goes to the waiting rooms to check for clients", barber)
			if len(b.ClientsChan) == 0 {
				color.Yellow("zzz - no clients, %s takes a nap", barber)
				isSleeping = true
			}

			client, ok := <-b.ClientsChan

			if ok {
				if isSleeping {
					color.Yellow("%s wakes %s up", client, barber)
					isSleeping = false
				}
				b.cutHair(barber, client)
			} else {
				b.finishShift(barber)
				return
			}

		}

	}()
}
func (b *Barbershop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair", barber, client)
	time.Sleep(b.HairCutDuration)
	color.Green("%s is finished cutting %s's hair", barber, client)
	b.addServedClient()
}

func (b *Barbershop) finishShift(barber string) {
	color.Cyan("%s is going home, shift is over", barber)
	b.BarbersDoneChan <- true
}

func (b *Barbershop) addClient(client string) {
	color.Green("*** %s arrives", client)

	if b.isOpen {
		select {
		case b.ClientsChan <- client:
			color.Yellow("%s takes a seat in the waiting room", client)
		default:
			color.Red("The waiting rooms is full, so %s leaves", client)
		}

	} else {
		color.Red("The barbershop is already closed, so %s leaves", client)
	}
}

func SleepingBarber() {
	color.Cyan("Sleeping barber problem solution")
	color.Cyan("--------------------------------")

	clientChan := make(chan string, barbershopCapacity)
	doneChan := make(chan bool)
	barbershopClosing := make(chan bool)
	closed := make(chan bool)

	barbershop := NewBarbershop(barbershopCapacity, workDurationPerClient, numberOfBarbers, doneChan, clientChan)
	barbershop.Open()
	color.Cyan("Barbershop is open for the day!")
	color.Cyan("--------------------------------")

	barbershop.startShift("Bob")
	barbershop.startShift("Maggie")
	barbershop.startShift("Frank")

	go func() {
		<-time.After(barbershopWorkingHours)
		barbershopClosing <- true
		barbershop.Close()
		closed <- true
	}()

	i := 1

	go func() {
		for {
			randomMilliseconds := rand.Int() % (2 * clientArrivalRate)

			select {
			case <-barbershopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilliseconds)):
				barbershop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}

		}
	}()

	<-closed

	color.Green("$$$ - Today we served %d clients", barbershop.served())

}
