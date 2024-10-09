package main

import (
	"fmt"
	"go_basic/utils"
)

func main() {
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("Its a weekend")
	// default:
	// 	fmt.Println("Its a weekday")
	// }

	// t:= time.Now().Hour()

	// switch {
	// case t == 10:
	// 	fmt.Println("Its before noon")
	// default:
	// 	fmt.Println("Its after noon", t)
	// }

	// whatAmi := func(i interface{}) {
	// 	switch t:= i.(type) {
	// 	case bool:
	// 		fmt.Println("I'm a boolean")
	// 	case int:
	// 		fmt.Println("I'm a int")
	// 	default:
	// 		fmt.Printf("Dont know type %T",t)

	// 	}

	// }
	// whatAmi(true)
	// whatAmi(1)
	// whatAmi("HEY")
	// var a[5]int
	// a[4] = 790
	// fmt.Println(a)
	// b := [...]string{"hi", "hello"}
	// fmt.Println(b)
	// b  = [...]string{"hi",1:"hello"}
	// var twoD [2][3]int
	// for i := 0; i < 2 ; i++ {
	// 	for j := 0; j < 3 ; j++ {
	// 		twoD[i][j] = i+ j
	// 	}
	// }
	// fmt.Println(twoD)
	// var s []string
	// fmt.Println("value:",s,s==nil, len(s)==0)

	// s = make([]string, 3)
	// // fmt.Println("emp:",s,"len:",len(s),"capacity:",cap(s))

	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"

	// s = append(s, "d")
	// s = append(s, "e", "f")

	// c := make([] string , len(s))

	// copy(c,s)
	// fmt.Println(c)

	// l := s[:5]
	// fmt.Println(l)

	// p := s[2:]
	// fmt.Println(p)

	// twoD := make([][]int , 3)

	// for i := 0; i < 3 ; i++ {
	// 	innerlen := i + 1
	// 	twoD[i] = make([]int, innerlen)
	// 	for j := 0; j < 3 ; j++ {
	// 		twoD[i][j] = i + j
	// 	}
	// }

	// m := make(map[string] int)
	// m["k1"] = 89
	// m["k3"] = 74

	// fmt.Println(m)

	// v1 := m["k1"]
	// fmt.Println(v1)

	// for  key, value :=  range m {
	// 	fmt.Println("key:",key, "Value:",value)

	// }

	// value, exists := m["k1"]
	// if exists && value > 0{
	// 	fmt.Println("The value is ", value)
	// } else{
	// 	fmt.Println("Couldnt find the value")
	// }

	// value, prs := m["k1"]
	// if prs {
	// 	fmt.Println("prs:",value )

	// }
	// delete(m, "k1")
	// fmt.Println(m)

	// clear(m)
	// fmt.Println(m)
	// s,r := utils.Hello()
	// fmt.Println(s,r)
	// a, _ := utils.Vals()
	// fmt.Println(a)

	var a int
	fmt.Println("Enter a number to convert in unsigned")
	fmt.Scanln(&a)
	y, _ := utils.Convert(a)
	fmt.Println("Binary of", a, "is", y)
	z := utils.Not(y)
	fmt.Println("Complement of binary is ", z)

	fmt.Println("Unsigned value of ", a, "is", utils.Unsigned(z))

	// dayofweek := [...]string{
	// 	"sun","mon",
	// }
	// dayofweek[1] = "hyt"
	// slice := dayofweek[:]
	// slice = append(slice, "spoi")
	// fmt.Println(slice)

}
