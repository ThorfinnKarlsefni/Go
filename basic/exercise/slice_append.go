package exercise

import (
	"bytes"
	"fmt"
)

var buffer bytes.Buffer

func SliceAppend() {
	sl := []byte{'a', 'u', 'z', 's'}

	str := []byte{'g', 'o', 'l', 'a', 'n'}

	//slice := Append(sl, str)
	//fmt.Printf("%s", slice)

	slice := WriteAppend(sl, str)
	bPrefix, bSuffix := buffer.Bytes()[:4], buffer.Bytes()[4:]
	fmt.Printf("%s\n", slice)
	fmt.Printf("%s,%s", bPrefix, bSuffix)
}

// buffer

func WriteAppend(slice, data []byte) []byte {

	buffer.Write(slice)
	buffer.Write(data)
	return buffer.Bytes()
}

// for

func Append(slice, data []byte) []byte {

	lengthNewSlice := len(slice) + len(data)
	capNewSlice := cap(slice)

	if lengthNewSlice > cap(slice) {
		capNewSlice = lengthNewSlice
	}

	newSlice := make([]byte, lengthNewSlice, capNewSlice)

	for k, v := range slice {
		newSlice[k] = v
	}

	for dataKey, dataValue := range data {
		newSlice[dataKey+len(slice)] = dataValue
	}

	return newSlice
}
