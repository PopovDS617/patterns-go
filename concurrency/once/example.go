package once

import (
	"fmt"
	"log"
	"sync"
)

type fileBuffer struct {
	data []byte
	once sync.Once
}

func (f *fileBuffer) getFile() []byte {
	var err error

	if f == nil {
		log.Fatalln("receiver must not be nil")
	}

	f.once.Do(func() {
		fmt.Println("loading file")

		f.data = []byte("hello world")

		if err != nil {
			log.Fatalln(err)

		}
	})
	return f.data

}

func OnceExample() {
	b := fileBuffer{}

	var wg sync.WaitGroup

	for i := 1; i < 15; i++ {
		wg.Add(1)

		go func() {
			data := b.getFile()
			if data == nil {
				log.Println("error: file not loaded")
			}
			fmt.Println("got data", data)
			wg.Done()
		}()

	}

	wg.Wait()

}
