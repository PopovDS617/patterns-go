package projects

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	hunger = 0
	chewingTime = time.Second * 0
	philosophingTime = time.Second * 0

	for i := 0; i < 10; i++ {
		finishOrder = FinishOrder{
			list: []string{},
			mu:   &sync.Mutex{},
		}

		dine()
		if len(finishOrder.list) != 5 {
			t.Errorf("expected 5, got %v", len(finishOrder.list))
		}
	}

}

func Test_dineWithVaryingDelays(t *testing.T) {
	var tests = []struct {
		name  string
		delay time.Duration
	}{
		{"zero delay", time.Second * 0},
		{"quarter second delay", time.Millisecond * 250},
		{"half second delay", time.Millisecond * 500},
	}

	for _, e := range tests {
		hunger = 5
		chewingTime = e.delay
		philosophingTime = e.delay
		fmt.Println(e.name)

		finishOrder = FinishOrder{
			list: []string{},
			mu:   &sync.Mutex{},
		}

		dine()

		if len(finishOrder.list) != 5 {
			t.Errorf("%s: expected 5, got %v", e.name, len(finishOrder.list))
		}
	}
}
