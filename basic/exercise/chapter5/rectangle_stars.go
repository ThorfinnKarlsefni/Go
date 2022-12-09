package exercise

import "fmt"

func Rectangle() {
	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("*")
	// 	if i == 1 || i == 10 {

	// 		fmt.Print("**********************")

	// 	} else {
	// 		fmt.Print("                     ")
	// 	}
	// 	fmt.Println("*")
	// }

	w, h := 20, 10

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
