package exercise

import (
	"fmt"
	"strconv"
)

type Stack1 struct {
	ix   int
	data [LIMIT]int
}

func StackStruct() {
	st1 := new(Stack1)
	fmt.Printf("%v\n", st1)
	st1.Push(3)
	fmt.Printf("%v\n", st1)
	st1.Push(7)
	fmt.Printf("%v\n", st1)
	st1.Push(8)
	fmt.Printf("%v\n", st1)
	st1.Push(10)
	fmt.Printf("%v\n", st1)
	st1.Pop()
	fmt.Printf("%v\n", st1)
}

func (st *Stack1) Push(n int) {
	if st.ix+1 > LIMIT {
		return
	}

	st.data[st.ix] = n
	st.ix++
}

func (st *Stack1) Pop() int {
	st.ix--
	return st.data[st.ix]
}

func (st Stack1) String() string {
	str := ""
	for ix := 0; ix < st.ix; ix++ {
		str += "[" + strconv.Itoa(ix) + ":" + strconv.Itoa(st.data[ix]) + "]"
	}

	return str
}
