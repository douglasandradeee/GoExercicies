package main

import (
	"fmt"

	"github.com/douglasandradeee/GoExercicies/hanoi"
)

func main() {

	fmt.Println("Torre de Hanoi : ")
	hanoi.Hanoi(0, 2, 1, 3)
}
