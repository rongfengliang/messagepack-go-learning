//go:generate msgp

package pkg

// Person person for search
type Person struct {
	Name       string `msg:"name" msgpack:"name"`
	Address    string `msg:"address" msgpack:"address"`
	Age        int    `msg:"age" msgpack:"age"`
	Hidden     string `msg:"-" msgpack:"_"` // this field is ignored
	unexported bool   // this field is also ignored
}
