package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//Book type
type Book struct {
	Title     string `json:"title"`
	Subtitle  string `json:"subtitle"`
	Author    string `json:"author"`
	Published string `json:"published"`
}

func main() {
	file, err := os.Open("./books.json")
	if err != nil {
		log.Fatalf("an error occured openinng the file")
	}
	defer file.Close()
	var books []Book
	err = json.NewDecoder(file).Decode(&books)
	if err != nil {
		log.Fatalf("couldn't decode file into json")
	}
	for _, book := range books {
		fmt.Println(book.Title, " ", book.Author)
	}
	bookNewfile, err := os.Create("booksnew.json")
	if err != nil {
		log.Fatalf("couldnt create file")
	}
	defer bookNewfile.Close()
	json.NewEncoder(bookNewfile).Encode(books)
}
