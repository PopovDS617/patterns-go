package faninfanout

import (
	"fmt"
	"math/rand"
)

type Application struct {
	name    string
	content string
}

func mockScan() string {
	if rand.Intn(100) > 90 {
		return "ALERT - vulterability found"
	}
	return "OK - all fine"
}

func scanSQLInjection(data Application, res chan<- string) {
	res <- fmt.Sprintf("SQL injection scan: %s scanned, result: %s", data.name, mockScan())
}
func scanTimingExploits(data Application, res chan<- string) {
	res <- fmt.Sprintf("Timing exploits scan: %s scanned, result: %s", data.name, mockScan())
}
func scanAuth(data Application, res chan<- string) {
	res <- fmt.Sprintf("Authentication scan: %s scanned, result: %s", data.name, mockScan())
}

func FanInFanOutExample2() {

	si := []Application{
		{name: "ms_comments", content: "package comments"},
		{name: "ms_posts", content: "package posts"},
		{name: "ms_hints", content: "package hints"},
		{name: "ms_video", content: "package video"},
		{name: "ms_sharing", content: "package sharing"},
		{name: "ms_bookmarks", content: "package bookmarks"},
	}

	res := make(chan string, len(si)*3)

	// на каждый элемент массива запускается 3 разные горутины
	for _, d := range si {
		d := d

		go scanSQLInjection(d, res)
		go scanTimingExploits(d, res)
		go scanAuth(d, res)
	}

	for i := 0; i < cap(res); i++ {
		fmt.Println(<-res)
	}

	fmt.Println("[OK] scanning is done")
}
