package calc

import(
"fmt"
)
func Add(n ...int) int{
	if  len(n)<2{
	fmt.Println("numbers should be greater than 2")
	return -1
	}
	var sum int
	for _,val := range n{
		sum += val
	}
return  sum
}
