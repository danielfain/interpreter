package main

import (
	"fmt"
	"interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s, welcome to my language.\n", user.Name)

	repl.Start(os.Stdin, os.Stdout)
}
