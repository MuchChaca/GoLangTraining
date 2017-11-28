package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var unitSource string   // unit to convert from
var valueSource float64 // value to convert from
var unitDest string     // unit to convert to
var typeUnit string     // type of unit (temperature, weight, etc...)

func init() {
	// usageUnit lists all units available
	usageUnit := `
	Available units:
		temp
			f	Farenhit
			c	Celcius
			k	Kelvin
		mass
			Metric system
				t	Tonne
				kg	Kilogram
				hg	Hectogram
				dkg	Dekagram
				g	Gram
				dg	Decigram
				cg	Centigram
				mg	Milligram
			American & British units
				lt	Long ton
				st	Short ton
				lh	Long hundredweight
				sh	Short hundredweight
				pd	Pound
				oc	Ounce
				dr	Dram
				gr	Grain
		`
	// usageType lists all types abailable
	usageType := `
	Available types:
		temp	Temperatures
		mass	Mass
		dist	Distances
		currency	Currencies
		volume	Volumes`

	// Set up the command-lin flags
	// TODO handle help on a flag !?
	flag.StringVar(&unitSource, "us", "N/A", "The unit of the value to convert "+usageUnit)
	flag.Float64Var(&valueSource, "v", 0, "The value to convert.")
	flag.StringVar(&unitDest, "ud", "N/A", "The unit to convert to "+usageUnit)
	flag.StringVar(&typeUnit, "t", "N/A", "The type of unit to convert "+usageType)
	flag.Parse()

	if isHelp(os.Args) {
		//
	}

	conv(typeUnit, valueSource, unitSource, unitDest)
}

// conv makes the conversion
func conv(typeU string, valS float64, unitS string, unitD string) error {
	switch typeU {
	case "temp":
		// TODO
	case "weight":
		// TODO

	case "height":
		// TODO
	case "N/A":
		log.Fatalln("No type of unit provided!")
	default:
		log.Fatalln("An error as occured while processing the unit type.")
	}
	return nil
}

func main() {
	//? code here ?
}

// isHelp returns a true if a help flag is used ('-h' or '--help')
func isHelp(args []string) bool {
	for _, arg := range args {
		if strings.Compare(arg, "-h") == 0 || strings.Compare(arg, "--help") == 0 {
			return true
		}
	}
	return false
}

// isZero returns the name of the variable in a slice if the flag is not null
// and a slice of the flags
// if it's not comparable, it returns an error
func isZero(unitS string, valueS float64, unitD string, typeU string) ([]string, error) {
	/* t := reflect.TypeOf(v)
	if !t.Comparable() {
		return false, fmt.Errorf("type is not comparable: %v", t)
	}
	return v == reflect.Zero(t).Interface(), nil */
	//! BUGGED
	var listDefFlags = make([]string, 10)
	gotOne := false
	if unitS != "" {
		listDefFlags = append(listDefFlags, "unitSource")
		gotOne = true
	}
	if valueS != 0 {
		listDefFlags = append(listDefFlags, "valueSource")
		gotOne = true
	}
	if unitD != "" {
		listDefFlags = append(listDefFlags, "unitDestination")
		gotOne = true
	}
	if typeU != "" {
		listDefFlags = append(listDefFlags, "typeUnit")
		gotOne = true
	}

	if gotOne {
		return listDefFlags, nil
	}
	err := fmt.Errorf("No flag set")
	return nil, err
}
