package main

import "fmt"

func main() {
	fmt.Println("Fastfood Gün sonu Uygulaması")
	fmt.Println("----------------------------")

	burgerPrice := 220.40
	friesPrice := 120.0
	drinkPrice := 45.90

	var burgerCount, friesCount, drinkCount int

	fmt.Print("Satılan Burger adedi:")
	fmt.Scanln(&burgerCount)
	fmt.Print("Satılan French Fries adedi:")
	fmt.Scanln(&friesCount)
	fmt.Print("Satılan Drink adedi:")
	fmt.Scanln(&drinkCount)

	total := (float64(burgerCount) * burgerPrice) + (float64(friesCount) * friesPrice) + (float64(drinkCount) * drinkPrice)

	fmt.Printf("Toplam Satış: %.2f TL\n", total)

	if total >= 1000 {
		fmt.Println("Gün sonu hedefi aşıldı!")
	} else if total >= 500 {
		fmt.Println("Gün sonu hedefi yarım kaldı.")
	} else {
		fmt.Println("Gün sonu hedefi tutturulamadı.")
	}
	fmt.Println("----------------------------")

	fmt.Println("Kapanış işlevi başlıyor: ")

	for i := 3; i > 0; i-- {
		fmt.Printf("%d...\n", i)
	}
	fmt.Println("Mağaza kapandı. İyi akşamlar!")

}
