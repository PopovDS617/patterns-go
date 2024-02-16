package concurrency

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Todo struct {
	UserID    int
	ID        int
	Title     string
	Completed bool
}

const baseURL = "https://jsonplaceholder.typicode.com/todos/"

func makeList() []int {
	list := make([]int, 0, 25)

	for i := 1; i <= 25; i++ {
		list = append(list, i)
	}
	return list
}

func FetchLineByLine() {

	list := makeList()

	start := time.Now()

	for _, v := range list {

		RegularFetch(baseURL + fmt.Sprintf("%v", v))
	}

	took := time.Since(start).Seconds()

	fmt.Printf("operation took %.2f seconds\n", took)
}

func PromiseAllLikeFunctionWithWG() {

	var wg sync.WaitGroup

	list := makeList()

	wg.Add(len(list))

	start := time.Now()

	for _, v := range list {

		go WGGoroutineFetch(baseURL+fmt.Sprintf("%v", v), &wg)
	}

	wg.Wait()

	took := time.Since(start).Seconds()

	fmt.Printf("operation took %.2f seconds\n", took)
}
func PromiseAllLikeFunctionWithUnbuffCh() {

	list := makeList()

	start := time.Now()

	for _, v := range list {

		ch := make(chan string)

		go ChGoroutineFetch(baseURL+fmt.Sprintf("%v", v), ch)
		fmt.Println(<-ch)
	}

	took := time.Since(start).Seconds()

	fmt.Printf("operation took %.2f seconds\n", took)
}
func PromiseAllLikeFunctionWithBuffCh() {

	list := makeList()

	start := time.Now()

	ch := make(chan string, len(list))

	for _, v := range list {

		go ChGoroutineFetch(baseURL+fmt.Sprintf("%v", v), ch)

	}

	// v 1
	// for elem := range ch {
	// 	fmt.Println(len(ch))
	// 	fmt.Println(elem)
	// }

	// v 2
	// for {
	// 	msg, ok := <-ch

	// 	if !ok {

	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }

	// v 3
	for i := 0; i < len(list); i++ {
		msg, ok := <-ch

		if !ok {
			break
		}
		fmt.Println(msg)
	}

	took := time.Since(start).Seconds()

	fmt.Printf("operation took %.2f seconds\n", took)
}

func RegularFetch(url string) {
	start := time.Now()

	var todo Todo

	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&todo); err != nil {
		fmt.Println(err)
		return
	}

	took := time.Since(start).Seconds()

	fmt.Printf("request took %.2f sec, data: %+v\n", took, todo)

}

func WGGoroutineFetch(url string, wg *sync.WaitGroup) {

	defer wg.Done()

	start := time.Now()

	var todo Todo

	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&todo); err != nil {
		fmt.Println(err)
		return
	}

	took := time.Since(start).Seconds()

	fmt.Printf("request took %.2f sec, data: %+v\n", took, todo)

}

func ChGoroutineFetch(url string, ch chan string) {

	start := time.Now()

	var todo Todo

	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&todo); err != nil {
		fmt.Println(err)
		return
	}

	took := time.Since(start).Seconds()

	ch <- fmt.Sprintf("request took %.2f sec, data: %+v\n", took, todo)

}
