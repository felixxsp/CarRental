package Entity

import (
	"fmt"
	"math/rand"
	"time"
)

type Electric struct {
	Battery  float64
	Motor    float64
	TopSpeed int
	CarMeter
}

func CreateElectric(year, speed int, battery, motor float64) Electric {
	return Electric{
		Battery:  battery,
		Motor:    motor,
		TopSpeed: speed,
		CarMeter: CarMeter{
			PlateID:    rand.Intn(8999) + 1000,
			Year:       year,
			Engine:     "Electric",
			KMs:        0,
			LastReturn: time.Now(),
			NextReturn: time.Now(),
		},
	}
}

func (car Electric) View() {
	fmt.Println("Plate ID: ", car.PlateID)
	fmt.Println("Year: ", car.Year)
	fmt.Println("Engine Type: ", car.Engine)
	fmt.Println("KMs: ", car.KMs)
	fmt.Printf("Last Returned on: 2023/10/%d-%d\n", car.LastReturn.Minute(), car.LastReturn.Second())
	fmt.Printf("Will Be Back on: 2023/10/%d-%d\n", car.NextReturn.Minute(), car.NextReturn.Second())
	fmt.Printf("Details===============\nBattery Capacity: %.2f MWH\n", car.Battery)
	fmt.Printf("Motor Power: %.2f KW\n", car.Motor)
	fmt.Printf("Top Speed: %.d Km/h\n", car.TopSpeed)
}

func (car Electric) RentOut(times int) {
	fmt.Println(times) //car.LastReturn.Add(time.Hour)
}
