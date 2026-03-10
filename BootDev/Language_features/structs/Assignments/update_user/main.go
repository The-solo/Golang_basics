package main

type User struct {
    MemberShip
	Name string    
}

type MemberShip struct {
    Type string
    MessageCharLimit int
}

func newUser(name string, membershipType string) User {
    membership := MemberShip{Type : membershipType}
   if membershipType == "premium" {
       membership.MessageCharLimit = 1000

   } else {
       membership.Type = "standard"
       membership.MessageCharLimit = 100
   }
   return User{MemberShip : membership, Name : name}
}

/*

my OG implementation.



func newUser(name string, membershipType string) User { 
    if membershipType == "premium" { 
        return User{ 
            Name: name,
            MemberShip : MemberShip { 
                Type : "premium", 
                MessageCharLimit : 1000, }, 
            } 
    } else { 
        return User{ 
            Name: name, 
            MemberShip : MemberShip{ 
                Type : "standard", 
                MessageCharLimit : 100, 
            }, 
        } 
    } 
}

*/

