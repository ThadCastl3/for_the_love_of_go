package creditcard_test

import (
	"creditcard"
	"testing"
)

func TestNewCardNumber(t *testing.T) {
	t.Parallel()
	want := "1234123412341234"
	cc, err := creditcard.NewCard(want)
	if err != nil {
		t.Fatal(err)
	}
	got := cc.Number()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestNewCardInvalid(t *testing.T) {
	t.Parallel()
	_, err := creditcard.NewCard("")
	if err == nil {
		t.Fatal("want error for invalid card number, got nil")
	}
}
