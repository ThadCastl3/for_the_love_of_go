package bookstore

import (
	"errors"
	"fmt"
)

// Book represents information about a book.
// type Book struct {
// 	Title string
// 	Author string
// 	Copies int
// 	ID int
// }

// creating a book object price in cents and a discount field
type Book struct {
	Title           string
	Author          string
	Copies          int
	ID              int
	PriceCents      int
	DiscountPercent int
	category        string
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

// func GetAllBooks(catalog map[int]Book) []Book {
// 	result := []Book{}
// 	for _, b := range catalog {
// 		result = append(result, b)
// 	}
// 	return result
// }

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

// func GetBook(catalog map[int]Book, ID int) (Book, error) {
//     b, ok := catalog[ID]
//     if !ok {
//         return Book{}, fmt.Errorf("Book not found: %d", ID)
//     }
//     return b, nil
// }

// func AddBook()

// func NetPriceCents(b Book) int {
// 	// check if DiscountPercent exists
// 	if b.DiscountPercent != 0 {
// 		return b.PriceCents - (b.PriceCents * b.DiscountPercent / 100)
// 	}
// 	return b.PriceCents
// }

func (b Book) NetPriceCents() int {
	// check if DiscountPercent exists
	if b.DiscountPercent != 0 {
		return b.PriceCents - (b.PriceCents * b.DiscountPercent / 100)
	}
	return b.PriceCents
}

// Catalog Type
type Catalog map[int]Book

func (c Catalog) GetAllBooks() []Book {
	result := []Book{}
	for _, b := range c {
		result = append(result, b)
	}
	return result
}

func (c Catalog) GetBook(ID int) (Book, error) {
	b, ok := c[ID]
	if !ok {
		return Book{}, fmt.Errorf("Book not found: %d", ID)
	}
	return b, nil
}

// method for validating book discounts (non working, since we need to modify the receiver and the receiver must be a pointer)
// func (b Book) SetPriceCents(price int) error {
// 	b.PriceCents = price
// 	return nil
// }

// method for validating book discounts (working, since we modify the receiver and the receiver must be a pointer) but it does not return an error if the price is invalid
// func (b *Book) SetPriceCents(price int) error {
// 	b.PriceCents = price
// 	return nil
// }

func (b *Book) SetPriceCents(price int) error {
	if price < 0 {
		return fmt.Errorf("negative price %d", price)
	}
	b.PriceCents = price
	return nil
}

// method for retrieving the category of a book
func (b *Book) SetCategory(category string) error {
	if category != "Autobiography" {
		return fmt.Errorf("unsupported category %q", category)
	}
	b.category = category
	return nil
}

func (b Book) Category() string {
	return b.category
}
