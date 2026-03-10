package main
// This func can take multiple arguments
func sum(nums ...int) int { // nums is a slice of type int
	finalSum := 0
	for i:=0; i<len(nums); i++{
		finalSum += nums[i]
	}
	return finalSum
}
