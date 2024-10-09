package utils

func Convert( x int) ([]int ,int){
	a := [] int {}
	val := x
	steps := Steps(x);
	
	for i := 0 ; val > 0 ; i++ {
		a = append(a, val%2)
		val /= 2
	}
	j := len(a) - 1
	for i := 0 ; i < j ; i++ {
		a[i] , a[j] = a[j] , a[i]
		j--
	}
	capacity :=  steps - len(a)
	b := make([]int, capacity + len(a))
	copy(b[capacity:], a)
	return b, steps
	
}
func Steps(x int) int{
	steps := 1
	for i := 0;x > 0; i++ {
		x /= 2
		steps++
	}
	return steps

}

func Not(x []int) []int {
   a := x
   not := [] int {}
   for i := 0 ; i < len(a) ; i++ {
	if a[i] == 1 {
		not = append(not, 0)
	} else if a[i] == 0 {
		not = append(not, 1)
	} else {
		break
	}
   } 
   return not
}

func Unsigned(x []int ) []int {
	
	a := []int{}
	
	if (len(x) == 0) {
		
			return []int{}
		
	}
	carry := 1
	for j := len(x) - 1 ; j >= 0 ; j-- {
		if(x[j] == 1 ) {
			if carry == 1 {
				a = append(a, 0)
				carry = 1
			} else if carry == 0 {
				a = append(a, 1)
				carry = 0
			}
			
		} else if (x[j] == 0) {
			if carry == 1 {
				a = append(a, 1)
				carry = 0
			} else if carry == 0 {
				a = append(a, 0)
				carry = 0
			}
		} 

	}
	j := len(a) - 1
	for i := 0 ; i < j ; i++ {
		a[i] , a[j] = a[j] , a[i]
		j--
	}

	return a

}
