package projects

import (
	"sync"
	"time"

	"github.com/fatih/color"
)

type Coordinates struct {
	X float64
	Y float64
}

type Order struct {
	Corps       string
	Coordinates Coordinates
}

func RadioMan(generalCh <-chan Order, tankBatallionCh chan<- Coordinates, artilleryCh chan<- Coordinates, quitCh <-chan struct{}, batallionQuitCh chan<- struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case msg := <-generalCh:
			color.Yellow("received order from general, sending it to %s batallion", msg.Corps)
			if msg.Corps == "tank" {

				tankBatallionCh <- msg.Coordinates
			}
			if msg.Corps == "artillery" {
				artilleryCh <- msg.Coordinates
			}
		case <-quitCh:
			color.Yellow("readio man received a cease fire order! All units must return to the base")
			batallionQuitCh <- struct{}{}
			batallionQuitCh <- struct{}{}
			return

		}
	}

}

func TankBatallion(tankBatCh <-chan Coordinates, quitCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case msg := <-tankBatCh:
			color.Green("-------------------------\n    .--._____\n.-='=='==-,\n(O_o_o_o_o_O)\n-------------------------\ntank batallion received order to move and defend this coordinates - x:%.2f, y:%.2f\n-------------------------", msg.X, msg.Y)
		case <-quitCh:
			color.Green("tank batallion received order to return to the base")
			return

		}

	}
}

func ArtilleryBatalliion(artBatCh <-chan Coordinates, quitCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case msg := <-artBatCh:
			color.Red("-------------------------\n      -.     \n       |========  \n-========-\n-------------------------\nartillery batallion received order to attack this coordinates - x:%.2f, y:%.2f\n-------------------------", msg.X, msg.Y)
		case <-quitCh:
			color.Red("artillery batallion received a ceasefire order")
			return

		}

	}
}

func ArmyCommunication() {
	tankBatCh := make(chan Coordinates)
	artBatCh := make(chan Coordinates)
	generalCh := make(chan Order)
	radioQuitCh := make(chan struct{})
	batallionQuitCh := make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(3)

	go RadioMan(generalCh, tankBatCh, artBatCh, radioQuitCh, batallionQuitCh, &wg)
	go TankBatallion(tankBatCh, batallionQuitCh, &wg)
	go ArtilleryBatalliion(artBatCh, batallionQuitCh, &wg)

	time.Sleep(1 * time.Second)
	generalCh <- Order{Corps: "tank", Coordinates: Coordinates{X: 245.2, Y: 114.2}}
	time.Sleep(1 * time.Second)
	generalCh <- Order{Corps: "artillery", Coordinates: Coordinates{X: 245.2, Y: 114.2}}

	time.Sleep(2 * time.Second)
	radioQuitCh <- struct{}{}

	wg.Wait()

	close(tankBatCh)
	close(artBatCh)
	close(generalCh)
	close(radioQuitCh)
	close(batallionQuitCh)

}
