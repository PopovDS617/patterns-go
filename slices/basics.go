package slicespkg

import "fmt"

func Slice1() {
	a1 := make([]int, 0, 10) // len 0, cap 10

	a1 = append(a1, []int{1, 2, 3, 4, 5}...) // len 5, cap 10

	a2 := append(a1, 6) // len 6, cap 10

	a3 := append(a1, 7) // len 6, cap 10

	a4 := append(a1, 8) // len 6, cap 10

	fmt.Println(a1, a2, a3, a4) // a1=[1 2 3 4 5] a2=[1 2 3 4 5 8] a3=[1 2 3 4 5 8] a4=[1 2 3 4 5 8]

}
