package chapter12

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type Address1 struct {
	Type    string
	City    string
	Country string
}

type VCard1 struct {
	FirstName string
	LastName  string
	Address   []*Address1
	Remark    string
}

var content string

var vc VCard1

func Gob2() {
	pa := &Address1{"private", "Aartselaar", "Belgium"}
	wa := &Address1{"work", "Boom", "Belgium"}
	vc := VCard1{"Jan", "Kersschot", []*Address1{pa, wa}, "none"}

	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)

	if err != nil {
		log.Println("Error in encoding gob")
	}
}

func Degob() {
	file, _ := os.Open("vcard.gob")
	defer file.Close()
	inReader := bufio.NewReader(file)
	dec := gob.NewDecoder(inReader)
	err := dec.Decode(&vc)
	if err != nil {
		log.Println("error in decoding gob")
	}
	fmt.Println(vc)
}
