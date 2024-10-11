package packet

import "fmt"

func packets(val int) {
	if val < 50 {
		fmt.Println("Its low")
	} else {
		fmt.Println("Its high")

	}
}