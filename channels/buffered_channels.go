package channels

import (
	"fmt"
	"strings"
)

/*
	Buffered Channels
*/
func BufferedChannelsBasic(){
	phrase := "These are the times that try men's souls.\n"

	words := strings.Split(phrase, " ")

	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}

	for i:=0; i < len(words); i++ {
		fmt.Print(<-ch + " ")
	}
}