package control_struct

import "fmt"

func For6() {
LABEL1:

	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is : %d,and j is : %d\n", i, j)
		}
	}
}
