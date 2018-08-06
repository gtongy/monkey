package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/gtongy/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel gree to type in command\n")
	repl.Start(os.Stdin, os.Stdout)
}
