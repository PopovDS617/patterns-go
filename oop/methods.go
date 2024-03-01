package oop

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

type Path []Point

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (path Path) Distance() float64 {

	sum := 0.0

	for i := range path {

		if i > 0 {
			sum += path[i-1].Distance(path[i])

		}

	}
	return sum
}

var PerimOne Path = Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
var PerimTwo Path = Path{Point{x: 15, y: 24}, Point{x: 42, y: 125}}

type Books map[string][]string

func (b Books) Size() int {
	return len(b)
}

func (b Books) Get(key string) string {
	if vs := b[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

func (b Books) Add(key string, value string) {
	b[key] = append(b[key], value)
}

func BookMethods() {

	var BookMap = make(Books)
	BookMap.Add("horrors", "Drakula")
	BookMap.Add("horrors", "It")
	BookMap.Add("plays", "Romeo and Juliet")

	fmt.Println(BookMap.Get("horrors"))
	fmt.Println(BookMap.Get("dramas"))
	fmt.Println(BookMap.Get("plays"))
	fmt.Println(BookMap["horrors"])
	fmt.Println(BookMap.Size())

}
