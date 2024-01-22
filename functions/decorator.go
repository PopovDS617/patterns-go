package functions

import "fmt"

type LogStore struct {
	data []int
}

func (s LogStore) log() {
	fmt.Println(s.data)
}

func (s *LogStore) saveLog(updData int) {
	s.data = append(s.data, updData)
}

func multiply(x, y int) int {
	return x * y
}

func multiplyAndSaveLog(logger *LogStore) func(int, int) int {

	return func(x, y int) int {

		res := multiply(x, y)

		logger.saveLog(res)

		return res
	}
}

func Calc() {

	logger := LogStore{}

	decoratedFunc := multiplyAndSaveLog(&logger)

	decoratedFunc(5, 400)
	decoratedFunc(800, 1200)

	logger.log()

}
