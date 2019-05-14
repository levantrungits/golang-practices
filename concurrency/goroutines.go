package concurrency

import (
	"runtime"
	"time"
	"fmt"
)

/*
	1. Create goroutines
	2. Introduce parallelism
*/
func GoroutinesBasic() {
	godur, _ := time.ParseDuration("10ms")

	runtime.GOMAXPROCS(2)

	go func(){
		for idx := 0; idx < 100; idx++ {
			fmt.Println("Hello")	
			time.Sleep(godur)
		}
	}()

	go func(){
		for jdx := 0; jdx < 100; jdx++ {
			fmt.Println("Go")
			time.Sleep(godur)
		}
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)

	fmt.Println("=======> DONE.")
}

