package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		time.Sleep(5 * time.Second)

		fmt.Println("Hello from outer space...")
	}
}
