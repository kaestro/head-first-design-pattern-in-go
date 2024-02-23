package main

import (
	"commandPattern/badRemoteControl"
	"commandPattern/goodRemoteControl"
	"fmt"
)

func main() {
	fmt.Println("Bad RemoteControl")
	badRemoteControl.Simulate()

	fmt.Print("\n\n\n")

	fmt.Println("Good RemoteControl")
	goodRemoteControl.Simulate()
}
