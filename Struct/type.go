package main

import "fmt"

type Age int
type Email string

func main() {
	var yas Age = 31
	fmt.Println(yas)
	var eposta Email = "rdvn.beklevic@gmail.com"
	fmt.Println(eposta)

	var sayi int = 31
	yas = Age(sayi)
	fmt.Println(yas)
}
