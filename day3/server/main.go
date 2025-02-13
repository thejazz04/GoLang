package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Flight struct {
	Id          string  `json:"id"`
	Number      string  `json :"number"`
	AirlineName string  `json :"airline_name"`
	Source      string  `json: "source"`
	Destination string  `json:"destination"`
	Capacity    int     `json: "capacity"`
	Price       float32 `json: "price"`
}

func readFlightById(c *gin.Context) {
	id := c.Param("id")
	flights := Flight{Id: id, Number: "AI 800", AirlineName: "Air India", Source: "Mumbai", Destination: "Banglore", Capacity: 180, Price: 15000.00}
	c.JSON(http.StatusOK, flights)
}
func readAllFlights(c *gin.Context) {
	flights := []Flight{
		{Id: "1001", Number: "AI 800", AirlineName: "Air India", Source: "Mumbai", Destination: "Banglore", Capacity: 180, Price: 15000.00},
		{Id: "1001", Number: "AI 845", AirlineName: "Air India", Source: "Mumbai", Destination: "Banglore", Capacity: 100, Price: 18000.00},
	}
	c.JSON(http.StatusOK, flights)
}

func createFlight(c *gin.Context) {
	var jbodyFlight Flight
	err := c.BindJSON(&jbodyFlight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error." + err.Error()})
		return
	}
	createdFlight := Flight{Id: "1001", Number: "AI 800", AirlineName: "Air India", Source: "Mumbai", Destination: "Banglore", Capacity: 180, Price: 15000.00}
	c.JSON(http.StatusCreated, gin.H{"message": "Flight Created Successfully.", "flight": createdFlight})
}

func updateFlight(c *gin.Context) {
	id := c.Param("id")
	var jbodyFlight Flight
	err := c.BindJSON(&jbodyFlight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error." + err.Error()})
		return
	}
	updatedFlight := Flight{Id: id, Number: "AI 800", AirlineName: "Air India", Source: "Mumbai", Destination: "Banglore", Capacity: 180, Price: 15000.00}
	c.JSON(http.StatusOK, gin.H{"message": "Flight Updated Successfully.", "flight": updatedFlight})
}

func deleteFlight(c *gin.Context) {
	//id := c.Param("id")
	//deletedFlight := Flight{Id: id, Number: "AI 800", AirlineName: "Air India", Source: "Mumbai", Destination: "Banglore", Capacity: 180, Price: 15000.00}
	c.JSON(http.StatusOK, gin.H{"message": "Flight deleted Successfully."})
}

func main() {
	//router
	r := gin.Default()
	//routes | API endpoints
	r.GET("/flights", readAllFlights)
	r.GET("/flights/:id", readFlightById)
	r.POST("/flights", createFlight)
	r.PUT("/flights/:id", updateFlight)
	r.DELETE("/flights/:id", deleteFlight)
	//server default port is 8080
	r.Run(":8081") //r.Run(":8081")

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
