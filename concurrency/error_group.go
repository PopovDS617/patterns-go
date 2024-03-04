package concurrency

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

func WithErrorGroup() {

	names := []string{"Mark", "Anna", "Mary", "Anthony"}

	eg := errgroup.Group{}

	for _, v := range names {
		v := v

		eg.Go(func() error {

			if len([]rune(v)) == 4 {
				return fmt.Errorf("short name")
			} else {
				fmt.Println(v)
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}
}
