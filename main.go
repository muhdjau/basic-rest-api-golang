package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Grade string `json:"grade"`
}

var students = []Student{
	{ID: 1, Name: "John Doe", Grade: "A"},
	{ID: 2, Name: "Jane Doe", Grade: "B"},
	{ID: 3, Name: "Steve Doe", Grade: "C"},
}

func main() {
	r := gin.Default()
	r.GET("/students", getStudents)
	r.GET("/students/:id", getStudent)
	r.POST("/students", createStudent)
	r.PUT("/students/:id", updateStudent)
	r.DELETE("/students/:id", deleteStudent)
	r.Run(":9090")
}

func getStudents(c *gin.Context) {
	c.JSON(http.StatusOK, students)
}

func getStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, student := range students {
		if student.ID == id {
			c.JSON(http.StatusOK, student)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}

func getNextID() int {
	lastStudent := students[len(students)-1]
	return lastStudent.ID + 1
}

func createStudent(c *gin.Context) {
	var student Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student.ID = getNextID()

	students = append(students, student)
	c.JSON(http.StatusCreated, student)
}

func updateStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedStudent Student
	if err := c.ShouldBindJSON(&updatedStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, student := range students {
		if student.ID == id {
			updatedStudent.ID = id
			students[i] = updatedStudent
			c.JSON(http.StatusOK, updatedStudent)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}

func deleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, student := range students {
		if student.ID == id {
			students = append(students[:i], students[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
}
