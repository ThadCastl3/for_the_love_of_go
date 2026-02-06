package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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

// func TestBuyTrivial(t *testing.T) {
// 	t.Parallel()
// 	bookstore.Buy(bookstore.Book{})
// 	bookstore.Buy(bookstore.Book{Copies: 1})
// }

// func TestGetAllBooks(t *testing.T) {
// 	t.Parallel()
// 	catalog := []bookstore.Book{
// 		{Title: "McMuffins and Puffins"},
// 		{Title: "El Chupacabra y El Gato"},
// 	}
// 	want := []bookstore.Book{
// 		{Title: "McMuffins and Puffins"},
// 		{Title: "El Chupacabra y El Gato"},
// 	}
// 	// want := []string{"same", "same", "same"}
// 	// got := []string{"same", "different", "same"}
// 	got := bookstore.GetAllBooks(catalog)
// 	if !cmp.Equal(want, got) {
// 		t.Error(cmp.Diff(want, got))
// 	}
// }
// refactoring to use a map instead of a slice

// func TestGetAllBooks(t *testing.T) {
// 	t.Parallel()
// 	catalog := map[int]bookstore.Book{
// 		1: {ID: 1, Title: "McMuffins and Puffins"},
// 		2: {ID: 2, Title: "El Chupacabra y El Gato"},
// 	}
// 	want := []bookstore.Book{
// 		{ID: 1, Title: "McMuffins and Puffins"},
// 		{ID: 2, Title: "El Chupacabra y El Gato"},
// 	}
// 	got := bookstore.GetAllBooks(catalog)
// 	//sorting the slice to match our want
// 	sort.Slice(got, func(i, j int) bool {
// 		return got[i].ID < got[j].ID
// 	})
// 	if !cmp.Equal(want, got) {
// 	t.Error(cmp.Diff(want, got))
// 	}
// }

// refactoring to utilized custom defined type and method
func TestGetAllBooks(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "McMuffins and Puffins"},
		2: {ID: 2, Title: "El Chupacabra y El Gato"},
	}
	want := []bookstore.Book{
		{ID: 1, Title: "McMuffins and Puffins"},
		{ID: 2, Title: "El Chupacabra y El Gato"},
	}
	got := catalog.GetAllBooks()
	//sorting the slice to match our want
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

// func TestGetBook(t *testing.T) {
// 	t.Parallel()
// 	// catalog := []bookstore.Book{
// 	// 	{ID: 1, Title: "Le Monkai"},
// 	// 	{ID: 2, Title: "The Second Coming of The Apes"},
// 	// }
// 	catalog := map[int]bookstore.Book{
// 		1: {ID: 1, Title: "Le Monkai"},
// 		2: {ID: 2, Title: "The Second Coming of The Apes"},
// 	}
// 	// want := bookstore.Book{ID: 1, Title: "Le Monkai"}
// 	want := bookstore.Book{ID: 2, Title: "The Second Coming of The Apes"}
// 	// got := bookstore.GetBook(catalog, 1)
// 	// got := bookstore.GetBook(catalog, 2)
// 	got, err := bookstore.GetBook(catalog, 2)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if !cmp.Equal(want, got) {
// 		t.Error(cmp.Diff(want, got))
// 	}
// }

// refactoring TestGetBook to use custom defined type and method
func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "Le Monkai"},
		2: {ID: 2, Title: "The Second Coming of The Apes"},
	}
	want := bookstore.Book{ID: 2, Title: "The Second Coming of The Apes"}
	got, err := catalog.GetBook(2)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookBadIDReturnsError(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{}
	_, err := catalog.GetBook(42069)
	if err == nil {
		t.Error("Expected an error when getting a book with a bad ID, got nil instead")
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:           "The Price is Right",
		PriceCents:      4000,
		DiscountPercent: 25,
	}
	want := 3000
	got := b.NetPriceCents()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:      "For The Love of Go",
		PriceCents: 4000,
	}
	want := 3000
	err := b.SetPriceCents(want)
	if err != nil {
		t.Fatal(err)
	}
	got := b.PriceCents
	if want != got {
		t.Errorf("want updated price %d, got %d", want, got)
	}
}

func TestSetPriceCentsInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title:      "For the Love of Go",
		PriceCents: 4000,
	}
	err := b.SetPriceCents(-1)
	if err == nil {
		t.Fatal("want error setting invalid price -1, got nil")
	}
}

// func TestSetCategory(t *testing.T) {
// 	t.Parallel()
// 	b := bookstore.Book{
// 		Title:    "For The Love of Go",
// 		category: "Science Fiction",
// 	}
// 	want := "Autobiography"
// 	err := b.SetCategory(want)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	got := b.category
// 	if want != got {
// 		t.Errorf("want updated category %s, got %s", want, got)
// 	}
// }

func TestSetCategoryInvalid(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "For the Love of Go",
	}
	err := b.SetCategory("Science Fiction")
	if err == nil {
		t.Fatal("want error for invalid category, got nil")
	}
}

func TestSetCategory(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "For the Love of Go",
	}
	want := "Autobiography"
	err := b.SetCategory(want)
	if err != nil {
		t.Fatal(err)
	}
	got := b.Category()
	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Errorf("want category %q, got %q", want, got)
	}
}
