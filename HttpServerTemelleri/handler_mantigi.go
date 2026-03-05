package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Go HTTP Server Başlıyor...")
	fmt.Println("Tarayıcıya yaz: http://localhost:8081")

	// Route tanımlama

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/json", JsonHandler)

	// Server Başlatma
	//ListenAndServe
	// Önce port ardından handler/router yazılır default geçtiğimiz için nil verdik

	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/ridvan", RidvanHandler)

	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		fmt.Println("Server Hata verdi : ", err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Anasayfaya Hoş Geldiniz!")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Merhaba Dünya !!")
}
func JsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsonData := `{"message":"Go HTTP Server temelleri dersinden selamlar"}`

	fmt.Fprintln(w, jsonData)
}
func RidvanHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Rıdvan Merhaba!!")
}
