package packet

import "fmt"

func Packet(val int) {
	if val < 50 {
		fmt.Println("Its low")
	} else {
		fmt.Println("Its high")

	}
}