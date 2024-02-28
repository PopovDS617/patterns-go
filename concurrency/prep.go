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

type User struct {
	ID     int
	Active bool
}

func MakeUsersList(v int) []User {
	users := make([]User, 0, v)

	for i := 0; i < v; i++ {

		user := User{
			ID:     i + 1,
			Active: true,
		}

		users = append(users, user)
	}
	return users
}

func (u *User) Deactivate() error {
	u.Active = false
	return nil
}

type ResultWithError struct {
	User User
	Err  error
}
