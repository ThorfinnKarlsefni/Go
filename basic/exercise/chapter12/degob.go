package chapter12

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

var content string
var vc VCard

func Degob() {
	file, _ := os.Open("basic/example/chapter12/vcard.god")
	defer file.Close()
	inReader := bufio.NewReader(file)
	dec := gob.NewDecoder(inReader)
	err := dec.Decode(&vc)
	if err != nil {
		log.Println("error in decoding gob")
	}
	fmt.Println(vc)
}
