package main

import "fmt"

func main() {
	var input int
	fmt.Print("Ketikkan nilai input: ")
	fmt.Scanln(&input)

	if input%7 == 0 {
		fmt.Println(input, "adalah bilangan kelipatan 7")
	} else {
		fmt.Println(input, "bukan bilangan kelipatan 7")
	}
}
