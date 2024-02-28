package internals

import "fmt"

func CaptLoopVars() {

	ch := make(chan bool)

	for i := 0; i <= 10; i++ {

		// go func() {
		// 	fmt.Println(i)  - неверно, i постоянно меняется и на момент старта горутины будет равна 10
		// 	ch <- true
		// }()
		go func(i int) {
			fmt.Println(i) // верно - значение i каждый раз фиксируется и замыкается в горутине
			ch <- true
		}(i)

	}

	for range ch {
		<-ch
	}

}
