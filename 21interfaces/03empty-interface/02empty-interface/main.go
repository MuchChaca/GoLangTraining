package main

import "fmt"

// Empty interface
type vehicles interface{}

type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

// type tst struct {
// 	notEvenVehicule int
// }

type car struct {
	vehicle // embedds vehicle
	Wheels  int
	Doors   int
}

type plane struct {
	vehicle // embedds vehicle
	Jet     bool
}

type boat struct {
	vehicle // embedds vehicle
	Length  int
}

func (v vehicle) Specs() {
	fmt.Printf("Seats %v, max speed %v, color %v\n", v.Seats, v.MaxSpeed, v.Color)
}

func main() {
	prius := car{}
	tacoma := car{}
	bmw528 := car{}

	boeing747 := plane{}
	boeing757 := plane{}
	boeing767 := plane{}

	sanger := boat{}
	nautique := boat{}
	malibu := boat{}

	// smallTest := tst{}

	// rides := []vehicles{prius, tacoma, bmw528, boeing747, boeing757, boeing767, sanger, nautique, malibu, smallTest}

	// We can put anything in here
	rides := []vehicles{prius, tacoma, bmw528, boeing747, boeing757, boeing767, sanger, nautique, malibu}

	for key, value := range rides {
		fmt.Println(key, " - ", value)
	}

}
