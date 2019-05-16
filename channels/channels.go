package channels

import (
	"fmt"
)

/*	Channel
	1. Communication
	2. Isolation (Memory Isolation)
	3. Synchronization
*/
func ChannelsBasic(){
	// Basic Channels
	ch := make(chan string, 1)
	ch <- "Hello Word"
	fmt.Println(<-ch)
}