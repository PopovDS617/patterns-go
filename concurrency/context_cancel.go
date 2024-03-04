package concurrency

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

func queryShardOne(ctx context.Context, query string, shard int, shardRes chan<- string) {

	start := time.Now()

	queryTime := time.After(time.Duration(5+rand.Intn(2)) * time.Second)

	for {

		select {
		case <-ctx.Done():
			fmt.Printf("queryShard: %s on shard %d after %s\n", ctx.Err(), shard, time.Since(start))
			return

		case <-time.After(time.Duration(rand.Intn(1500)+500) * time.Microsecond):
			shardRes <- fmt.Sprintf("queryShard: found an occurence of %s in shard %d at index %d", query, shard, 100000+rand.Intn(899999))

		case <-queryTime:
			fmt.Printf("queryShard: finished query '%s' on shard %d after %s\n", query, shard, time.Since(start))
			return
		}
	}

}

func ContextConcurrencyCancel() {
	var g errgroup.Group

	bgctx := context.Background()
	ctx, cancel := context.WithCancel(bgctx)

	const numShards = 5
	queries := []string{"pid=5543,HTTP_418", "CON_RST"}

	logs := make(chan string)

	start := time.Now()

	for _, query := range queries {
		q := query
		for shard := 0; shard < numShards; shard++ {
			sh := shard

			g.Go(func() error {
				queryShardOne(ctx, q, sh, logs)
				return nil
			})

		}
	}

	limit := 1

	g.Go(func() error {

		for i := 1; i <= limit; i++ {

			res := <-logs
			fmt.Printf("Result %d: %s\n", i, res)

		}
		fmt.Printf("Receiving goroutine: all results received. Query completed afte %s\n", time.Since(start))
		cancel()
		return nil
	})

	g.Wait()

	fmt.Printf("All goroutines finished after %s\n", time.Since(start))

}
