package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// Dog is a struct
type Dog struct {
	DogName string `json:"name"`
	DogType string `json:"type"`
}

// Cat is a struct
type Cat struct {
	CatName string `json:"name"`
	CatType string `json:"type"`
}

// Hamster is a struct
type Hamster struct {
	HamsterName string `json:"name"`
	HamsterType string `json:"type"`
}

func main() {
	fmt.Println("Welcome to the echo server!")
	e := echo.New()

	e.GET("/", yallo)
	e.GET("/dogs/:data", getDogs)

	//* "Normal" way
	e.POST("/dogs", addDog)
	e.POST("/cats", addCat)

	//* "Echo" way
	e.POST("/hamsters", addHamster)

	e.Start(":8080")
}

func yallo(c echo.Context) error {
	return c.String(http.StatusOK, "yallo from the web side")
}

func getDogs(c echo.Context) error {
	// these would get query params such as:
	// localhost:8080?name=debby&type=terrier
	dogName := c.QueryParam("name")
	dogType := c.QueryParam("type")

	// this get data from url params such as:
	// localhost:8080/dogs/:data
	// where :data can be
	dataType := c.Param("data")

	/*
		//* The "if" way
		if dataType == "string"{
			return c.String(http.StatusOK, fmt.Sprintf("your dog name: %s\nand his type is: %s.", dogName, dogTYpe))
		}
		if dataType == "json"{
			return c.JSON(http.StatusOK, map[string]string{
				"name" : dogName,
				"type" : dogType
			})
		}

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error" : "Wrong data type"
		})
	*/

	//* The "switch" way
	switch dataType {
	case "json":
		return c.JSON(http.StatusOK, map[string]string{
			"name": dogName,
			"type": dogType,
		})
	case "string":
		return c.String(http.StatusOK, fmt.Sprintf("your dog name: %s\nand his type is: %s.", dogName, dogType))
	default:
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Wrong data type",
		})
	}

	//? localhost:8080/dogs/json?name=snoop&type=singer
	// or
	//? localhost:8080/dogs/string?name=snoop&type=singer
	// both work
}

func addDog(c echo.Context) error {
	dog := Dog{}

	defer c.Request().Body.Close()

	// ReadAll returns the body and an error
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body for addDog: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}
	//* Faster way
	err = json.Unmarshal(b, &dog)
	if err != nil {
		log.Printf("Failed to unmarchal in addDog: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("This is the dog: %#v", dog)
	return c.String(http.StatusOK, "We got your doggy!")
}

func addCat(c echo.Context) error {
	cat := Cat{}

	defer c.Request().Body.Close()
	//* 2nd faster way
	err := json.NewDecoder(c.Request().Body).Decode(&cat)
	if err != nil {
		log.Printf("Failed processing addCat request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("This is the cat: %#v", cat)
	return c.String(http.StatusOK, "We got your kitty!")
}

//* "Echo" Way
//* Note that this is slower
func addHamster(c echo.Context) error {
	hamster := Hamster{}

	err := c.Bind(&hamster)
	if err != nil {
		log.Printf("Failed processing addHamster request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("This is the hamster: %#v", hamster)
	return c.String(http.StatusOK, "We got your hamham!")
}
