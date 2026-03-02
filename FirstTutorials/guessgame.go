package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var remainingAttempts int = 5

	var successfullyMessage string = "Tebrikler sayıyı buldunuz"

	randomNumber := rand.Intn(100) + 1

	fmt.Println(randomNumber)

	for remainingAttempts > 0 {
		fmt.Print("Sayıyı tahmin edin:")
		var guess int
		fmt.Scanln(&guess)
		remainingAttempts--
		if randomNumber == guess {
			attemptUsed := 5 - remainingAttempts
			fmt.Println(successfullyMessage, guess)
			fmt.Printf("%d. denemede bildiniz!\n", attemptUsed)

			break
		} else {
			if guess < randomNumber {
				fmt.Println("Daha yukarıda tahminde bulunun.")
			} else {
				fmt.Println("Daha aşağıda tahminde bulunun.")
			}
		}
	}
	if remainingAttempts == 0 {
		fmt.Println("Tahmin hakkınız kalmadı.")
		fmt.Printf("Doğru sayı: %d\n", randomNumber)
		for i := 3; i > 0; i-- {
			fmt.Println(" sistem kapatılıyor..", i)
		}
	}

}
