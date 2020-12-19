package main

import "fmt"

func main() {
	months := 0
	situation := false
	city := "Salvador"

	if months > 0 && months <= 6 {
		fmt.Println("this creditor owes less than six months")
	}

	if situation {
		fmt.Println("He owes")
	}

	if !situation {
		fmt.Println("He compliant")
	}

	if city == "Salvador" {
		fmt.Println("The client lives on Bahia")
	}

	if descripition, status := debtTime(months); status {
		fmt.Println("What is the client's situation? ", descripition)
		return
	}

	fmt.Println("Thank you for consulting us.")
}

func debtTime(months int) (descripition string, status bool) {
	if months > 0 {
		status = true
		descripition = "He owes!"
		return
	}
	descripition = "He compliant!"
	return
}
