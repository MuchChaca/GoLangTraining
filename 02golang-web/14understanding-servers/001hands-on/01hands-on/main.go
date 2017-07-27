package main

import "fmt"

type person struct {
	fName, lName string
	age          int
}

type secretAgent struct {
	person
	codeName      string
	licenseToKill bool
}

func (p person) pSpeak(s string) {
	fmt.Println(s)
}

func (sa secretAgent) saSpeak(s string) {
	fmt.Println(s)
}

func main() {
	pers1 := person{"Sarah", "Walker", 26}

	sAgent := secretAgent{person{"Leon", "Wahington", 24}, "Leon", true}

	fmt.Println("Field from person:", pers1.fName)
	pers1.pSpeak("I am a person tested.")

	fmt.Println("Field from secretAgent:", sAgent.codeName)
	sAgent.saSpeak("I'm defenitly not an agent ...")
	sAgent.pSpeak("I am a tested agent. lol.")
}
