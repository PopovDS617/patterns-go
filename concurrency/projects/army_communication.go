package projects

import (
	"sync"

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

func RadioMan(generalCh <-chan Order, tankBatallionCh chan<- Coordinates, artilleryCh chan<- Coordinates, wg *sync.WaitGroup) {
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
		default:
		}
	}

}

func TankBatallion(tankBatCh <-chan Coordinates, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case msg := <-tankBatCh:
			color.Green("-------------------------\n    .--._____\n.-='=='==-,\n(O_o_o_o_o_O)\n-------------------------\ntank batallion received order to move and defend this coordinates - x:%.2f, y:%.2f\n-------------------------", msg.X, msg.Y)
		default:
		}

	}
}

func ArtilleryBatalliion(artBatCh <-chan Coordinates, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case msg := <-artBatCh:
			color.Red("-------------------------\n      -.     \n       |========  \n-========-\n-------------------------\nartillery batallion received order to attack this coordinates - x:%.2f, y:%.2f\n-------------------------", msg.X, msg.Y)
		default:
		}

	}
}

func ArmyCommunication() {
	tankBatCh := make(chan Coordinates)
	artBatCh := make(chan Coordinates)
	generalCh := make(chan Order)
	var wg sync.WaitGroup

	wg.Add(3)

	go RadioMan(generalCh, tankBatCh, artBatCh, &wg)
	go TankBatallion(tankBatCh, &wg)
	go ArtilleryBatalliion(artBatCh, &wg)

	generalCh <- Order{Corps: "tank", Coordinates: Coordinates{X: 245.2, Y: 114.2}}
	generalCh <- Order{Corps: "artillery", Coordinates: Coordinates{X: 245.2, Y: 114.2}}

	wg.Wait()

}
