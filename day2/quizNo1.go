package main

import (
	"fmt"
)

func bilPrima(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	for i := 0; i <= 100; i++ {
		if bilPrima(i) {
			fmt.Printf("%d ", i)
		}
	}
}
