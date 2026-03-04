package main

import (
	"errors"
	"fmt"
)

func main() {

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Hata oluştu: ", err)
	} else {
		fmt.Println("Sonuç: ", result)
	}

	result2, err1 := GetUser(5)
	if err1 != nil {
		fmt.Println("Hata oluştu: ", err1)
	} else {
		fmt.Println("Sonuç: ", result2)
	}

}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("0'a bölme hatası")
	}
	return a / b, nil
}

func GetUser(id int) (string, error) {
	users := map[int]string{
		1: "Sercan",
		2: "Rıdvan",
		3: "Mehmet",
		4: "İbrahim",
	}

	if name, exists := users[id]; exists {
		return name, nil
	} else {
		return "", errors.New("Kullanıcı bulunamadı")
	}
}
