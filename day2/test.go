package main

import "fmt"

// func main(){
// 	var usia int = 22

// 	fmt.Println(usia)
// }

// func main() {
// 	var fruits = []string{"apple", "grape", "banana", "melon"}
// 	var newFruits = fruits[0:2]

// 	fmt.Println(newFruits)
// }
type Point struct {
	X int
	Y int
}

func main() {
	var p Point
	p.X = 1
	p.Y = 2

	var q *Point
	q.X = 3
	q.Y = 4
	fmt.Println(p, q)
}
