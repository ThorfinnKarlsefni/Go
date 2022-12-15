package chapter12

import (
	"bufio"
	"fmt"
	"go/basic/exercise/chapter12/stack"
	"os"
	"strconv"
)

func Calculator()  {
	buf := bufio.NewReader(os.Stdin)
	fmt.Println("Give a numbers,or operator(+,-,/,*),or p of stop")
	calc1 := new(stack.Stack)
	for {
		input, err := buf.ReadString('\n')

		if err != nil {
			fmt.Println("Input Error")
			os.Exit(1)
		}
		//fmt.Printf("%s---",input)
		token := input[:len(input)-1]
		switch  {
		case token == "q":
			fmt.Println("Goodbye")
			return
		case token > "-99999" && token < "99999":
			i,_ := strconv.Atoi(token)
			calc1.Push(i)
		case token == "+":
			q := calc1.Pop()
			p := calc1.Pop()
			fmt.Printf("this result %d %s %d = %d\n",p,token,q,q+p)
		default:
			fmt.Println("unknown error")
		}

	}
}