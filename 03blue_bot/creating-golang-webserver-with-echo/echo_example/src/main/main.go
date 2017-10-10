package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

// JWTClaims is a struct to set the claims of the JWT
type JWTClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func main() {
	fmt.Println("Welcome to the echo server!")
	e := echo.New()

	e.Use(ServerHeader)

	//* A new Group
	// We give it a prefix "/admin"
	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")
	jwtGroup := e.Group("/jwt")

	// e.Use(middleware.Static("../../static"))
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{}))

	// A way to add a middleware:
	adminGroup.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}]  ${status}  ${method}  ${host}${path}  ${latency_human}` + "\n",
	}))

	//* Basic Authentification middleware
	// in basic auth, the username:password is encoded in the header in base 64
	// can be decoded at https://www.base64decode.org/
	adminGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "admin" && password == "admin" {
			return true, nil
		}
		return false, nil
	}))

	jwtGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte("mySecret"),
	}))

	cookieGroup.Use(checkCookie)

	cookieGroup.GET("/main", mainCookie)

	adminGroup.GET("/main", mainAdmin)

	jwtGroup.GET("/main", mainJWT)

	e.GET("/login", login)

	e.GET("/yallo", yallo)
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

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to the Admin main page")
}

func mainCookie(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to the secret cookie page mained")
}

func mainJWT(c echo.Context) error {
	user := c.Get("user")      // this returns an interface
	token := user.(*jwt.Token) // becomes of type *jwt.Token

	claims := token.Claims.(jwt.MapClaims) // Now we have the claims

	log.Println("User Name: ", claims["name"], "User ID", claims["jti"])

	return c.String(http.StatusOK, "You are on the secret JWT page!")
}

func login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	// check in the db after hashing but here juste for the test:
	if username == "test" && password == "test" {
		cookie := &http.Cookie{}
		// this is the same:
		// cookie := new(http.Cookie)

		cookie.Name = "sessionID"
		cookie.Value = "some_string" // a hash ? to represent the session of this user
		cookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(cookie)

		//// Create JWT token
		token, err := createJWT()
		if err != nil {
			log.Println("Error creating JWT token? NANII!!??")
			return c.String(http.StatusInternalServerError, "something went wrong")
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "You are logged in!",
			"token":   token,
		})
	}

	return c.String(http.StatusUnauthorized, "Your username or password are incorrect!")
}

func createJWT() (string, error) {
	claims := JWTClaims{
		"Jack",
		jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString([]byte("mySecret")) // do something more secured if for real
	if err != nil {
		return "", err
	}

	return token, nil
}

//************  MIDDLEWARES *************//

// ServerHeader is a middleware that will add to any response, the server name
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "YellowBot/1.0")

		return next(c)
	}
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")
		if err != nil {
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "Need cookies to join the dark side!")
			}
			log.Println(err)
			return err
		}

		if cookie.Value == "some_string" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "Need cookies to join the dark side!")
	}
}
