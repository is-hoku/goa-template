// Code generated by goa v3.7.6, DO NOT EDIT.
//
// students HTTP server types
//
// Command:
// $ goa gen github.com/is-hoku/goa-template/webapi/design

package server

import (
	students "github.com/is-hoku/goa-template/webapi/gen/students"
	studentsviews "github.com/is-hoku/goa-template/webapi/gen/students/views"
)

// GetStudentResponseBody is the type of the "students" service "get_student"
// endpoint HTTP response body.
type GetStudentResponseBody struct {
	// 学生を一意に表す ID
	ID int64 `form:"id" json:"id" xml:"id"`
	// 学生の氏名
	Name string `form:"name" json:"name" xml:"name"`
	// 学生の氏名のフリガナ
	Ruby string `form:"ruby" json:"ruby" xml:"ruby"`
	// 学生の学籍番号
	StudentNumber int `form:"student_number" json:"student_number" xml:"student_number"`
	// 学生の生年月日 (RFC3339)
	DateOfBirth string `form:"date_of_birth" json:"date_of_birth" xml:"date_of_birth"`
	// 学生の住所
	Address string `form:"address" json:"address" xml:"address"`
	// 学生証の有効期間 (RFC3339)
	ExpirationDate string `form:"expiration_date" json:"expiration_date" xml:"expiration_date"`
}

// GetStudentsResponseBody is the type of the "students" service "get_students"
// endpoint HTTP response body.
type GetStudentsResponseBody struct {
	Students []*StudentResponseBody `form:"students" json:"students" xml:"students"`
}

// CreateStudentResponseBody is the type of the "students" service
// "create_student" endpoint HTTP response body.
type CreateStudentResponseBody struct {
	// 学生を一意に表す ID
	ID int64 `form:"id" json:"id" xml:"id"`
	// 学生の氏名
	Name string `form:"name" json:"name" xml:"name"`
	// 学生の氏名のフリガナ
	Ruby string `form:"ruby" json:"ruby" xml:"ruby"`
	// 学生の学籍番号
	StudentNumber int `form:"student_number" json:"student_number" xml:"student_number"`
	// 学生の生年月日 (RFC3339)
	DateOfBirth string `form:"date_of_birth" json:"date_of_birth" xml:"date_of_birth"`
	// 学生の住所
	Address string `form:"address" json:"address" xml:"address"`
	// 学生証の有効期間 (RFC3339)
	ExpirationDate string `form:"expiration_date" json:"expiration_date" xml:"expiration_date"`
}

// GetStudentInternalErrorResponseBody is the type of the "students" service
// "get_student" endpoint HTTP response body for the "internal_error" error.
type GetStudentInternalErrorResponseBody struct {
	// Name of error
	Name string `form:"name" json:"name" xml:"name"`
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
}

// GetStudentNotFoundResponseBody is the type of the "students" service
// "get_student" endpoint HTTP response body for the "not_found" error.
type GetStudentNotFoundResponseBody struct {
	// Name of error
	Name string `form:"name" json:"name" xml:"name"`
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
}

// GetStudentsInternalErrorResponseBody is the type of the "students" service
// "get_students" endpoint HTTP response body for the "internal_error" error.
type GetStudentsInternalErrorResponseBody struct {
	// Name of error
	Name string `form:"name" json:"name" xml:"name"`
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
}

// CreateStudentInternalErrorResponseBody is the type of the "students" service
// "create_student" endpoint HTTP response body for the "internal_error" error.
type CreateStudentInternalErrorResponseBody struct {
	// Name of error
	Name string `form:"name" json:"name" xml:"name"`
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
}

// CreateStudentBadRequestResponseBody is the type of the "students" service
// "create_student" endpoint HTTP response body for the "bad_request" error.
type CreateStudentBadRequestResponseBody struct {
	// Name of error
	Name string `form:"name" json:"name" xml:"name"`
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
}

// StudentResponseBody is used to define fields on response body types.
type StudentResponseBody struct {
	// 学生を一意に表す ID
	ID int64 `form:"id" json:"id" xml:"id"`
	// 学生の氏名
	Name string `form:"name" json:"name" xml:"name"`
	// 学生の氏名のフリガナ
	Ruby string `form:"ruby" json:"ruby" xml:"ruby"`
	// 学生の学籍番号
	StudentNumber int `form:"student_number" json:"student_number" xml:"student_number"`
	// 学生の生年月日 (RFC3339)
	DateOfBirth string `form:"date_of_birth" json:"date_of_birth" xml:"date_of_birth"`
	// 学生の住所
	Address string `form:"address" json:"address" xml:"address"`
	// 学生証の有効期間 (RFC3339)
	ExpirationDate string `form:"expiration_date" json:"expiration_date" xml:"expiration_date"`
}

// NewGetStudentResponseBody builds the HTTP response body from the result of
// the "get_student" endpoint of the "students" service.
func NewGetStudentResponseBody(res *studentsviews.StudentView) *GetStudentResponseBody {
	body := &GetStudentResponseBody{
		ID:             *res.ID,
		Name:           *res.Name,
		Ruby:           *res.Ruby,
		StudentNumber:  *res.StudentNumber,
		DateOfBirth:    *res.DateOfBirth,
		Address:        *res.Address,
		ExpirationDate: *res.ExpirationDate,
	}
	return body
}

// NewGetStudentsResponseBody builds the HTTP response body from the result of
// the "get_students" endpoint of the "students" service.
func NewGetStudentsResponseBody(res *studentsviews.StudentsView) *GetStudentsResponseBody {
	body := &GetStudentsResponseBody{}
	if res.Students != nil {
		body.Students = make([]*StudentResponseBody, len(res.Students))
		for i, val := range res.Students {
			body.Students[i] = marshalStudentsviewsStudentViewToStudentResponseBody(val)
		}
	}
	return body
}

// NewCreateStudentResponseBody builds the HTTP response body from the result
// of the "create_student" endpoint of the "students" service.
func NewCreateStudentResponseBody(res *studentsviews.StudentView) *CreateStudentResponseBody {
	body := &CreateStudentResponseBody{
		ID:             *res.ID,
		Name:           *res.Name,
		Ruby:           *res.Ruby,
		StudentNumber:  *res.StudentNumber,
		DateOfBirth:    *res.DateOfBirth,
		Address:        *res.Address,
		ExpirationDate: *res.ExpirationDate,
	}
	return body
}

// NewGetStudentInternalErrorResponseBody builds the HTTP response body from
// the result of the "get_student" endpoint of the "students" service.
func NewGetStudentInternalErrorResponseBody(res *students.CustomError) *GetStudentInternalErrorResponseBody {
	body := &GetStudentInternalErrorResponseBody{
		Name:    res.Name,
		Message: res.Message,
	}
	return body
}

// NewGetStudentNotFoundResponseBody builds the HTTP response body from the
// result of the "get_student" endpoint of the "students" service.
func NewGetStudentNotFoundResponseBody(res *students.CustomError) *GetStudentNotFoundResponseBody {
	body := &GetStudentNotFoundResponseBody{
		Name:    res.Name,
		Message: res.Message,
	}
	return body
}

// NewGetStudentsInternalErrorResponseBody builds the HTTP response body from
// the result of the "get_students" endpoint of the "students" service.
func NewGetStudentsInternalErrorResponseBody(res *students.CustomError) *GetStudentsInternalErrorResponseBody {
	body := &GetStudentsInternalErrorResponseBody{
		Name:    res.Name,
		Message: res.Message,
	}
	return body
}

// NewCreateStudentInternalErrorResponseBody builds the HTTP response body from
// the result of the "create_student" endpoint of the "students" service.
func NewCreateStudentInternalErrorResponseBody(res *students.CustomError) *CreateStudentInternalErrorResponseBody {
	body := &CreateStudentInternalErrorResponseBody{
		Name:    res.Name,
		Message: res.Message,
	}
	return body
}

// NewCreateStudentBadRequestResponseBody builds the HTTP response body from
// the result of the "create_student" endpoint of the "students" service.
func NewCreateStudentBadRequestResponseBody(res *students.CustomError) *CreateStudentBadRequestResponseBody {
	body := &CreateStudentBadRequestResponseBody{
		Name:    res.Name,
		Message: res.Message,
	}
	return body
}

// NewGetStudentPayload builds a students service get_student endpoint payload.
func NewGetStudentPayload(id int64) *students.GetStudentPayload {
	v := &students.GetStudentPayload{}
	v.ID = &id

	return v
}
