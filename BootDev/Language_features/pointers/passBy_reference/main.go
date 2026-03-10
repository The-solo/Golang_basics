package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

// ?

func analyzeMessage(msgPool *Analytics, msg Message) {

    if msg.Success {
        msgPool.MessagesSucceeded++ // dereference syntax can be *msgPool. 
    } else {
        msgPool.MessagesFailed++
    }
    msgPool.MessagesTotal = msgPool.MessagesSucceeded + msgPool.MessagesFailed
}

