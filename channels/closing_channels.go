package channels

import (
	"fmt"
	"strings"
)

/*
	Closing Channels
*/
func ClosingChannelsBasic(){
	phrase := "These are the times that try men's souls.\n"
	
	words := strings.Split(phrase, " ")

	// Create Channel
	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}

	// Close Channel
	close(ch)

	for i:=0; i < len(words); i++ {
		fmt.Print(<-ch + " ")
	}

	ch <- "trunglv" // =>  panic: send on closed channel
}