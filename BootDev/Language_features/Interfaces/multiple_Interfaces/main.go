package main

import ("fmt")

func (e email) cost() int { 
    if !e.isSubscribed{
        return (len(e.body) * 5)
    }
    return (len(e.body) * 2)
}

func (e email) format() string {
    var content string = e.body
    status := "Subscribed"
    if !e.isSubscribed{
        status = "Not Subscribed"
    }
    return fmt.Sprintf("'%s' | %s",content, status)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}

