package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	fmt.Println("REST API Başlıyor.....")
	fmt.Println("GET: http://localhost:8080/users")
	fmt.Println("POST: http://localhost:8080/users")

	// "/users" endpointini UsersHandler fonksiyonuna bağladık
	http.HandleFunc("/users", UsersHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server Hata verdi : ", err)
		return
	}

}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getUsers(w)
		return
	}
	if r.Method == http.MethodPost {
		createUser(w, r)
		return
	}
	if r.Method == http.MethodDelete {
		deleteUser(w, r)
		return
	}
	http.Error(w, "Sadece GET ve POST methodları desteklenmektedir.", http.StatusMethodNotAllowed)

}

func getUsers(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json") // Response JSON olacak

	fileName := "users.json"
	info, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			// Dosya yoksa -> Oluşturma işlemi yap
			fmt.Println(fileName + " dosyası yok, oluşturuluyor...")

			empty := []User{}

			data, marshallErr := json.MarshalIndent(empty, "", " ")
			if marshallErr != nil {
				http.Error(w, "Json oluşturulamadı"+marshallErr.Error(), http.StatusInternalServerError)
				return
			}
			writerErr := os.WriteFile(fileName, data, 0644)
			if writerErr != nil {
				http.Error(w, "Dosya oluşturulamadı"+writerErr.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			// Diğer hatalar -> Dosyayı okumaya izin yoksa vs , Disk
			http.Error(w, "Dosya erişim hatası:"+err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if info != nil {
			fmt.Println("Dosya Adı:", info.Name())
			fmt.Println("Dosya Boyutu:", info.Size())
			fmt.Println("Dosya Son değiştirilme tarihi:", info.ModTime())
			fmt.Println("Dizin mi ?", info.IsDir())
			fmt.Println("Dosyanın izinleri:", info.Mode())

		}

		// dosyayı oku

		data, err := os.ReadFile(fileName)
		if err != nil {
			http.Error(w, "Dosya okunamadı", http.StatusInternalServerError)
			return
		}

		// Dosyadaki JSON'u GO Struct Listesine Çevirme
		var users []User
		json.Unmarshal(data, &users)

		// Kullanıcı listesini JSON olarak döndür

		json.NewEncoder(w).Encode(users)

	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Response JSON olacak

	// Request body ' i oku

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Body okunamadı", http.StatusBadRequest)
		return
	}
	var newUser User
	err = json.Unmarshal(body, &newUser)
	if err != nil {
		http.Error(w, "Geçersiz JSON formatı", http.StatusBadRequest)
		return
	}
	if newUser.Name == "" {
		http.Error(w, "Name alanı zorunlu", http.StatusBadRequest)
		return
	}
	if newUser.Email == "" {
		http.Error(w, "E-posta alanı zorunlu", http.StatusBadRequest)
		return
	}
	if newUser.Age <= 0 {
		http.Error(w, "Yaş alanı zorunlu", http.StatusBadRequest)
		return
	}

	// users.json dosyasını oku; yoksa oluştur
	fileName := "users.json"
	var users []User
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// Dosya yoksa boş liste oluştur
		users = []User{}
		emptyFile, _ := json.MarshalIndent(users, "", " ")
		os.WriteFile(fileName, emptyFile, 0644)
	} else {
		// Dosya varsa içeriğini oku
		fileData, _ := os.ReadFile(fileName)

		// JSON 'u listeye aktar
		json.Unmarshal(fileData, &users)
	}

	// Yeni kullanıcı listeye ekleme
	users = append(users, newUser)

	// Güncellenmiş kullanıcı listesini JSON'a çevir

	fileJson, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		http.Error(w, "JSON oluşturulamadı", http.StatusInternalServerError)
		return
	}
	os.WriteFile(fileName, fileJson, 0644)

	// API Cevabı

	/*
			interface{} -> Her türden veriyi kabul edebilen en genel tiptir. empty interface

			Burada mapin value kısmı yani interface{},
			hem string değerini tutuyor:
				"message":"Kullanıcı kaydedildi"
			hem de User structını tutuyor:
				"user" : newUser

			Böylece farklı türlerdeki verileri tek map içinde tutabilmiş oluyoruz.


		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Kullanıcı kaydedildi",
			"user":    newUser,
		})
	*/
	response := CreateUserResponse{
		Message: "Kullanıcı kaydedildi",
		User:    newUser,
	}
	json.NewEncoder(w).Encode(response)
}
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Body okunamadı", http.StatusBadRequest)
		return
	}
	var deleteUserRequest DeleteUserRequest
	err = json.Unmarshal(body, &deleteUserRequest)
	if err != nil {
		http.Error(w, "Geçersiz JSON formatı", http.StatusBadRequest)
		return
	}
	if deleteUserRequest.Name == "" {
		http.Error(w, "İsim zorunlu", http.StatusBadRequest)
		return
	}

	// users.json dosyasını oku; yoksa oluştur

	fileName := "users.json"

	var users []User
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		users = []User{}
		emptyFile, _ := json.MarshalIndent(users, "", " ")
		os.WriteFile(fileName, emptyFile, 0644)
	} else {
		fileData, _ := os.ReadFile(fileName)
		json.Unmarshal(fileData, &users)
	}

	// Silinecek kullanıcıyı bul
	for i, user := range users {
		if user.Name == deleteUserRequest.Name {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}
	// Güncellenmiş kullanıcı listesini JSON'a çevir
	fileJson, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		http.Error(w, "JSON oluşturulamadı", http.StatusInternalServerError)
		return
	}
	os.WriteFile(fileName, fileJson, 0644)

	response := DeleteUserResponse{
		Message: "Kullanıcı silindi..",
		Name:    deleteUserRequest.Name,
	}
	json.NewEncoder(w).Encode(response)

}

type CreateUserResponse struct {
	Message string `json:"message"`
	User    User   `json:"user"`
}
type DeleteUserResponse struct {
	Message string `json:"message"`
	Name    string `json:"name"`
}

type DeleteUserRequest struct {
	Name string `json:"name"`
}
