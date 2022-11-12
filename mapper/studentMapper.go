package mapper

import (
	"student/dto"
	"student/model"
	"student/request"
)

func UpdateStudent(student *model.Student, putRequest request.StudentPutRequest) {
	student.FirstName = putRequest.FirstName
	student.LastName = putRequest.LastName
}

func PatchStudent(student *model.Student, patchRequest request.StudentPatchRequest) {
	if patchRequest.FirstName != "" {
		student.FirstName = patchRequest.FirstName
	}
	if patchRequest.LastName != "" {
		student.LastName = patchRequest.LastName
	}
}

func StudentPostRequestToStudent(request request.StudentPostRequest) model.Student {
	return model.Student{FirstName: request.FirstName, LastName: request.LastName}
}

func StudentToStudentDto(student model.Student) dto.StudentDto {
	return dto.StudentDto{Id: student.ID, FirstName: student.FirstName, LastName: student.LastName, CreatedAt: student.CreatedAt, UpdatedAt: student.UpdatedAt}
}

func StudentsArrayToStudentDtoArray(students []model.Student) []dto.StudentDto {
	if len(students) == 0 {
		return []dto.StudentDto{}
	}
	var studentDtoArray []dto.StudentDto
	for _, student := range students {
		studentDtoArray = append(studentDtoArray, StudentToStudentDto(student))
	}
	return studentDtoArray
}
