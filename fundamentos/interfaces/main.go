package main

import (
	"fmt"

	"github.com/douglasandradeee/GoExercicies/fundamentos/interfaces/model"
)

func main() {

	jojo := model.Bird{}
	jojo.Name = "Jojo da Silva"

	//wakeUpWithCackel(jojo)
	duckSound(jojo)

}

func wakeUpWithCackel(g model.Chicken) {
	fmt.Println(g.Cackle())
}

func duckSound(d model.Duck) {
	fmt.Println(d.Quack())
}
