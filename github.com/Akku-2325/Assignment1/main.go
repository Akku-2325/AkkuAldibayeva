package main

import (
	"awesomeProject/github.com/Akku-2325/Assignment1/Library"
	"bufio"
	"fmt"
	"os"
)

func main() {
	lib := Library.NewLibrary()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		var command string
		fmt.Println("Enter command (Add, Borrow, Return, List, Exit):")
		scanner.Scan()
		command = scanner.Text()

		switch command {
		case "Add":
			var id, title, author string
			fmt.Println("Enter book ID:")
			scanner.Scan()
			id = scanner.Text()
			fmt.Println("Enter book title:")
			scanner.Scan()
			title = scanner.Text()
			fmt.Println("Enter book author:")
			scanner.Scan()
			author = scanner.Text()
			book := Library.Book{ID: id, Title: title, Author: author, IsBorrowed: false}
			lib.AddBook(book)
			fmt.Println("Book added successfully.")
		case "Borrow":
			var id string
			fmt.Println("Enter book ID to borrow:")
			scanner.Scan()
			id = scanner.Text()
			if err := lib.BorrowBook(id); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Book borrowed successfully.")
			}
		case "Return":
			var id string
			fmt.Println("Enter book ID to return:")
			scanner.Scan()
			id = scanner.Text()
			if err := lib.ReturnBook(id); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Book returned successfully.")
			}
		case "List":
			lib.ListBooks()
		case "Exit":
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid command. Try again.")
		}
	}
}
