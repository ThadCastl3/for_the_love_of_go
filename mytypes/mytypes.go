package mytypes

import "strings"

// Twice multiplies its receiver by 2 and returns the result

type MyInt int

func (i MyInt) Twice() MyInt {
	return i * 2
}

type MyString string

// Len returns the length of the string it receives as a receiver
func (s MyString) Len() int {
	return len(s)
}

// type MyBuilder strings.Builder

func (mb MyBuilder) Hello() string {
	return "Hello, Gophers!"
}

// MyBuilder is a struct that contains a strings.Builder
type MyBuilder struct {
	Contents strings.Builder
}

// StringUppercaser is a struct that contains a strings.Builder and converts the string to uppercase
type StringUppercaser struct {
	Contents strings.Builder
}

func (su StringUppercaser) ToUpper() string {
	return strings.ToUpper(su.Contents.String())
}

// func Double(input int) {
// 	input *= 2
// }

// modified func Double to use the int pointer *int
// func Double(input *int) {
// 	input *= 2
// }

//fixed the mismatch types
// func Double(input *int) {
// 	*input *= 2
// }

func (input *MyInt) Double() {
	*input *= 2 // dereference the pointer and multiply by 2
}
