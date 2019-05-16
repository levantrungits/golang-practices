package channels

import (
	"fmt"
)

type Message struct {
	To []string
	From string
	Content string
}

type FailedMessage struct {
	ErrorMessage string
	OriginalMessage Message
}

/*
	Switching Between Channels
*/
func SwitchingBetweenChannels(){
	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	/*
	msg := Message{
		To: []string{"levantrungits@gmail.com"},
		From: "levantrungits@hotmail.com.vn",
		Content: "Keep it secret, keep it safe.",
	}

	failedMessage := FailedMessage{
		ErrorMessage: "Message intercepted by black rider",
		OriginalMessage: Message{},
	}

	msgCh<- msg
	errCh<- failedMessage
	*/

	select {
	case receivedMsg := <-msgCh:
		fmt.Println(receivedMsg)
	case receivedError := <-errCh:
		fmt.Println(receivedError)
	default:
		fmt.Println("ok ga den.")
	}
}