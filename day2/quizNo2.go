package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%7 == 0 {
			fmt.Printf("%d adalah bilangan kelipatan 7\n", i)
		}
	}
}
