package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Inisiasi Constructor
type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade string `json:"grade"`
}

// Inisiasi Database
var students []Student

func main() {
	// Inisiasi Echo Framework
	e := echo.New()

	// Routing
	e.GET("/students", getStudents)
	e.GET("/student/:id", getStudent)
	e.POST("/students", createStudent)
	e.PUT("/students/:id", updateStudent)
	e.DELETE("/student/:id", deleteStudent)

	// Inisiasi Port
	e.Start(":8080")
}

// Mengambil semua data yang ada dan mengembalikan HTTP Status: GET
func getStudents(c echo.Context) error {
	return c.JSON(http.StatusOK, students)
}

// Mengambil data student berdasarkan ID dan mengembalikan HTTP Status: GET
func getStudent(c echo.Context) error {
	// Mengambil informasi ID yang string dan mengkonversi menjadi integer
	id, _ := strconv.Atoi(c.Param("id"))
	for _, student := range students {
		if student.ID == id {
			return c.JSON(http.StatusOK, student)
		}
	}
	return c.JSON(http.StatusNotFound, "Student not found")
}

// Membuat data baru: POST
func createStudent(c echo.Context) error {
	// Inisiasi Student Baru
	student := new(Student)

	// Mengembalikan informasi student
	if err := c.Bind(student); err != nil {
		return err
	}

	// ID Baru dan menambahkan data kedalam database
	student.ID = len(students) + 1
	students = append(students, *student)

	return c.JSON(http.StatusCreated, student)
}

// Update data student: PUT
func updateStudent(c echo.Context) error {
	// Mengambil informasi ID yang string dan mengkonversi menjadi integer
	id, _ := strconv.Atoi(c.Param("id"))
	for i := range students {
		if students[i].ID == id {
			updatedStudent := new(Student)
			if err := c.Bind(updatedStudent); err != nil {
				return err
			}

			// Update student data
			students[i].Name = updatedStudent.Name
			students[i].Age = updatedStudent.Age
			students[i].Grade = updatedStudent.Grade

			return c.JSON(http.StatusOK, students[i])
		}
	}
	return c.JSON(http.StatusNotFound, "Student not found")
}

// Menghapus informasi data student: DELETE
func deleteStudent(c echo.Context) error {
	// Mengambil informasi ID yang string dan mengkonversi menjadi integer
	id, _ := strconv.Atoi(c.Param("id"))
	for i := range students {
		if students[i].ID == id {
			students = append(students[:i], students[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, "Student not found")
}
