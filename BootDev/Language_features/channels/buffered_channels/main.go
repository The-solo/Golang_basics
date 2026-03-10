package main

func addEmailsToQueue(emails []string) chan string {
	ch := make(chan string, len(emails)) //making a channel of buffer size emails[]
	for _, mail := range emails{
		ch <- mail // hading out the mails to channel ch.
	}
	close(ch) // make sure to close the channel.
	return ch //returning the channel with all the emails into it.
}
