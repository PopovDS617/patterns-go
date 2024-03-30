package projects

import (
	"io"
	"net/http"
	"sync"
)

func HTTPGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)

}

// Func is the type of the function to memoize.
type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// !+
type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {

		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)

		close(e.ready) // broadcast ready condition
	} else {

		memo.mu.Unlock()

		<-e.ready // wait for ready condition
	}
	return e.res.value, e.res.err
}

// var 2

// package memo

// // Func is the type of the function to memoize.
// type Func func(key string) (interface{}, error)

// // A result is the result of calling a Func.
// type result struct {
// 	value interface{}
// 	err   error
// }

// type entry struct {
// 	res   result
// 	ready chan struct{} // closed when res is ready
// }

// // A request is a message requesting that the Func be applied to key.
// type request struct {
// 	key      string
// 	response chan<- result // the client wants a single result
// }

// type Memo struct{ requests chan request }

// // New returns a memoization of f.  Clients must subsequently call Close.
// func New(f Func) *Memo {
// 	memo := &Memo{requests: make(chan request)}
// 	go memo.server(f)
// 	return memo
// }

// func (memo *Memo) Get(key string) (interface{}, error) {
// 	response := make(chan result)
// 	memo.requests <- request{key, response}
// 	res := <-response
// 	return res.value, res.err
// }

// func (memo *Memo) Close() { close(memo.requests) }

// func (memo *Memo) server(f Func) {
// 	cache := make(map[string]*entry)
// 	for req := range memo.requests {
// 		e := cache[req.key]
// 		if e == nil {
// 			// This is the first request for this key.
// 			e = &entry{ready: make(chan struct{})}
// 			cache[req.key] = e
// 			go e.call(f, req.key) // call f(key)
// 		}
// 		go e.deliver(req.response)
// 	}
// }

// func (e *entry) call(f Func, key string) {
// 	// Evaluate the function.
// 	e.res.value, e.res.err = f(key)
// 	// Broadcast the ready condition.
// 	close(e.ready)
// }

// func (e *entry) deliver(response chan<- result) {
// 	// Wait for the ready condition.
// 	<-e.ready
// 	// Send the result to the client.
// 	response <- e.res
// }
