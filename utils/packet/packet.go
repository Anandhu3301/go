package packet

import "fmt"

func Packet12(val int) {
	if val < 50 {
		fmt.Println("Its low")
	} else {
		fmt.Println("Its high")

	}
}