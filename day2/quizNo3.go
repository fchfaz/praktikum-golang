package main

import "fmt"

func main() {
	var alas1, alas2, tinggi float64

	fmt.Print("Masukkan panjang dari alas pertama: ")
	fmt.Scanln(&alas1)

	fmt.Print("Masukkan panjang dari alas kedua: ")
	fmt.Scanln(&alas2)

	fmt.Print("Masukkan tinggi trapesium: ")
	fmt.Scanln(&tinggi)

	luas := ((alas1 + alas2) / 2) * tinggi

	fmt.Printf("Luas trapesium adalah: %.2f", luas)
}
