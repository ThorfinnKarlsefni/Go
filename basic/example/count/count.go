package count

import (
	"fmt"
	"strings"
)

func Init() {
	var str string = "Hello,how is it going, Hugo?"

	var manyG = "ggggggg"

	fmt.Printf("Number of H's in %s is:", str)
	fmt.Printf("%d\n", strings.Count(str, "H"))

	fmt.Printf("Number of double g's in %s is:", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))
}
