package pointer

import (
	"fmt"
)

func TestPointer() {
	i := 12

	p := &i 		// point to i
	fmt.Println("*p: ", *p) // read i through the pointer
	*p = 21			// set i through the pointer
	fmt.Println("new i: ", i)  // see the new value of i

	j := 100
	p = &j
	fmt.Println("*p: ", *p)
	*p = *p / 10
	fmt.Println("new j: ", j)
}