package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var baseURL = "http://localhost:3000"

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	fmt.Println("========")
	fmt.Println("Fetch Books")
	fmt.Println("========")
	books, err := fetchBooks()
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	for _, book := range books {
		fmt.Printf("ID: %d\t Title: %s\t Author: %s\t\n", book.ID, book.Title, book.Author)
	}

	fmt.Println("\n")
	fmt.Println("========")
	fmt.Println("Post Book")
	fmt.Println("========")

	// Post book
	success, err := postBook()
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	fmt.Println(success)

	fmt.Println("\n")
	fmt.Println("========")
	fmt.Println("Fetch Single Book")
	fmt.Println("========")

	// Fetch data single book
	idBook := strconv.Itoa(3)
	book, err := fetchBookByID(idBook)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	fmt.Printf("ID: %d\t Title: %s\t Author: %s\t\n", book.ID, book.Title, book.Author)
}

func fetchBooks() ([]Book, error) {
	var client = &http.Client{}
	var books []Book

	// Create new request
	request, err := http.NewRequest("GET", baseURL+"/books", nil)
	if err != nil {
		return nil, err
	}

	// Execute request
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Decode from json to object Book
	err = json.NewDecoder(response.Body).Decode(&books)
	if err != nil {
		return nil, err
	}

	// Return books and error nil
	return books, nil
}

func fetchBookByID(ID string) (Book, error) {
	var client = &http.Client{}
	var book Book

	// Create new request
	request, err := http.NewRequest("GET", baseURL+"/book?id="+ID, nil)
	if err != nil {
		return book, err
	}

	// Execute request
	response, err := client.Do(request)
	if err != nil {
		return book, err
	}
	defer response.Body.Close()

	// Decode from json to object Book
	err = json.NewDecoder(response.Body).Decode(&book)
	if err != nil {
		return book, err
	}

	// Return books and error nil
	return book, nil
}

func postBook() (string, error) {
	var client = &http.Client{}

	data := Book{
		ID:     3,
		Title:  "Harry Potter",
		Author: "J. K. Rowling",
	}
	dataBody := fmt.Sprintf(`{"id": %d, "title": "%s", "author": "%s"}`, data.ID, data.Title, data.Author)
	requestBody := strings.NewReader(dataBody)

	// Create new request
	request, err := http.NewRequest("POST", baseURL+"/post-book", requestBody)
	if err != nil {
		return "", err
	}

	// Execute request
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Decode from json to object Book
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	// Return response string and error nil
	return string(respBody), nil
}
