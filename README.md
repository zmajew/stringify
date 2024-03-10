# Stringify

This parser should stringify anything into a json valid string. 

Options map contains instruction functions for defining a patterns of how to parse some user specific or common types. 

Defaults for some common types are defined here.

Example of usage:

```
package main

import (
	"fmt"
	"reflect"

	"github.com/zmajew/stringify"
)

type A struct {
	SomeField string
	field     string
}

func main() {
	a := A{
		SomeField: "test_string",
		field:     "field_2",
	}

	aString := stringify.Parse(a)
	fmt.Println(aString)
}
```
output:
```
{"SomeField": "test_string", "field": "field_2"}
```

with options:
```
package main

import (
	"fmt"
	"reflect"

	"github.com/zmajew/stringify"
)

func init() {
	stringify.Options["main.Password"] = func(v interface{}) string {
		return "******"
	}
}

type Password string

type A struct {
	Username string
	Password Password
}

func main() {
	a := A{
		Username: "admin",
		Password: "admin",
	}

	aString := stringify.Parse(a)
	fmt.Println(aString)
}
```
output
```
{"SomeField": "test_string", "Password": ******}
```
