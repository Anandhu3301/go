package main

import (
	"fmt"
	"go_basic/utils"
)

func main() {
	

	// var a int
	// fmt.Println("Enter a number to convert in unsigned")
	// fmt.Scanln(&a)
	// y, _ := utils.Convert(a)
	// fmt.Println("Binary of", a, "is", y)
	// z := utils.Not(y)
	// fmt.Println("Complement of binary is ", z)

	// fmt.Println("Unsigned value of ", a, "is", utils.Unsigned(z))

	// dayofweek := [...]string{
	// 	"sun","mon",
	// }
	// dayofweek[1] = "hyt"
	// slice := dayofweek[:]
	// slice = append(slice, "spoi")
	// fmt.Println(slice)

	var s *utils.Stack = new(utils.Stack)
	var maxsize int
	fmt.Println("What should be the length of stack")
	fmt.Scanln(&maxsize)         
	s.Initialize(maxsize)

	s.Push(4)
	s.Push(3)
	s.Push(8)

	for !s.IsEmpty() {
		fmt.Println("Top element",s.Peek())
		fmt.Println("Popped element",s.Pop())
	}



}
