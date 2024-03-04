package faninfanout

import (
	"fmt"
	"math/rand"

	"golang.org/x/sync/errgroup"
)

func fanOut[T any](ch chan T, n int) []chan T {
	chans := make([]chan T, 0, n)

	for i := 0; i < n; i++ {
		chans = append(chans, make(chan T, 1))
	}

	go func() {
		for item := range ch {
			for _, c := range chans {
				c <- item
				// select {
				// case c <- item:
				// case <-time.After(time.Microsecond * 100):

				// }
			}
		}

		for _, c := range chans {
			close(c)
		}
	}()

	return chans
}

func fanIn[T any](chans ...chan T) chan T {
	res := make(chan T)

	var g errgroup.Group

	for _, c := range chans {
		c := c

		g.Go(func() error {
			for s := range c {
				res <- s
			}
			return nil
		})

	}

	go func() {
		g.Wait()
		close(res)
	}()

	return res
}

type ApplicationEx struct {
	name    string
	content string
}

func mockScanEx() string {
	if rand.Intn(100) > 90 {
		return "ALERT - vulterability found"
	}
	return "OK - all fine"
}

func scanSQLInjectionEx(data <-chan ApplicationEx, res chan<- string) error {
	for d := range data {
		res <- fmt.Sprintf("SQL injection scan: %s scanned, result: %s", d.name, mockScanEx())
	}

	close(res)
	return nil

}
func scanTimingExploitsEx(data <-chan ApplicationEx, res chan<- string) error {
	for d := range data {
		res <- fmt.Sprintf("Timing exploits scan: %s scanned, result: %s", d.name, mockScanEx())
	}

	close(res)
	return nil

}
func scanAuthEx(data <-chan ApplicationEx, res chan<- string) error {
	for d := range data {
		res <- fmt.Sprintf("Authentication scan: %s scanned, result: %s", d.name, mockScanEx())
	}

	close(res)
	return nil
}

func FanInFanOutExampleEx() {

	si := []ApplicationEx{
		{name: "ms_comments", content: "package comments"},
		{name: "ms_posts", content: "package posts"},
		{name: "ms_hints", content: "package hints"},
		{name: "ms_video", content: "package video"},
		{name: "ms_sharing", content: "package sharing"},
		{name: "ms_bookmarks", content: "package bookmarks"},
	}

	input := make(chan ApplicationEx, len(si))

	res1 := make(chan string, len(si))
	res2 := make(chan string, len(si))
	res3 := make(chan string, len(si))

	chans := fanOut(input, 3)

	var g errgroup.Group

	g.Go(func() error {
		return scanSQLInjectionEx(chans[0], res1)
	})
	g.Go(func() error {
		return scanTimingExploitsEx(chans[1], res2)
	})
	g.Go(func() error {
		return scanAuthEx(chans[2], res3)
	})

	g.Go(func() error {
		for _, d := range si {
			input <- d

		}

		close(input)
		return nil
	})

	g.Go(func() error {

		res := fanIn(res1, res2, res3)
		for r := range res {
			fmt.Println(r)
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("[OK] scanning is done")
}
