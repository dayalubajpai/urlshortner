package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID          string    `json:"id"`
	OriginalURL string    `json:"original_url"`
	ShortURL    string    `json:"shorturl"`
	CreatedAT   time.Time `json:"created_date"`
}

var urlDB = make(map[string]URL)

func generateShortURL(originalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(originalURL)) //[]byte(originalURL) its convert url to byte
	data := hasher.Sum(nil)
	hexString := hex.EncodeToString(data)
	return hexString[0:8]
}

func createURL(originalURL string) string {
	shortURL := generateShortURL(originalURL)
	Id := shortURL
	urlDB[Id] = URL{
		ID:          Id,
		OriginalURL: originalURL,
		ShortURL:    shortURL,
		CreatedAT:   time.Now(),
	}

	return shortURL
}

func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}

	return url, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func shortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, "Invalid Parameters", http.StatusBadRequest)
		return
	}

	url := createURL(data.URL)

	response := struct {
		ShortURL string `json:"short_url"`
	}{ShortURL: url}

	w.Header().Set("Content-Type:", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectURL(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]

	url, err := getURL(id)

	if err != nil {
		http.Error(w, "something went wrong during url finding", http.StatusBadRequest)
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func main() {
	fmt.Println("Server is starting...")
	http.HandleFunc("/shortner", shortURLHandler)
	http.HandleFunc("/redirect/", redirectURL)
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println("Server not responding", err)
	}
}
