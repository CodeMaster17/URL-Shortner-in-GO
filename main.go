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
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

// keys of string type
// values of URL type
var UrlDB = make(map[string]URL)

func generateShortURL(OriginalURL string) string {

	// converts the given string to particular Hash value
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL)) // converts given URL into bytes
	data := hasher.Sum(nil)           // returns the hash value of the given URL
	hash := hex.EncodeToString(data)  // converts the hash value into string
	return hash[:8]                   // returns the first 8 characters of the hash value
}

func creatURL(OriginalURL string) string {
	shortURL := generateShortURL(OriginalURL)
	id := shortURL // id to be stored in db
	UrlDB[id] = URL{
		ID:           id,
		OriginalURL:  OriginalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}

func getURL(id string) (URL, error) {
	url, ok := UrlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// Short url handler
func shortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortURL_ := creatURL(data.URL)
	// fmt.Fprintf(w, "Short URL: %s", shortURL)
	response := struct {
		ShortURL string `json:"short_url"`
	}{ShortURL: shortURL_}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func main() {
	// Handler for the root URL

	fmt.Println("Starting the server at port 8080")
	http.HandleFunc("/", handler)
	http.HandleFunc("/shorten", shortURLHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server")
	}
}
