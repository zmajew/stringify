package stringify_test

import (
	"fmt"
	"testing"

	"github.com/zmajew/stringify"
)

type A struct {
	first int
}

type C interface {
	Read()
}

func (a *A) Read() {}

type B struct {
	First  int
	first  int
	s      []int
	m      map[string]int
	ptr    *A
	intrfc C
	fnc    func()
}

func (b B) Read() {}

func TestToString(t *testing.T) {
	s := []int{1, 3}
	m := make(map[string]int)
	m["g"] = 5
	a := A{first: 4}
	b := B{
		First:  1,
		first:  2,
		s:      s,
		m:      m,
		ptr:    &a,
		intrfc: &a,
		fnc:    a.Read,
	}
	// b.intrfc = b
	fmt.Println("string:", stringify.ToString(b))
}
