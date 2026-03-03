package main

import "fmt"

func main() {
	users := map[string]int{
		"ridvan": 25,
		"ali":    30,
		"ayse":   28,
	}
	fmt.Println("Map ", users)

	var userName string

	fmt.Print("Kullanıcı adı giriniz:")
	fmt.Scanln(&userName)

	deleteMap(users, userName)
	fmt.Println("Map ", users)
}

func hasContainsInMap(users map[string]int, userName string) bool {
	value, ok := users[userName]

	if ok {
		fmt.Println("Username bulundu: ", value)
		return true
	} else {
		fmt.Println("Username bulunamadı: ", value)
		return false
	}
}

func addMap(users map[string]int, userName string) bool {
	if age, ok := users[userName]; ok {
		users[userName] = age + 1
		return true
	} else {
		users[userName] = 18
		return false
	}
}

func addIleriMap(users map[string]int, userName string) bool {
	if _, ok := users[userName]; ok {
		users[userName]++
		return true
	}
	users[userName] = 18
	return false
}
func deleteMap(users map[string]int, userName string) {
	_, ok := users[userName]
	if ok {
		delete(users, userName)
		fmt.Printf("%s adlı kullanıcı başarıyla silindi\n", userName)
	} else {
		fmt.Println("Kullanıcı bulunamadı")
	}
}

func counterPattern(word string) {
	counts := map[rune]int{}

	for _, ch := range word {
		counts[ch]++
	}

	fmt.Println("input:", word)
	fmt.Println("output:")
	for letter, count := range counts {
		fmt.Printf("%c: %d\n", letter, count)
	}
}
