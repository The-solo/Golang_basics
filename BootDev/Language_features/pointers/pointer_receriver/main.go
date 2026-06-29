package main

// The e of type email is the pointer receiver on method setMessage().
func (e *email) setMessage(newMessage string) {
	e.message = newMessage
// Ideally you want to derefernce the pointer but go does it for you here.

// *** you only do manual dereferncing when you want to override the entire object.
}

// don't edit below this line

type email struct {
	message     string
	fromAddress string
	toAddress   string
}
