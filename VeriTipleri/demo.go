package main

import "fmt"

func main() {
	fmt.Println("Kafe Sipariş Uygulamasına Hoş geldiniz!")

	var customerName string = "Rıdvan"
	var orderName string = "Americano"
	var quantity int = 2
	var price float64 = 129.99

	var isMember bool = true

	fmt.Println("--------------------------------")
	fmt.Printf("Müşteri Adı: %s\n", customerName)
	fmt.Printf("Sipariş: %s\n", orderName)
	fmt.Printf("Adet: %d\n", quantity)
	fmt.Printf("Üye mi? %t\n", isMember)
	fmt.Println("--------------------------------")

	totalPrice := float64(quantity) * price

	var toplamFiyatString string

	if isMember {
		toplamFiyatString := fmt.Sprintf("Toplam Fiyat (Üye İndirimi): %.2f TL", totalPrice*0.9) // Üyelere %10 indirim
		fmt.Println(toplamFiyatString)
	}

	toplamFiyatString = fmt.Sprintf("Toplam Fiyat: %.2f TL", totalPrice)

	fmt.Println(toplamFiyatString)

}
