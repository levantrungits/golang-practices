package integrating_goroutines_channels

import (
	"fmt"
	"runtime"
	// "sync"
)

/*
	Mutex Lock with Goroutines and Channels
*/
func MutexLock()  {
	runtime.GOMAXPROCS(4)

	// mutex := new(sync.Mutex)
	mutex := make(chan bool, 1)

	for i:=1; i < 10; i++ {
		for j:=1; j < 10; j++ {
			// mutex.Lock()
			mutex<- true
			go func(){
				fmt.Printf("%d + %d = %d\n", i, j, i+j)
				// mutex.Unlock()
				<-mutex
			}()
		}
	}

	fmt.Scanln()
}