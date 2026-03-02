package main

import "fmt"

func main() {
	var welcomeMessage string = "Hesap Makinesi Uygulamasına Hoşgeldiniz!"
	var inputFirstNumberMessage string = "Lütfen ilk sayıyı giriniz: "
	var inputSecondNumberMessage string = "Lütfen ikinci sayıyı giriniz: "

	for {
		fmt.Println(welcomeMessage)
		fmt.Println("-----------------------------")
		fmt.Println("Lütfen yapmak istediğiniz işlemi seçiniz:")
		fmt.Println("1. Toplama")
		fmt.Println("2. Çıkarma")
		fmt.Println("3. Çarpma")
		fmt.Println("4. Bölme")
		fmt.Println("5. Çıkış")

		var choice int
		fmt.Scanln(&choice)

		if choice == 5 {
			fmt.Println("Çıkılıyor...")
			return
		}

		// 👇 SAYILARI BURADA AL
		var num1, num2 float64
		fmt.Print(inputFirstNumberMessage)
		fmt.Scanln(&num1)

		fmt.Print(inputSecondNumberMessage)
		fmt.Scanln(&num2)

		switch choice {
		case 1:
			fmt.Printf("Toplam: %.2f\n", num1+num2)
		case 2:
			fmt.Printf("Fark: %.2f\n", num1-num2)
		case 3:
			fmt.Printf("Çarpım: %.2f\n", num1*num2)
		case 4:
			if num2 == 0 {
				fmt.Println("0'a bölünemez.")
				continue
			}
			fmt.Printf("Bölüm: %.2f\n", num1/num2)
		default:
			fmt.Println("Geçersiz seçim.")
		}
	}
}
