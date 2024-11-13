package main

import (
	"crypto/md5"
	"encoding/hex"
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

func main() {

}
