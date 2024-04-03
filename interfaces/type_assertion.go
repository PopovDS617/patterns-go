package interfaces

import "fmt"

type MyInt interface{}

func UseTypeAssertion() {

	var unknown MyInt = "hello"
	s, ok := unknown.(string)

	fmt.Println(s, ok)

}

func UseTypeSwitch() {
	var unknown MyInt = "hello"

	switch v := unknown.(type) {
	case int:
		fmt.Printf("%v is a int\n", v)
	case string:
		fmt.Printf("%v is a string\n", v)
	}

}
