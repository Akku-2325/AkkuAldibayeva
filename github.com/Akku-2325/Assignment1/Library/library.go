package Library

import (
	"errors"
	"fmt"
)

type Library struct {
	Books map[string]*Book
}

func NewLibrary() *Library {
	return &Library{Books: make(map[string]*Book)}
}

func (l *Library) AddBook(book Book) {
	l.Books[book.ID] = &book
}

func (l *Library) BorrowBook(id string) error {
	if book, exists := l.Books[id]; exists {
		if book.IsBorrowed {
			return errors.New("this book is already borrowed")
		}
		book.IsBorrowed = true
		return nil
	}
	return errors.New("book not found")
}

func (l *Library) ReturnBook(id string) error {
	if book, exists := l.Books[id]; exists {
		if !book.IsBorrowed {
			return errors.New("this book was not borrowed")
		}
		book.IsBorrowed = false
		return nil
	}
	return errors.New("book not found")
}

func (l *Library) ListBooks() {
	fmt.Println("List of available books:")
	anyAvailable := false
	for _, book := range l.Books {
		if !book.IsBorrowed {
			anyAvailable = true
			fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
	if !anyAvailable {
		fmt.Println("No available books.")
	}
}
