package main

import (
	"encoding/json"
	"fmt"

	"github.com/douglasandradeee/GoExercicies/advancedStructs/model"
)

func main() {

	house := model.Immobile{}
	house.Name = "Yellow House"
	house.X = 18
	house.Y = 25
	house.SetValue(60000)

	fmt.Printf("House is: %+v\r\n", house)
	fmt.Printf("The value of the house is: %+v\r\n", house.GetValue())
	objJSON, _ := json.Marshal(house)

	fmt.Println("The house in JSON: ", string(objJSON))
}
