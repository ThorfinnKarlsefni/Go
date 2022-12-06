package struct_func

import (
	"fmt"
	"time"
)

type Address struct {
	Street           string
	HouseNumber      uint32
	HouseNumberAddOn string
	POBox            string
	ZipCode          string
	City             string
	Country          string
}

type VCard struct {
	FirstName string
	LastName  string
	NickName  string
	BirDate   time.Time
	Photo     string
	Address   map[string]*Address
}

func Vcard() {
	addr1 := &Address{"VINLAND", 12, "", "", "2600", "Mechelen", "Belgie"}
	addr2 := &Address{"Heideland", 28, "", "", "2640", "Mortsel", "Belgie"}
	addrs := make(map[string]*Address)
	addrs["youth"] = addr1
	addrs["now"] = addr2

	birthdt := time.Date(1956, 1, 17, 15, 4, 5, 0, time.Local)
	photo := "MyDocuments/MyPhotos/photo1.jpg"
	vcard := &VCard{"Ivo", "Balbert", "", birthdt, photo, addrs}
	fmt.Printf("Here is the full Vcard %v\n", vcard)
	fmt.Printf("My address are: \n %v\n %v", addr1, addr2)
}
