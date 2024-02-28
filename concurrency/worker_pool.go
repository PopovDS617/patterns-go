package concurrency

import (
	"fmt"
	"sync"
)

func deactivateUser(wg *sync.WaitGroup, inCh <-chan User, outCh chan<- ResultWithError) {
	defer wg.Done()

	for usr := range inCh {
		err := usr.Deactivate()
		outCh <- ResultWithError{
			User: usr,
			Err:  err,
		}
	}
}

func DeactivateUsersWorkerPool(usrs []User, wgCount int) ([]User, error) {
	inputCh := make(chan User)
	outputCh := make(chan ResultWithError)
	wg := &sync.WaitGroup{}

	output := make([]User, 0, len(usrs))

	go func() {
		defer close(inputCh)

		for i := range usrs {
			inputCh <- usrs[i]
		}
	}()

	go func() {
		for i := 0; i < wgCount; i++ {
			wg.Add(1)

			go deactivateUser(wg, inputCh, outputCh)
		}
		wg.Wait()
		close(outputCh)
	}()

	for res := range outputCh {
		if res.Err != nil {
			return nil, fmt.Errorf("an error occurred: %w", res.Err)
		}

		output = append(output, res.User)
	}

	fmt.Println(output)

	return output, nil
}
