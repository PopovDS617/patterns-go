package workerpool

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func parallelDownload(ctx context.Context, urls <-chan string, numWorkders int) map[string]string {

	ctx, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*20))

	defer cancel()

	var wg sync.WaitGroup

	result := map[string]string{}

	var mx sync.Mutex

	for i := 0; i < numWorkders; i++ {

		wg.Add(1)

		go func(wg *sync.WaitGroup, i int) {
			time.Sleep(time.Duration(rand.Intn(int(time.Millisecond * 200))))

			defer wg.Done()

			for {
				select {
				case url, ok := <-urls:

					if !ok {
						return
					}
					mx.Lock()
					result[url] = url + " wow"
					mx.Unlock()

				case <-ctx.Done():
					return
				}

			}

		}(&wg, i)

	}

	wg.Wait()

	return result
}

func WorkerPool3() {

	urls := make(chan string)

	go func() {

		urls <- "site1"
		urls <- "site2"
		urls <- "site3"
		urls <- "site4"
		urls <- "site5"
		urls <- "site6"
		urls <- "site7"
		urls <- "site8"
		urls <- "site9"
		urls <- "site10"
		urls <- "site11"
		urls <- "site12"

		close(urls)
	}()

	result := parallelDownload(context.Background(), urls, 3)

	for _, v := range result {
		fmt.Printf("%s\n", v)
	}

}
