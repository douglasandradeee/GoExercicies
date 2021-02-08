package main

import (
	"encoding/json"
	"fmt"

	"github.com/douglasandradeee/GoExercicies/fundamentos/goBuild/model"
)

/*
GOOS=windows GOARCH=386 go build -o nameofproject
*/

func main() {

	house := model.Immobile{}
	house.Name = "Lucy's House"
	house.X = 17
	house.Y = 29
	house.SetValue(100)

	fmt.Printf("House is: %+v\r\n", house)
	fmt.Printf("The value of the house is: %+v\r\n", house.GetValue())
	objJSON, _ := json.Marshal(house)

	fmt.Println("The house in JSON: ", string(objJSON))
}
