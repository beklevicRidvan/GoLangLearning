package main

import "fmt"

func main() {
	users := map[string][]string{
		"ridvan": {"Go çalış", "Backend oku"},
		"ali":    {"Spor yap"},
	}

	fmt.Println("MİNİ BLOG YÖNETİM SİSTEMİNE HOŞELDİN")
	fmt.Println("--------------------------------")

	for {
		printMenu()

		var choice string
		fmt.Print("Seçiminiz:")
		fmt.Scanln(&choice)

		if choice == "q" {
			return
		}

		switch choice {
		case "1":
			value := addUser(users)
			if !value {
				continue
			}

		case "2":
			value := addNote(users)
			if !value {
				continue
			}

		case "3":
			value := showNotes(users)
			if !value {
				continue
			}

		case "4":
			value := deleteNote(users)
			if !value {
				continue
			}

		case "5":
			value := deleteUser(users)
			if !value {
				continue
			}
			break
		case "6":
			showAllUsers(users)

		default:
			defaultFunction()

		}
	}

}

func printMenu() {
	var menu string = "1.Kullanıcı Oluştur\n2.Not Ekle\n3.Notları Listele\n4.Not sil\n5.Kullanıcı sil\n6.Tüm Kullacıları Listele\nÇıkış için 'q' ya basın."
	fmt.Println(menu)
}

func addUser(users map[string][]string) bool {
	userName := inputUserName()

	if _, exists := users[userName]; exists {
		fmt.Println("Girdiğiniz kullanıcı adında zaten bir kullanıcı var!")
		return false
	} else {
		users[userName] = []string{}
		fmt.Println("Kullanıcı başarıyla oluşturuldu.")
		return true
	}
}

func addNote(users map[string][]string) bool {
	userName := inputUserName()
	var newNote string
	fmt.Print("Eklemek istediğiniz notu girin:")
	fmt.Scanln(&newNote)

	if _, exists := users[userName]; exists {
		users[userName] = append(users[userName], newNote)
		fmt.Println("Not başarıyla eklendi")
		return true
	} else {
		fmt.Println("Kullanıcı bulunamadı")
		return false
	}
}

func showNotes(users map[string][]string) bool {
	userName := inputUserName()
	if slice, exists := users[userName]; exists {
		if len(slice) > 0 {
			for i, note := range slice {
				fmt.Printf("%d) %s\n", i, note)
			}
		} else {
			fmt.Println("Bu Kullanıcının notları boş")
		}
		return true
	} else {
		fmt.Println("Kullanıcı bulunamadı")
		return false
	}
}

func inputUserName() string {
	var userName string
	fmt.Print("Kullanıcı ismini girin:")
	fmt.Scanln(&userName)
	return userName
}

func inputIndex() int {
	var removeAtIndex int
	fmt.Print("Silmek istediğiniz notun indexini girin:")
	fmt.Scanln(&removeAtIndex)
	return removeAtIndex
}

func deleteNote(users map[string][]string) bool {
	var userName string = inputUserName()
	if slice, exists := users[userName]; exists {
		fmt.Printf("%s adlı kişiye ait notlar: %v\n", userName, slice)
		if len(slice) > 0 {
			removeAtIndex := inputIndex()
			if removeAtIndex < 0 || removeAtIndex >= len(slice) {
				fmt.Println("Geçersiz index.")
				return false
			}
			users[userName] = append(slice[:removeAtIndex], slice[removeAtIndex+1:]...)
			fmt.Printf("%s adlı kullanıcının %d indexindeki not başarıyla kaldırıldı\n", userName, removeAtIndex)
			return true
		} else {
			fmt.Println("Kullanıcının notları boş sistemde silenecek not bulunmuyor!!")
			return false
		}
	} else {
		fmt.Println("Kullanıcı bulunamadı")
		return false
	}
}

func deleteUser(users map[string][]string) bool {

	userName := inputUserName()
	if _, exists := users[userName]; exists {
		delete(users, userName)
		fmt.Printf("%s isimli kullanıcı başarıyla  kaldırıldı.\n", userName)
		return true
	} else {
		fmt.Println("Sistemde zaten böyle bir kullanıcı bulunmuyor..")
		return false
	}
}

func defaultFunction() {
	fmt.Println("Lütfen menüdeki değerlerden bir değer giriniz!!")
}

func showAllUsers(users map[string][]string) {
	fmt.Println("Tüm Kullanıcılar")
	for k, v := range users {
		fmt.Println(k, v)
	}
}
