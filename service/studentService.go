package service

import (
	"github.com/labstack/echo"
	"net/http"
	"student/config"
	"student/dto"
	"student/mapper"
	"student/model"
	"student/request"
)

func AddStudent(request request.StudentPostRequest) {
	student := mapper.StudentPostRequestToStudent(request)
	config.Db.Create(&student)
}

func GetAllStudents() []dto.StudentDto {
	var students []model.Student
	config.Db.Find(&students)
	return mapper.StudentsArrayToStudentDtoArray(students)
}

func GetStudent(id int) (dto.StudentDto, error) {
	var student model.Student
	err := config.Db.First(&student, id).Error
	if err != nil {
		return dto.StudentDto{}, echo.NewHTTPError(http.StatusNotFound, "Student not found")
	}
	return mapper.StudentToStudentDto(student), nil
}

func UpdateStudent(request request.StudentPutRequest, id int) error {
	var saved model.Student
	err := config.Db.First(&saved, id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Student not found")
	}
	mapper.UpdateStudent(&saved, request)
	config.Db.Save(&saved)
	return nil
}

func PatchStudent(request request.StudentPatchRequest, id int) error {
	var saved model.Student
	err := config.Db.First(&saved, id).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Student not found")
	}
	mapper.PatchStudent(&saved, request)
	config.Db.Save(&saved)
	return nil
}

func DeleteStudent(id int) error {
	result := config.Db.Delete(&model.Student{}, id)
	if result.RowsAffected == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Student not found")
	}
	return nil
}
