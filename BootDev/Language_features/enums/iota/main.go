package main

type emailStatus int 

// Type safe set of constants.
const(
	//Default type.
	Unknown emailStatus = iota // good practice to initialize with Unknown / invalid state.
	EmailBounced// 0
	EmailInvalid // 1
	EmailDelivered // 2
	EmailOpened // 3
)

// This is not really an enum but is enum like
// The iota keyword is used to group together the related values.
// These so called types are actually the sequence of numbers
