package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Still running")
		time.Sleep(5 * time.Second)
	}
}
