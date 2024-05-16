package bookstore

import (
	"errors"
	"fmt"
)

// Book represents information about a book.
type Book struct {
	Title string
	Author string
	Copies int
	ID int
}


// Buy decrements the number of copies of a book by 1.
func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("No copies remaining.")
	}
	b.Copies--
	return b, nil
}

// func GetAllBooks(catalog []Book) []Book {
// 	return catalog
// }
// refactoring to use a map instead of a slice

func GetAllBooks(catalog map[int]Book) []Book {
	result := []Book{}
	for _, b := range catalog {
		result = append(result, b)
	}
	return result
}


// func GetBook(catalog []Book, ID int) Book {
// 	for _, b := range catalog {
// 		if b.ID == ID {
// 			return b
// 		}
// 	}
// 	return Book{}
// }

// func GetBook(catalog map[int]Book, ID int) Book {
// 	return catalog[ID]
// }

// func GetBook(catalog map[int]Book, ID int) (Book, error) {
// 	return catalog[ID], nil
// }

// func GetBook(catalog map[int]Book, ID int) (Book, error) {
// 	if _, ok := catalog[ID]; !ok {
// 		return Book{}, errors.New("Book not found.")
// 	}
// 	return Book{}, nil
// }
// scrapped this function for a more informative version that returns the ID of the book

func GetBook(catalog map[int]Book, ID int) (Book, error) {
    b, ok := catalog[ID]
    if !ok {
        return Book{}, fmt.Errorf("Book not found: %d", ID)
    }
    return b, nil
}

// func AddBook()
