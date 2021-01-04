package main

import (
	"bytes"
	"demoapp/pkg"
	"log"

	"github.com/tinylib/msgp/msgp"
	"github.com/vmihailenco/msgpack/v5"
)

func main() {
	var buf bytes.Buffer

	myPerson := pkg.Person{
		Name:    "荣锋亮",
		Address: "beijing",
		Age:     33,
	}
	err := msgp.Encode(&buf, &myPerson)
	if err != nil {
		panic("encode some wrong" + err.Error())
	}
	var dstPerson pkg.Person
	err = msgpack.Unmarshal(buf.Bytes(), &dstPerson)
	if err != nil {
		panic("uncode:" + err.Error())
	}
	log.Println("from msgp: ", string(buf.Bytes()))
	log.Printf("%v", dstPerson)
	log.Println("from msgpack:", dstPerson.Name)
}
