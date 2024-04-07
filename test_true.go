package main

// Importing fmt and time
import (
    "fmt"
)

func main() {

    // Calling Tick method
    // using range keyword
	fileAccessInt := 6
    hasExec := (fileAccessInt & 0o111) != 0
	if hasExec {
		fmt.Println(hasExec)
	}

	fileAccessInt = 5
    hasExec = (fileAccessInt & 0o111) != 0
	if hasExec {
		fmt.Println(!hasExec)
	}
}

