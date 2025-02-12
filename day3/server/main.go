package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Flight struct {
	Id          int
	Number      string
	AirlineName string
	Source      string
	Destination string
	Capacity    int
	Price       float32
}

func readAllFlights(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello from readFlights")
}

func main() {
	//router
	r := gin.Default()
	//routes | API endpoints
	r.GET("/flights", readAllFlights)
	//server default port is 8080
	r.Run() //r.Run(":8081")

	/*flight1 := Flight{Id: 1001, Number: "AI 800", AirlineName: "Air India", Source: "Mumbai", Destination: "Banglore", Capacity: 180, Price: 15000.00}
	fmt.Println(flight1)*/

	// car1 := Car{Id: 101, Number: "TN48 0001", Model: "Benz", Type: "Sclass"}

	// var cars []Car = []Car{
	// 	{Id: 101, Number: "TN48 0001", Model: "Benz", Type: "Sclass"},
	// 	{Id: 103, Number: "KA18 001", Model: "Benz", Type: "Cclass"},
	// }
	// fmt.Println(cars)

	// var car2 *Car = &car1
	// fmt.Println(car1)
	// fmt.Println(car2.Model)
	// fmt.Println(car1.Type)
	// fmt.Println("Hellllllloooooooooooo")
	// var first int = 10
	// var second int = 20
	// var sum int = first + second
	// var salaries[] float32 = [] float32 {3.32,4.53}
	// fmt.Println(salaries)
	// fmt.Println(sum)
}
