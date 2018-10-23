package main

import (
	user "os/user"
	"fmt"
	"myinterpreter/repl"
	"os"
)

func main() {
	us, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s!", us.Username)
	fmt.Println(" Type a command!")
	repl.Start(os.Stdin, os.Stdout)
}
