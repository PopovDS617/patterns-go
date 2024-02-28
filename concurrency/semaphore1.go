package concurrency

import (
	"fmt"
	"time"
)

type Semaphore struct {
	C chan struct{}
}

func (s *Semaphore) Acquire() {
	s.C <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.C
}

func DeactivateUsersSemaphore(users []User, gCount int) ([]User, error) {

	sem := Semaphore{
		C: make(chan struct{}, gCount),
	}

	outputCh := make(chan ResultWithError, len(users))
	signalCh := make(chan struct{})

	output := make([]User, 0, len(users))

	for _, v := range users {

		go func(user User) {
			fmt.Println("waiting", user.ID)
			sem.Acquire()       // let the first batch in
			defer sem.Release() // clean one item from buffer

			fmt.Println("passed semaphore", user.ID)
			time.Sleep(time.Second * 2) // go routine doing something heavy

			err := user.Deactivate()

			select {
			case outputCh <- ResultWithError{
				User: user,
				Err:  err,
			}:
			case <-signalCh:
				return

			}

		}(v)

	}

	for i := len(users); i > 0; i-- {
		res := <-outputCh

		if res.Err != nil {
			close(signalCh)
			return nil, fmt.Errorf("an error occured: %w", res.Err)
		}

		output = append(output, res.User)

	}

	fmt.Println(output)

	return output, nil
}
