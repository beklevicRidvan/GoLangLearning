package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserGetResponse struct {
	Data []User `json:"data"`
	code int    `json:"code"`
}

func main() {

	fmt.Println("REST API Başlıyor.....")
	fmt.Println("GET: http://localhost:8080/users")
	fmt.Println("POST: http://localhost:8080/users")

	http.HandleFunc("/users", UsersHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Sunucu ayağa kalkarken hata: ", err)
		return
	}

}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w)
	case http.MethodPost:
		createUser(w, r)
	case http.MethodDelete:
		deleteUser(w, r)
	case http.MethodPut:
		updateUser(w, r)
	case http.MethodPatch:
		updateUserSet(w, r)

	default:
		http.Error(w, "Yanlış yöntem denediniz", http.StatusMethodNotAllowed)
	}
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

		// Dosyayı oku

		data, err := os.ReadFile(fileName)

		if err != nil {
			http.Error(w, "Dosya okunamadı", http.StatusInternalServerError)
			return
		}
		users := []User{}
		json.Unmarshal(data, &users)

		userGetResponse := UserGetResponse{code: http.StatusOK, Data: users}

		json.NewEncoder(w).Encode(userGetResponse)
	}

}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Response JSON olacak

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Response JSON olacak

}
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Response JSON olacak

}
func updateUserSet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Response JSON olacak

}
