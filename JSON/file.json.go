package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	user := User{
		Name:  "Sercan",
		Age:   31,
		Email: "sercan@example.com",
	}
	user1 := User{
		Name:  "Rıdvan",
		Age:   31,
		Email: "ridvan@example.com",
	}
	user2 := User{
		Name:  "Mehmet",
		Age:   31,
		Email: "mehmet@example.com",
	}
	user3 := User{
		Name:  "Ali",
		Age:   31,
		Email: "ali@example.com",
	}
	user4 := User{
		Name:  "Ayşe",
		Age:   31,
		Email: "ayse@example.com",
	}
	userSlice := []User{
		user,
		user1,
		user2,
		user3,
		user4,
	}

	data, err := json.MarshalIndent(userSlice, "", "  ")
	if err != nil {
		fmt.Println("Json oluşturma hatası:", err)
		return
	}
	err = os.WriteFile("user.json", data, 0644)
	if err != nil {
		fmt.Println("Dosya yazma hatası:", err)
		return
	}
	fmt.Println("Json dosyaya yazıldı")

	/*
		Kod 		|	 Anlamı
		4			|	sadece okuma
		2			| 	sadece yazma
		1			| 	sadece çalıştırma
		5 (4 + 1)	|   Okuma ve çalıştırma
		6 (4 + 2)   | 	okuma + yazma
		7 (4 + 2 + 1)| okuma + yazma + çalıştırma

		3 rakamın anlamı nedir
		1- Owner / Sahibi
		2- Group - aynı gruptaki diğer kullanıcılar
		3- Diğleri - Sistemdeki herkes

		644 => Sahibi Okur ve yazar  | Gruptakiler sadece okur | Sistemdeke diğerleri sadece okur
		600 => Sahibi okur ve yazar  | Gruptakiler ve Sistemdekiler erişemez
		755 => Sahibi Okur yazar çalıştırır | Gruptakiler okur ve çalıştırır | Sistemdeki diğerleri okur ve çalıştırır
		666 => Herkes okur ve yazar
	*/

	fileData, err := os.ReadFile("user.json")
	if err != nil {
		fmt.Println("Dosya okuma hatası:", err)
		return
	}
	var readUser []User
	err = json.Unmarshal(fileData, &readUser)
	if err != nil {
		fmt.Println("Json çözümleme hatası: ", err)
		return
	}
	fmt.Println("Dosyadan Okunan Kullanıcı")
	fmt.Println("İsim:", readUser[0].Name)
	fmt.Println("Yaş:", readUser[0].Age)
	fmt.Println("E-posta:", readUser[0].Email)
}
