package main

import (
	"fmt"
	"os"
)

const (
	usage       = "Usage: [username] [password]"
	wrongpass   = "Invalid password for %q"
	userWrong   = "acces %q is wrong "
	acces       = "acces %q is accepted"
	user, user2 = "niko", "rio"
	pass, pass1 = "123", "456"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Printf(usage)
		return
	}

	u, p := args[1], args[2]

	if u != user && u != user2 {
		fmt.Printf(userWrong, u)
	} else if p != pass && p != pass1 {
		fmt.Printf(wrongpass, p)
	} else {
		fmt.Printf(acces, u)
	}
}
