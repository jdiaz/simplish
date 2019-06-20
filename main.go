package main

import (
	"fmt"
	"os"
	"os/user"
	"simplish/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf(`
 ____ ___ __  __ ____  _     ___ ____  _   _ 
/ ___|_ _|  \/  |  _ \| |   |_ _/ ___|| | | |
\___ \| || |\/| | |_) | |    | |\___ \| |_| |
 ___) | || |  | |  __/| |___ | | ___) |  _  |
|____/___|_|  |_|_|   |_____|___|____/|_| |_|

Welcome %s! This is the simplish programming language!`, user.Username)
	fmt.Printf("\nFeel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
