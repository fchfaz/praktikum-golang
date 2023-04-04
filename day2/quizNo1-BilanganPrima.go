package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Print("Ketik sebuah bilangan: ")
	fmt.Scan(&n)
	if isPrime(n) {
		fmt.Printf("%d adalah bilangan prima\n", n)
		fmt.Println("Program selesai. Tekan tombol enter untuk keluar.")
		fmt.Scanln()
	} else {
		fmt.Printf("%d bukan bilangan prima\n", n)
		fmt.Println("Program selesai. Tekan tombol enter untuk keluar.")
		fmt.Scanln()
	}

}
