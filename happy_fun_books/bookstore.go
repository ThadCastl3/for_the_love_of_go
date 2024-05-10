package bookstore

import "errors"

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

func GetAllBooks(catalog []Book) []Book {
	return catalog
}

func GetBook(catalog []Book, ID int) Book {
	return Book{}
}
