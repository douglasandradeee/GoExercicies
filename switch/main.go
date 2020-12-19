package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	number := 3
	fmt.Print("The number ", number, " is written like this: ")
	switch number {
	case 1:
		fmt.Println("One.")
	case 2:
		fmt.Println("Two.")
	case 3:
		fmt.Println("Three.")
	}

	fmt.Println("Are you from the Unix family?")
	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Println("Yes, I am!")
	default:
		fmt.Println("Never mind that question")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's Weekend", time.Now().Weekday())
	default:
		fmt.Println("It's a business day,go to work!", time.Now().Weekday())
	}

	number = 101
	fmt.Println("Does this number fit in a digit?")
	switch {
	case number < 10:
		fmt.Println("Yes!")
	case number >= 10 && number < 100:
		fmt.Println("Serves two digits...?")
	case number >= 100:
		fmt.Println("No, this number is too big!!")
	}
}
