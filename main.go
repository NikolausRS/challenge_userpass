package main

import (
	"fmt"
	"os"
)

const (
	userInvalid = "Access denied for user %q\n"
	userPass    = "Invalid password for %q\n"
	succesLogin = "Acces granted to %q\n"
	uSage       = "Usage :[username] [password]"
	user        = "niko"
	pass        = "12345"
)

func main() {

	args := os.Args

	if len(os.Args) != 3 {
		fmt.Println(uSage)
		return
	}

	u, p := args[1], args[2]

	if u != user {
		fmt.Printf(userInvalid, u)
	} else if p != pass {
		fmt.Printf(userPass, p)
	} else {
		fmt.Printf(succesLogin, u)
	}

}
