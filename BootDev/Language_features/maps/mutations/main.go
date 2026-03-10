package main
import("errors")

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error){
	val, ok := users[name] //val is the local copy of that entire object with name
	if !ok {
		return false, errors.New("not found")
	}
	if val.scheduledForDeletion{ 
		delete(users, name)
		return true, nil
	}
	return false, nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}


