package concurrency

type Todo struct {
	UserID    int
	ID        int
	Title     string
	Completed bool
}

type Post struct {
	UserID int
	ID     int
	Title  string
	Body   string
}

const baseTodoURL = "https://jsonplaceholder.typicode.com/todos/"
const basePostURL = "https://jsonplaceholder.typicode.com/posts/"

func makeList(size int) []int {

	res := make([]int, 0, size)
	for i := 0; i < size; i++ {
		res = append(res, i)
	}

	return res
}
