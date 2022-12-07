package struct_func

import (
	"fmt"
	"reflect"
)

type TagType struct {
	field1 bool   "An important answer"
	field2 string "The name of the thing"
	field3 int    "How much there are"
}

func Tag() {
	//tt := TagType{true, "Barak Obama", 1}
	//fmt.Println(tt)
	for i := 0; i < 3; i++ {
		refTag(TagType{}, i)
	}
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)

	fmt.Printf("%v\n", ixField.Tag)
}
