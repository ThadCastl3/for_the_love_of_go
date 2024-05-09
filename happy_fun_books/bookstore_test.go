package bookstore_test

import (
	"bookstore"
	"testing"
)

// bookstore.Book{
// 	Title: "For The Love of Go",
// 	Author: "The Gopher",
// 	Copies: 100,
// }

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "For The Love of Go",
		Author: "The Gopher",
		Copies: 100,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "Le Go",
		Author: "ThadCastl3",
		Copies: 7,
	}
	want := 6
	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies
	if want != got {
		t.Errorf("want %d, got %d after buying 1 copy from a stock of %d", want, got, b.Copies)
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:  "I'm Going to Bed",
		Author: "ThadCastl3",
		Copies: 0,
	}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Error("Expected an error when buying a book with no copies left, got nil instead")
	}
}
