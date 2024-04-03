package interfaces

import "fmt"

type MyInt interface{}

func UseTypeAssertion() {

	var unknown MyInt = "hello"
	s, ok := unknown.(string)

	fmt.Println(s, ok)

}
