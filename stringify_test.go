package stringify_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/uuid"
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
	bts    []byte
	jsn    json.RawMessage
	tm     time.Time
	uid    uuid.UUID
}

func (b B) Read() {}

func TestToString(t *testing.T) {
	s := []int{1, 3}
	m := make(map[string]int)
	m["g"] = 5
	a := A{first: 4}
	bts := []byte("some text")
	b := B{
		First:  1,
		first:  2,
		s:      s,
		m:      m,
		ptr:    &a,
		intrfc: &a,
		fnc:    a.Read,
		bts:    bts,
		tm:     time.Now(),
		uid:    uuid.New(),
	}
	jsn, err := json.Marshal(b)
	if err != nil {
		t.Fatal(err)
	}
	b.jsn = jsn

	result := stringify.Parse(b)
	ok := json.Valid([]byte(result))
	if !ok {
		t.Fatal("result string not a valid json", result)
	}
}
