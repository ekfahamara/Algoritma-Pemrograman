package main

import "fmt"

func hitung_luas_lingkaran(r int) float64 {
	var luas float64
	luas = 22 / 7 * float64(r*r)
	return luas
}

func main() {
	var r int
	var hasil float64

	fmt.Scan(&r)
	hasil = hitung_luas_lingkaran(r)
	fmt.Println(hasil)
}
