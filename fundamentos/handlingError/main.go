package main

import (
	"encoding/json"
	"fmt"

	"github.com/douglasandradeee/GoExercicies/fundamentos/handlingError/model"
)

func main() {

	house := model.Immobile{}
	house.Name = "Yellow House"
	house.X = 18
	house.Y = 25
	if err := house.SetValue(11000000); err != nil {
		fmt.Println("[main] there was an error in assigning value: ", err.Error())
		if err == model.ErrVeryHighPropertyValue {
			fmt.Println("Realtor, please check your rating")
		}
		return
	}

	fmt.Printf("House is: %+v\r\n", house)
	fmt.Printf("The value of the house is: %+v\r\n", house.GetValue())

	objJSON, err := json.Marshal(house)
	if err != nil {
		fmt.Println("[main] there was an error in generating the JSON object: ", err.Error())
		return
	}
	fmt.Println("The house in JSON: ", string(objJSON))
}
