package main

import "fmt"

func main() {

	var remainingAttempts int = 3
	var welcomeMessage string = "ATM Uygulamasına Hoşgeldiniz!"
	var inputPinMessage string = "Lütfen 4 haneli PIN kodunuzu giriniz: "
	var inputAmountMessage string = "Lütfen çekmek istediğiniz tutarı giriniz: "
	var atmMenu string = "Lütfen yapmak istediğiniz işlemi seçiniz:\n1. Para Çekme\n2.Para Yatırma\n3. Bakiye Sorgulama\n4. Çıkış"
	var exitMessage string = "İyi günler dileriz..."
	var defaultMessage string = "Geçersiz seçim. Lütfen tekrar deneyiniz."
	authenticated := false

	var balance = 1000.00
	fmt.Println(welcomeMessage)
	fmt.Println("-----------------------------")
	fmt.Print(inputPinMessage)
	pinCode := "1234"

	for remainingAttempts > 0 {
		var pin string
		fmt.Scanln(&pin)
		remainingAttempts--
		if pin != pinCode {
			if remainingAttempts > 0 {
				fmt.Println("Hatalı PIN kodu. Lütfen tekrar deneyiniz.")
				fmt.Printf("Kalan deneme hakkınız: %d\n", remainingAttempts)
			}
			continue
		}
		authenticated = true
		fmt.Println("PIN kodu doğrulandı. İşlemlere yönlendiriliyorsunuz...")
		break

	}

	if !authenticated {
		fmt.Println("Hesabınız bloke olmuştur. Lütfen bankanızla iletişime geçiniz.")
		return
	}

	for {

		fmt.Println(atmMenu)

		var choice int
		fmt.Scanln(&choice)

		if choice == 4 {
			for i := 3; i > 0; i-- {
				fmt.Printf("Çıkış yapılıyor... %d saniye\n", i)
			}
			fmt.Println(exitMessage)
			return
		}
		switch choice {
		case 1:
			fmt.Print(inputAmountMessage)
			var amount float64
			fmt.Scanln(&amount)
			if amount <= 0 {
				fmt.Println("Lütfen geçerli bir tutar giriniz.")
				continue
			} else if amount > balance {
				fmt.Println("Yetersiz bakiye.")
				continue
			} else {
				balance -= amount
				fmt.Printf("Lütfen paranızı alınız. Kalan bakiye: %.2f\n", balance)
			}
		case 2:
			fmt.Print("Lütfen yatırmak istediğiniz tutarı giriniz: ")
			var amount float64
			fmt.Scanln(&amount)
			if amount <= 0 {
				fmt.Println("Lütfen geçerli bir tutar giriniz.")
				continue
			}
			balance += amount
			fmt.Printf("Para yatırıldı. Güncel bakiyeniz: %.2f\n", balance)

		case 3:
			fmt.Printf("Mevcut bakiyeniz: %.2f\n", balance)
		default:
			fmt.Println(defaultMessage)

		}

	}

}
