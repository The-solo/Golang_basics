package main

func getLast[T any](s []T) T {
	var zero T //zero value for the current type & is valid for every type.
	if len(s) == 0 {
		return zero
	}
	return s[len(s)-1]	
}
