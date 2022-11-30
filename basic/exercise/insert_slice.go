package exercise

import "fmt"

func SliceInsert() {
	s := []string{"M", "N", "O", "P", "Q", "R"}
	in := []string{"A", "B", "C"}
	res := InserStringSlice(s, in, 0)
	fmt.Println(res)
	res = InserStringSlice(s, in, 3)
	fmt.Println(res)
}

func InserStringSlice(slice []string, insertion []string, index int) []string {
	result := make([]string, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	fmt.Println(at)
	at += copy(result[at:], insertion)
	fmt.Println(at)
	copy(result[at:], slice[index:])
	return result
}
