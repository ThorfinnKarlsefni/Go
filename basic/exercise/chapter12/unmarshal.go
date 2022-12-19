package chapter12

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var v VCard

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Address   []*Address
	Remark    string
}

func Unmarshal() {
	data, err := os.Open("../../example/chapter12/vcard.json")
	if err != nil {
		fmt.Println(nil)
		return
	}
	temp, _ := ioutil.ReadAll(data)
	json.Unmarshal(temp, &v)
	fmt.Println(v)

}
