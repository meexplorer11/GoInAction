package main

import (
	"GoInAction/greeting"
	"fmt"
)

func init() {
	fmt.Println("Even before main method..")
}

func main() {
	fmt.Println("Main method...")

	greeting.Greet()
}
