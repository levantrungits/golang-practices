package channels

import (
	"fmt"
	"strings"
)

/*
	Ranging Over a Channel
*/
func RangingOverAChannelBasic(){
	phrase := "These are the times that try men's souls.\n"
	
	words := strings.Split(phrase, " ")

	// Create Channel
	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}

	close(ch)

	/*
	for {
		if msg, ok := <-ch; ok {
			fmt.Print(msg + " ")
		} else {
			break
		}
	} 
	*/

	for msg := range ch {
		fmt.Print(msg + " ")
	}
}