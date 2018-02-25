//usr/bin/env go run $0 $@; exit $?

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "%v args are not enough\n", len(os.Args))
		os.Exit(11)
	} else {
		fmt.Printf("Thanks for %v\n", os.Args[1:])
	}
}
