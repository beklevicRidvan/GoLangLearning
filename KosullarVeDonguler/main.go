package main

import "fmt"

func main() {
	age := 17

	if age >= 18 {
		fmt.Println("Ehliyet alabilir.")
	} else {
		fmt.Println("Ehliyet alamaz.")
	}

	// burada tanımlananan score değişkeni sadece if bloğu içerisinde geçerlidir.Scoped variable
	if score := 45; score >= 50 {
		fmt.Println("Geçtiniz.")
	} else {
		fmt.Println("Kaldınız.")
	}

	temp := 25
	if temp < 10 {
		fmt.Println("Hava soğuk.")
	} else if temp >= 10 && temp < 20 {
		fmt.Println("Hava ılık.")
	} else {
		fmt.Println("Hava sıcak.")
	}

	// switch-case yapısı

	day := 3
	switch day {
	case 1:
		fmt.Println("Pazartesi")
	case 2:
		fmt.Println("Salı")
	case 3:
		fmt.Println("Çarşamba")
	case 4:
		fmt.Println("Perşembe")
	case 5:
		fmt.Println("Cuma")
	case 6:
		fmt.Println("Cumartesi")
	case 7:
		fmt.Println("Pazar")
	default:
		fmt.Println("Geçersiz gün.")
	}
	grade := 'B'
	switch grade {
	case 'A':
		fmt.Println("Mükemmel")
	case 'B':
		fmt.Println("Çok İyi")
	case 'C':
		fmt.Println("İyi")
	case 'D':
		fmt.Println("Geçer")
	case 'F':
		fmt.Println("Kaldı")
	default:
		fmt.Println("Geçersiz not.")

	}

	for i := 0; i < 5; i++ {
		fmt.Printf("Sayı: %d\n", i)
	}
	fmt.Println("--------------")
	// while döngüsü gibi çalışan for döngüsü
	count := 5
	for count > 0 {
		fmt.Printf("Count: %d\n", count)
		count--
	}

	for {
		fmt.Println("Bu bir sonsuz döngüdür.")
		break //  döngüyü kırmak için break kullanılır
	}

	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Printf("%d çift sayıdır.\n", i)
		} else {
			fmt.Printf("%d tek sayıdır.\n", i)
		}
	}

	for i := 1; i <= 9; i++ {
		if i == 3 {
			continue // 3'ü atla ve döngünün geri kalanını çalıştır
		}
		if i == 5 {
			break // 5'e geldiğinde döngüyü kır
		}
		fmt.Printf("Sayı: %d\n", i)
	}

}
