package main

import (
	"errors"
	"fmt"
)

func main() {
	ridvan := BankAccount{"Ridvan Beklevic", 2500}

	var acc Account = &ridvan

	result, err := acc.withdraw()
	if err != nil {
		fmt.Println("Hata oluştu", err)
	} else {
		fmt.Println("Result: ", result)
	}

	acc.seeBalance()

}

type Account interface {
	deposit() (string, error)
	withdraw() (string, error)
	seeBalance()
}

type BankAccount struct {
	owner   string
	balance float64
}

func (b *BankAccount) deposit() (string, error) {
	inputNumber := getInputNumber()
	if inputNumber < 0 {
		return "", errors.New("Lütfen pozitif bir sayı girin")
	} else {
		b.balance += float64(inputNumber)
		b.seeBalance()
		return "Para yatırma işlevi başarıyla tamamlandı.", nil
	}
}

func (b *BankAccount) withdraw() (string, error) {
	inputNumber := getInputNumber()
	if inputNumber < 0 {
		return "", errors.New("Lütfen pozitif bir sayı girin")
	} else {
		if inputNumber > int(b.balance) {
			return "", fmt.Errorf("... mevcut bakiyeniz: %.2f", b.balance)
		} else {
			b.balance -= float64(inputNumber)
			b.seeBalance()

			return "Para çekme işlemi başarılı", nil
		}
	}
}

func getInputNumber() int {
	var number int
	fmt.Print("Sayı girin:")
	fmt.Scanln(&number)
	return number
}

func (b BankAccount) seeBalance() {
	fmt.Printf("Banka hesabınızda bulunan bakiyeniz: %.2f\n", b.balance)
}
