package concurrency

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func FetchLineByLine() {

	list := makeList(200)

	start := time.Now()

	for _, v := range list {

		RegularFetch(baseTodoURL + fmt.Sprintf("%v", v))
	}

	took := time.Since(start).Seconds()

	fmt.Printf("operation took %.2f seconds\n", took)
}

func PromiseAllLikeFunctionWithWG() {

	var wg sync.WaitGroup

	list := makeList(200)

	wg.Add(len(list))

	start := time.Now()

	for _, v := range list {

		go WGGoroutineFetch(baseTodoURL+fmt.Sprintf("%v", v), &wg)
	}

	wg.Wait()

	took := time.Since(start).Seconds()

	fmt.Printf("operation took %.2f seconds\n", took)
}
func PromiseAllLikeFunctionWithUnbuffCh() {

	list := makeList(200)

	start := time.Now()

	for _, v := range list {

		ch := make(chan string)

		go ChGoroutineFetch(baseTodoURL+fmt.Sprintf("%v", v), ch)
		fmt.Println(<-ch)
	}

	took := time.Since(start).Seconds()

	fmt.Printf("operation took %.2f seconds\n", took)
}
func PromiseAllLikeFunctionWithBuffCh() {

	list := makeList(200)

	start := time.Now()

	ch := make(chan string, len(list))

	for _, v := range list {

		go ChGoroutineFetch(baseTodoURL+fmt.Sprintf("%v", v), ch)

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
	defer close(ch)

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
