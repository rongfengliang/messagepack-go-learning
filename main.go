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
	var dstPerson2 *pkg.Person = &pkg.Person{}
	err = msgpack.Unmarshal(buf.Bytes(), &dstPerson)
	datas, err := dstPerson2.UnmarshalMsg(buf.Bytes())
	if err != nil {
		panic("uncode:" + err.Error())
	}
	if len(datas) > 0 {
		log.Panicf("%d bytes left over after UnmarshalMsg(): %q", len(datas), datas)
	}
	log.Println("from msgp: ", string(buf.Bytes()))
	log.Printf("msgpack:%v,msgp: %v", dstPerson, *dstPerson2)
	log.Println("from msgpack:", dstPerson.Name)
}
