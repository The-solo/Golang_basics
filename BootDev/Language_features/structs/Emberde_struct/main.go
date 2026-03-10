package main

type sender struct {
    user //the user struct is now emberded into the sender struct.
	rateLimit int
}

type user struct {
	name   string
	number int
}

