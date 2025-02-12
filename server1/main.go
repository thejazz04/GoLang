package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	Id          string
	Name        string
	Designation string
	Technology  string
	Phone       string
	Commission  float32
	Salary      float32
}

func readEmployeeById(c *gin.Context) {
	id := c.Param("id")
	Employees := Employee{Id: id, Name: "Rishika", Designation: "Manager", Technology: "CS", Phone: "9988154662", Commission: 99999.99, Salary: 100000.00}
	c.JSON(http.StatusOK, Employees)
}
func readAllEmployees(c *gin.Context) {
	Employees := []Employee{
		{Id: "001", Name: "Rishika", Designation: "Manager", Technology: "CS", Phone: "9988154662", Commission: 99999.99, Salary: 100000.00},
		{Id: "002", Name: "Poojitha", Designation: "Asst. Manager", Technology: "CS", Phone: "9988154663", Commission: 99999.00, Salary: 50000.00},
	}
	c.JSON(http.StatusOK, Employees)
}

func createEmployee(c *gin.Context) {
	var jbodyEmployee Employee
	err := c.BindJSON(&jbodyEmployee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error." + err.Error()})
		return
	}
	createdEmployee := Employee{Id: "002", Name: "Poojitha", Designation: "Asst. Manager", Technology: "CS", Phone: "9988154663", Commission: 99999.00, Salary: 50000.00}
	c.JSON(http.StatusCreated, gin.H{"message": "Employee Created Successfully.", "Employee": createdEmployee})
}

func updateEmployee(c *gin.Context) {
	id := c.Param("id")
	var jbodyEmployee Employee
	err := c.BindJSON(&jbodyEmployee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error." + err.Error()})
		return
	}
	updatedEmployee := Employee{Id: id, Name: "Poojitha", Designation: "Asst. Manager", Technology: "CS", Phone: "9988154663", Commission: 99999.00, Salary: 50000.00}
	c.JSON(http.StatusOK, gin.H{"message": "Employee Updated Successfully.", "Employee": updatedEmployee})
}

func deleteEmployee(c *gin.Context) {
	//id := c.Param("id")
	//deletedEmployee := Employee{Id: id, Number: "AI 800", AirlineName: "Air India", Source: "Mumbai", Destination: "Banglore", Capacity: 180, Price: 15000.00}
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted Successfully."})
}

func main() {
	//router
	r := gin.Default()
	//routes | API endpoints
	r.GET("/Employees", readAllEmployees)
	r.GET("/Employees/:id", readEmployeeById)
	r.POST("/Employees", createEmployee)
	r.PUT("/Employees/:id", updateEmployee)
	r.DELETE("/Employees/:id", deleteEmployee)
	//server default port is 8080
	r.Run(":8080") //r.Run(":8081")

}
