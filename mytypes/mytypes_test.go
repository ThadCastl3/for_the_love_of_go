package mytypes_test

import (
	"mytypes"
	"testing"
	"strings"
)

func TestTwice(t *testing.T) {
	t.Parallel()
	input := mytypes.MyInt(9)
	want := mytypes.MyInt(18)
	got := input.Twice()
	if want != got {
		t.Errorf("twice %d: want %d, got %d", input, want, got)
	}
}

func TestMyStringLen(t *testing.T) {
	t.Parallel()
	input := mytypes.MyString("Hello, Gophers!")
	want := 15
	got := input.Len()
	if want != got {
		t.Errorf("len(%s): want %d, got %d", input, want, got)
	}
}

func TestStringBuilder(t *testing.T) {
	t.Parallel()
	var sb strings.Builder
	sb.WriteString("Hello, ")
	sb.WriteString("Gophers!")
	want := "Hello, Gophers!"
	got := sb.String()
	if want != got {
		t.Errorf("got %q, want %q", got, want)
	}
	wantlen := 15
	gotlen := sb.Len()
	if wantlen != gotlen {
		t.Errorf("%q: want len %d, got %d", sb.String(), wantlen, gotlen)
	}
}

// func TestMyBuilderlen(t *testing.T) {
// 	t.Parallel()
// 	var mb mytypes.MyBuilder
// 	want := 15
// 	got := mb.Len()
// 	if want != got {
// 		t.Errorf("want %d, got %d", want, got)
// 	}
// }

// func TestMyBuilderHello(t *testing.T) {
// 	t.Parallel()
// 	var mb mytypes.MyBuilder
// 	want := "Hello, Gophers!"
// 	got := mb.Hello()
// 	if want != got {
// 		t.Errorf("want %q, got %q", want, got)
// 	}
// }

func TestMyBuilder(t *testing.T) {
	t.Parallel()
	var mb mytypes.MyBuilder
	mb.Contents.WriteString("Hello, ")
	mb.Contents.WriteString("Gophers!")
	want := "Hello, Gophers!"
	got := mb.Contents.String()
	if want != got {
		t.Errorf("got %q, want %q", got, want)
	}
	wantlen := 15
	gotlen := mb.Contents.Len()
	if wantlen != gotlen {
		t.Errorf("%q: want len %d, got %d", mb.Contents.String(), wantlen, gotlen)
	}
}

func TestStringUppercaser(t *testing.T) {
	t.Parallel()
	var su mytypes.StringUppercaser
	su.Contents.WriteString("Hello, Gophers!")
	want := "HELLO, GOPHERS!"
	got := su.ToUpper()
	if want != got {
		t.Errorf("got %q, want %q", got, want)
	}
}

// func TestDouble(t *testing.T) {
// 	t.Parallel()
// 	var x int = 12
// 	want := 24
// 	mytypes.Double(x)
// 	if want != x {
// 		t.Errorf("want %d, got %d", want, x)
// 	}
// }

// modified this test to use a pointer and a sharing operator
// func TestDouble(t *testing.T) {
// 	t.Parallel()
// 	var x int = 12
// 	want := 24
// 	mytypes.Double(&x)
// 	if want != x {
// 		t.Errorf("want %d, got %d", want, x)
// 	}
// }

// modifying test to use a pointer for the MyInt type
//
