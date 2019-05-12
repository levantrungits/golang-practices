package pointer

import (
	"fmt"
)

/* pointer receiver
	func (s *MyStruct) pointerMethod() { }
*/
/* value receiver
	func (s MyStruct)  valueMethod()   { }
*/

func TestPointer() {
	fmt.Println("==================================================")
	a := 555
	pt := &a
	fmt.Println(pt) // pt is a address value: 0xc000014100

	//fmt.Println("==================================================")
	// A literal has no address.
	//pt = &123  // error: "cannot take the address of 123"
	//fmt.Println(pt)
	// A constant has no address.
	//const c int = 123
	//pt = &c  // error: "cannot take the address of c"
	//fmt.Println(pt)

	fmt.Println("==================================================")
	// Retrieving a value through a pointe
	fmt.Println("*pt yields a's value:", *pt)

	fmt.Println("==================================================")
	i := 12
	p := &i 		// point to i
	fmt.Println("*p: ", *p) // read i through the pointer
	*p = 21			// set i through the pointer
	fmt.Println("new i: ", i)  // see the new value of i

	fmt.Println("==================================================")
	j := 100
	p = &j
	fmt.Println("*p: ", *p)
	*p = *p / 10
	fmt.Println("new j: ", j)

	fmt.Println("==================================================")
	// Pointer “side effect”
	// When two pointers point to the same variable and the variable is updated through one pointer
	// , then accessing the variable through the other pointer retrieves the updated value.
	p1 := &a
	p2 := &a
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
	fmt.Println("*p2 is:", *p2)
	*p1 = 444
	fmt.Println("*p2 is:", *p2)

	fmt.Println("==================================================")
	// Pointer comparisons
	fmt.Println("p1 == p2:", p1 == p2)
	fmt.Println("*p1 == *p2:", *p1 == *p2)
	
	fmt.Println("==================================================")
	// A pointer can also be compared to nil
	p = nil
    fmt.Println("p == nil:", p == nil)
    p = &a
    if p != nil {
        *p = 5
	}
	fmt.Println("p:", p)
	fmt.Println("*p:", *p)

	fmt.Println("==================================================")
	// The new function
	p = new(int) // default is 0 value.
	fmt.Println("p points to an unnamed int:", p, *p)
	*p = 6
	fmt.Println("The unnamed int has been changed to:", *p)

	// fmt.Println("==================================================")
	// No pointer arithmetic
	// p = &a + 64  // error: "invalid operation: &a + 64 (mismatched types *int and int)"
	// fmt.Println("p:", p)
	// fmt.Println("*p:", *p)

	// fmt.Println("==================================================")
	// Not covered in the video
}