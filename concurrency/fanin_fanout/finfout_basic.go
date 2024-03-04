package faninfanout

import "fmt"

type data int

// запускаем три воркера, которые читают из канала wch, делают калькуляции и записывают в канал res
// обходим массив задач и каждую задачу записываем в канал wch
// закрываем wch
// читаем канал res

func worker(wch <-chan data, res chan<- data) {
	for {
		w, ok := <-wch
		if !ok {
			return
		}

		w *= 2

		res <- w

	}
}

func FanInFanOutExample1() {
	work := []data{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	const numWorkers = 3

	wch := make(chan data, len(work))
	res := make(chan data, len(work))

	for i := 0; i < numWorkers; i++ {
		go worker(wch, res)
	}

	for _, w := range work {
		wch <- w
	}

	close(wch)

	for range work {
		w := <-res
		fmt.Println(w)
	}

}
