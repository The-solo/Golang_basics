package main
// In golang the memory is arranged in contiguous blocks.
// So here the "Field ordering" matters. 
// Arranging the struct with the similar data type together helps reduce the memory wasteage.
// The allocations order mismatches results in a lots of unused memeory because the golang adds some padding
// To make it up for the execution speed.
type contact struct {
	sendingLimit int32	
	age          int32
    userID       string
}

type perms struct {
	permissionLevel int
	canSend         bool
	canReceive      bool
	canManage       bool
}
