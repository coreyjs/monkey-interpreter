package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Monkey Language v 0.0.1;  Welcome %s\n", user.Username)
	fmt.Printf("REPL\n")

	repl.Start(os.Stdin, os.Stdout)
}
