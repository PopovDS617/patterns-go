package pointers

import "fmt"

type Person struct {
	Name string
}

func ChangeName(person *Person) {
	person = &Person{
		Name: "Alice",
	}

	fmt.Println(person) // Alice => 0xc00018c028
}

func Pointer1() {

	person := &Person{Name: "Bob"}

	fmt.Println(person) // Bob => 0xc00018c018

	ChangeName(person)

	fmt.Println(person) // Bob => 0xc00018c018

}
