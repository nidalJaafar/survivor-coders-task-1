package controller

import (
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
	"student/request"
	"student/service"
)

var v = validator.New()

func GetStudent(context echo.Context) error {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		return handleError(err)
	}
	if student, err := service.GetStudent(id); err != nil {
		return err
	} else {
		return context.JSON(http.StatusOK, student)
	}
}

func UpdateStudent(context echo.Context) error {
	var req request.StudentPutRequest

	if bindError := handleError(context.Bind(&req)); bindError != nil {
		return bindError
	}
	if validationError := handleError(v.Struct(req)); validationError != nil {
		return validationError
	}
	id, convertError := strconv.Atoi(context.Param("id"))
	if convertError != nil {
		return handleError(convertError)
	}
	if err := service.UpdateStudent(req, id); err != nil {
		return err
	}
	context.Response().WriteHeader(http.StatusCreated)
	return nil
}

func PatchStudent(context echo.Context) error {
	var req request.StudentPatchRequest

	if bindError := handleError(context.Bind(&req)); bindError != nil {
		return bindError
	}
	if validationError := handleError(v.Struct(req)); validationError != nil {
		return validationError
	}
	id, convertError := strconv.Atoi(context.Param("id"))
	if convertError != nil {
		return handleError(convertError)
	}
	if err := service.PatchStudent(req, id); err != nil {
		return err
	}
	context.Response().WriteHeader(http.StatusCreated)
	return nil
}

func GetAllStudents(context echo.Context) error {
	return handleError(context.JSON(http.StatusOK, service.GetAllStudents()))
}

func AddStudent(context echo.Context) error {
	var req request.StudentPostRequest
	err := handleError(context.Bind(&req))
	if err != nil {
		return err
	}
	err = handleError(v.Struct(req))
	if err != nil {
		return err
	}
	service.AddStudent(req)
	context.Response().WriteHeader(http.StatusCreated)
	return nil
}

func DeleteStudent(context echo.Context) error {
	id, _ := strconv.Atoi(context.Param("id"))
	if err := service.DeleteStudent(id); err != nil {
		return err
	}
	context.Response().WriteHeader(http.StatusNoContent)
	return nil
}

func handleError(err error) error {
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
