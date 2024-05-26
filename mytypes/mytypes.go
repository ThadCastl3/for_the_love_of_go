package mytypes

// Twice multiplies its receiver by 2 and returns the result

type MyInt int

func (i MyInt) Twice() MyInt {
	return i * 2
}

type MyString string

func (s MyString) Len() int {
	return len(s)
}
