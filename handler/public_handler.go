package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/letanthang/nc_student/db"
)

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func TestDB(c echo.Context) error {
	return c.JSON(http.StatusOK, db.Test)
}

func GetAllStudents(c echo.Context) error {
	// var students []db.Student
	// inputJson := `[{"first_name":"Tam","last_name":"Nguyen","age":100,"email":"tamnguyen@gmail.com"},{"first_name": "Hieu", "last_name": "Nguyen", "age":200,"email":"hieunguyen@gmail.com"}]`
	// json.Unmarshal([]byte(inputJson), &students)

	students, err := db.GetAllStudent()
	if err != nil {
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, students)
}

func SearchStudentSimple(c echo.Context) error {

	var req db.StudentSearchRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	students, err := db.SearchStudentSimple(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, db.Error{Code: http.StatusBadRequest, Msg: "bad request"})
	}

	return c.JSON(http.StatusOK, students)
}
