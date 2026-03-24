package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
)

type UserData struct {
	Name string `json:"name"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/{$}", handleRoot)
	mux.HandleFunc("/goodbye", handleGoodbye)
	mux.HandleFunc("/hello/", handleHelloParameterized)
	mux.HandleFunc("responses/{user}/hello/", handleUserResponsesHello)
	mux.HandleFunc("/user/hello/", handleHelloHeader)
	mux.HandleFunc("/json/", handleJson)

	log.Fatal(http.ListenAndServe(":8080", mux))

}

func handleRoot(w http.ResponseWriter, _ *http.Request) {
	wc, err := w.Write([]byte("Welcome to our Homepage!\n"))
	if err != nil {
		slog.Error("errror writng response", "err", err)
		return
	}

	fmt.Printf("%d bytes writtten\n", wc)
}

func handleGoodbye(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte("Goodbye!"))
	if err != nil {
		slog.Error("error writing response", "err", err)
		return
	}
}

func handleHelloParameterized(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	userList := params["user"]

	username := "User"

	if len(userList) > 0 {
		username = userList[0]
	}

	writeBody(w, username)

}

func handleUserResponsesHello(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("user")

	writeBody(w, username)
}

func handleHelloHeader(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("user")

	if username == "" {
		http.Error(w, "invalid username provided", http.StatusBadRequest)
		return
	}

	writeBody(w, username)

}

func writeBody(w http.ResponseWriter, username string) {
	var output bytes.Buffer
	output.WriteString("Hello, ")
	output.WriteString(username)

	output.WriteString("!\n")

	_, err := w.Write(output.Bytes())
	if err != nil {
		slog.Error("error writing response", "err", err)
		return
	}
}

func handleJson(w http.ResponseWriter, r *http.Request) {
	byteData, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("error reading request body", "err", err)
		http.Error(w, "bad request body", http.StatusBadRequest)
		return
	}

	var reqData UserData

	err = json.Unmarshal(byteData, &reqData)

	if err != nil {
		slog.Error("error unmarshalling request body", "err", err)
		http.Error(w, "error parsing request json", http.StatusBadRequest)
		return
	}

	if reqData.Name == "" {
		slog.Error("invalid Username provided\n", "err", err)

		http.Error(w, "invalid Username provided\n", http.StatusBadRequest)
		return
	}

	writeBody(w, reqData.Name)
}
