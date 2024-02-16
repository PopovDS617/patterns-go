package concurrency

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

func FetchTodo(url string, ch chan string, count *atomic.Int32) {

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

	ch <- fmt.Sprintf("200ok - TODO - request took %.2f sec, data: %+v\n", took, todo)

	count.Add(1)

}

func FetchPost(url string, ch chan string, count *atomic.Int32) {

	start := time.Now()

	var post Post

	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&post); err != nil {
		fmt.Println(err)
		return
	}

	took := time.Since(start).Seconds()

	ch <- fmt.Sprintf("200ok - POST - request took %.2f sec, data: %+v\n", took, post)

	count.Add(1)

}

func GetTodos(ch chan string, count *atomic.Int32) {
	list := makeList(100)

	for i := 0; i < len(list); i++ {

		time.Sleep(time.Second * 1)

		go FetchTodo(baseTodoURL+fmt.Sprintf("%v", list[i]), ch, count)

	}
}
func GetPosts(ch chan string, count *atomic.Int32) {
	list := makeList(200)

	for i := 0; i < len(list); i++ {
		time.Sleep(time.Second * 2)

		go FetchPost(basePostURL+fmt.Sprintf("%v", list[i]), ch, count)

	}
}

func GetAll() {

	var count atomic.Int32

	todoCh := make(chan string, 200)
	postCh := make(chan string, 100)

	// optional - break the loop after 100
	qCh := make(chan bool, 1)

	go func() {
		for count.Load() < 10 {
		}
		qCh <- true
	}()

	go GetTodos(todoCh, &count)
	go GetPosts(postCh, &count)

outer:
	for {
		select {
		case todo := <-todoCh:
			fmt.Println(todo)
		case post := <-postCh:
			fmt.Println(post)
		case <-qCh:
			break outer

		}

	}
}
