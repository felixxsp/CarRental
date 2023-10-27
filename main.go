package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Car interface {
	RentOut(days int)
	View()
}

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	//var year, speed int
	//var battery, motor float64
	//fmt.Scan(&year, &speed, &battery, &motor)
	//newCar := Entity.CreateElectric(year, speed, battery, motor)
	file, err := os.Open("database.json")
	if err != nil {
		panic(err)
	}
	temp, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var newCar Car
	json.Unmarshal(temp, &newCar)
	fmt.Println(newCar)
	Garage := make(map[int]Car)
	Garage[100] = newCar
	for _, car := range Garage {
		car.View()
		jsonEncoded, err := json.Marshal(car)
		if err != nil {
			panic(err)
		}

		if err := os.WriteFile("database.json", jsonEncoded, 0644); err != nil {
			panic(err)
		}
	}
}
