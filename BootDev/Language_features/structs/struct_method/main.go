package main

// This is just the blueprint of the struct.
type authenticationInfo struct {
	username string
	password string
}

// This is the actual copy with values.
var a = authenticationInfo {
    username : "USERNAME",
    password : "PASSWORD",
}

//The method on struct.
func (a authenticationInfo) getBasicAuth() string {
    return  "Authorization: Basic " + a.username+":" +a.password
}


