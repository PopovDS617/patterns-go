package interfaces

import (
	"fmt"
	"math/rand"
)

type Kicker interface {
	Kick()
}

type Player struct {
	Stamina int
	Power   int
}

type SuperPlayer struct {
	Player
	AdditionalPower int
}

func (p Player) Kick() {
	res := p.Power * p.Stamina
	fmt.Printf("regular player kicks for %d\n", res)
}

func (sp SuperPlayer) Kick() {
	res := sp.Power * sp.Stamina * sp.AdditionalPower
	fmt.Printf("super player kicks for %d\n", res)
}

func Football() {

	var fteam = make([]Kicker, 11)

	for i := 0; i <= 11; i++ {

		fmt.Println(i)

		if i < 10 {
			fteam[i] = Player{
				Stamina: rand.Intn(10),
				Power:   rand.Intn(10)}
		} else {
			fteam[10] = SuperPlayer{
				Player: Player{
					Stamina: rand.Intn(10),
					Power:   rand.Intn(10),
				},
				AdditionalPower: 5,
			}

		}
	}

	for i := range fteam {
		fteam[i].Kick()
	}

}
