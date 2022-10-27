// Code generated by goa v3.7.6, DO NOT EDIT.
//
// student HTTP server types
//
// Command:
// $ goa gen github.com/is-hoku/goa-sample/webapi/design

package server

import (
	student "github.com/is-hoku/goa-sample/webapi/gen/student"
	studentviews "github.com/is-hoku/goa-sample/webapi/gen/student/views"
	goa "goa.design/goa/v3/pkg"
)

// CreateStudentRequestBody is the type of the "student" service
// "create_student" endpoint HTTP request body.
type CreateStudentRequestBody struct {
	// 学生の氏名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// 学生の氏名のフリガナ
	Ruby *string `form:"ruby,omitempty" json:"ruby,omitempty" xml:"ruby,omitempty"`
	// 学生の学籍番号
	StudentNumber *uint32 `form:"student_number,omitempty" json:"student_number,omitempty" xml:"student_number,omitempty"`
	// 学生の生年月日 (RFC3339)
	DateOfBirth *string `form:"date_of_birth,omitempty" json:"date_of_birth,omitempty" xml:"date_of_birth,omitempty"`
	// 学生の住所
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// 学生証の有効期間 (RFC3339)
	ExpirationDate *string `form:"expiration_date,omitempty" json:"expiration_date,omitempty" xml:"expiration_date,omitempty"`
}

// UpdateStudentRequestBody is the type of the "student" service
// "update_student" endpoint HTTP request body.
type UpdateStudentRequestBody struct {
	// 学生の氏名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// 学生の氏名のフリガナ
	Ruby *string `form:"ruby,omitempty" json:"ruby,omitempty" xml:"ruby,omitempty"`
	// 学生の学籍番号
	StudentNumber *uint32 `form:"student_number,omitempty" json:"student_number,omitempty" xml:"student_number,omitempty"`
	// 学生の生年月日 (RFC3339)
	DateOfBirth *string `form:"date_of_birth,omitempty" json:"date_of_birth,omitempty" xml:"date_of_birth,omitempty"`
	// 学生の住所
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// 学生証の有効期間 (RFC3339)
	ExpirationDate *string `form:"expiration_date,omitempty" json:"expiration_date,omitempty" xml:"expiration_date,omitempty"`
}

// GetStudentResponseBody is the type of the "student" service "get_student"
// endpoint HTTP response body.
type GetStudentResponseBody struct {
	// 学生を一意に表す ID
	ID uint64 `form:"id" json:"id" xml:"id"`
	// 学生の氏名
	Name string `form:"name" json:"name" xml:"name"`
	// 学生の氏名のフリガナ
	Ruby string `form:"ruby" json:"ruby" xml:"ruby"`
	// 学生の学籍番号
	StudentNumber uint32 `form:"student_number" json:"student_number" xml:"student_number"`
	// 学生の生年月日 (RFC3339)
	DateOfBirth string `form:"date_of_birth" json:"date_of_birth" xml:"date_of_birth"`
	// 学生の住所
	Address string `form:"address" json:"address" xml:"address"`
	// 学生証の有効期間 (RFC3339)
	ExpirationDate string `form:"expiration_date" json:"expiration_date" xml:"expiration_date"`
}

// GetStudentsResponseBody is the type of the "student" service "get_students"
// endpoint HTTP response body.
type GetStudentsResponseBody struct {
	Students []*StudentResponseBody `form:"students" json:"students" xml:"students"`
}

// CreateStudentResponseBody is the type of the "student" service
// "create_student" endpoint HTTP response body.
type CreateStudentResponseBody struct {
	// 学生を一意に表す ID
	ID uint64 `form:"id" json:"id" xml:"id"`
	// 学生の氏名
	Name string `form:"name" json:"name" xml:"name"`
	// 学生の氏名のフリガナ
	Ruby string `form:"ruby" json:"ruby" xml:"ruby"`
	// 学生の学籍番号
	StudentNumber uint32 `form:"student_number" json:"student_number" xml:"student_number"`
	// 学生の生年月日 (RFC3339)
	DateOfBirth string `form:"date_of_birth" json:"date_of_birth" xml:"date_of_birth"`
	// 学生の住所
	Address string `form:"address" json:"address" xml:"address"`
	// 学生証の有効期間 (RFC3339)
	ExpirationDate string `form:"expiration_date" json:"expiration_date" xml:"expiration_date"`
}

// UpdateStudentResponseBody is the type of the "student" service
// "update_student" endpoint HTTP response body.
type UpdateStudentResponseBody struct {
	// 学生を一意に表す ID
	ID uint64 `form:"id" json:"id" xml:"id"`
	// 学生の氏名
	Name string `form:"name" json:"name" xml:"name"`
	// 学生の氏名のフリガナ
	Ruby string `form:"ruby" json:"ruby" xml:"ruby"`
	// 学生の学籍番号
	StudentNumber uint32 `form:"student_number" json:"student_number" xml:"student_number"`
	// 学生の生年月日 (RFC3339)
	DateOfBirth string `form:"date_of_birth" json:"date_of_birth" xml:"date_of_birth"`
	// 学生の住所
	Address string `form:"address" json:"address" xml:"address"`
	// 学生証の有効期間 (RFC3339)
	ExpirationDate string `form:"expiration_date" json:"expiration_date" xml:"expiration_date"`
}

// GetStudentInternalErrorResponseBody is the type of the "student" service
// "get_student" endpoint HTTP response body for the "internal_error" error.
type GetStudentInternalErrorResponseBody struct {
	// Name of error
	Name string `form:"name" json:"name" xml:"name"`
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
}

// GetStudentNotFoundResponseBody is the type of the "student" service
// "get_student" endpoint HTTP response body for the "not_found" error.
type GetStudentNotFoundResponseBody struct {
	// Name of error
	Name string `form:"name" json:"name" xml:"name"`
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
}

// GetStudentBadRequestResponseBody is the type of the "student" service
// "get_student" endpoint HTTP response body for the "bad_request" error.
type GetStudentBadRequestResponseBody struct {
	// Name of error
	Name string `form:"name" json:"name" xml:"name"`
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
}

// GetStudentsInternalErrorResponseBody is the type of the "student" service
// "get_students" endpoint HTTP response body for the "internal_error" error.
type GetStudentsInternalErrorResponseBody struct {
	// Name of error
	Name string `form:"name" json:"name" xml:"name"`
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
}

// CreateStudentInternalErrorResponseBody is the type of the "student" service
// "create_student" endpoint HTTP response body for the "internal_error" error.
type CreateStudentInternalErrorResponseBody struct {
	// Name of error
	Name string `form:"name" json:"name" xml:"name"`
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
}

// CreateStudentBadRequestResponseBody is the type of the "student" service
// "create_student" endpoint HTTP response body for the "bad_request" error.
type CreateStudentBadRequestResponseBody struct {
	// Name of error
	Name string `form:"name" json:"name" xml:"name"`
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
}

// UpdateStudentInternalErrorResponseBody is the type of the "student" service
// "update_student" endpoint HTTP response body for the "internal_error" error.
type UpdateStudentInternalErrorResponseBody struct {
	// Name of error
	Name string `form:"name" json:"name" xml:"name"`
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
}

// UpdateStudentBadRequestResponseBody is the type of the "student" service
// "update_student" endpoint HTTP response body for the "bad_request" error.
type UpdateStudentBadRequestResponseBody struct {
	// Name of error
	Name string `form:"name" json:"name" xml:"name"`
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
}

// DeleteStudentInternalErrorResponseBody is the type of the "student" service
// "delete_student" endpoint HTTP response body for the "internal_error" error.
type DeleteStudentInternalErrorResponseBody struct {
	// Name of error
	Name string `form:"name" json:"name" xml:"name"`
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
}

// DeleteStudentNotFoundResponseBody is the type of the "student" service
// "delete_student" endpoint HTTP response body for the "not_found" error.
type DeleteStudentNotFoundResponseBody struct {
	// Name of error
	Name string `form:"name" json:"name" xml:"name"`
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
}

// DeleteStudentBadRequestResponseBody is the type of the "student" service
// "delete_student" endpoint HTTP response body for the "bad_request" error.
type DeleteStudentBadRequestResponseBody struct {
	// Name of error
	Name string `form:"name" json:"name" xml:"name"`
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
}

// StudentResponseBody is used to define fields on response body types.
type StudentResponseBody struct {
	// 学生を一意に表す ID
	ID uint64 `form:"id" json:"id" xml:"id"`
	// 学生の氏名
	Name string `form:"name" json:"name" xml:"name"`
	// 学生の氏名のフリガナ
	Ruby string `form:"ruby" json:"ruby" xml:"ruby"`
	// 学生の学籍番号
	StudentNumber uint32 `form:"student_number" json:"student_number" xml:"student_number"`
	// 学生の生年月日 (RFC3339)
	DateOfBirth string `form:"date_of_birth" json:"date_of_birth" xml:"date_of_birth"`
	// 学生の住所
	Address string `form:"address" json:"address" xml:"address"`
	// 学生証の有効期間 (RFC3339)
	ExpirationDate string `form:"expiration_date" json:"expiration_date" xml:"expiration_date"`
}

// NewGetStudentResponseBody builds the HTTP response body from the result of
// the "get_student" endpoint of the "student" service.
func NewGetStudentResponseBody(res *studentviews.StudentView) *GetStudentResponseBody {
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
// the "get_students" endpoint of the "student" service.
func NewGetStudentsResponseBody(res *studentviews.StudentsView) *GetStudentsResponseBody {
	body := &GetStudentsResponseBody{}
	if res.Students != nil {
		body.Students = make([]*StudentResponseBody, len(res.Students))
		for i, val := range res.Students {
			body.Students[i] = marshalStudentviewsStudentViewToStudentResponseBody(val)
		}
	}
	return body
}

// NewCreateStudentResponseBody builds the HTTP response body from the result
// of the "create_student" endpoint of the "student" service.
func NewCreateStudentResponseBody(res *studentviews.StudentView) *CreateStudentResponseBody {
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

// NewUpdateStudentResponseBody builds the HTTP response body from the result
// of the "update_student" endpoint of the "student" service.
func NewUpdateStudentResponseBody(res *studentviews.StudentView) *UpdateStudentResponseBody {
	body := &UpdateStudentResponseBody{
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
// the result of the "get_student" endpoint of the "student" service.
func NewGetStudentInternalErrorResponseBody(res *student.CustomError) *GetStudentInternalErrorResponseBody {
	body := &GetStudentInternalErrorResponseBody{
		Name:    res.Name,
		Message: res.Message,
	}
	return body
}

// NewGetStudentNotFoundResponseBody builds the HTTP response body from the
// result of the "get_student" endpoint of the "student" service.
func NewGetStudentNotFoundResponseBody(res *student.CustomError) *GetStudentNotFoundResponseBody {
	body := &GetStudentNotFoundResponseBody{
		Name:    res.Name,
		Message: res.Message,
	}
	return body
}

// NewGetStudentBadRequestResponseBody builds the HTTP response body from the
// result of the "get_student" endpoint of the "student" service.
func NewGetStudentBadRequestResponseBody(res *student.CustomError) *GetStudentBadRequestResponseBody {
	body := &GetStudentBadRequestResponseBody{
		Name:    res.Name,
		Message: res.Message,
	}
	return body
}

// NewGetStudentsInternalErrorResponseBody builds the HTTP response body from
// the result of the "get_students" endpoint of the "student" service.
func NewGetStudentsInternalErrorResponseBody(res *student.CustomError) *GetStudentsInternalErrorResponseBody {
	body := &GetStudentsInternalErrorResponseBody{
		Name:    res.Name,
		Message: res.Message,
	}
	return body
}

// NewCreateStudentInternalErrorResponseBody builds the HTTP response body from
// the result of the "create_student" endpoint of the "student" service.
func NewCreateStudentInternalErrorResponseBody(res *student.CustomError) *CreateStudentInternalErrorResponseBody {
	body := &CreateStudentInternalErrorResponseBody{
		Name:    res.Name,
		Message: res.Message,
	}
	return body
}

// NewCreateStudentBadRequestResponseBody builds the HTTP response body from
// the result of the "create_student" endpoint of the "student" service.
func NewCreateStudentBadRequestResponseBody(res *student.CustomError) *CreateStudentBadRequestResponseBody {
	body := &CreateStudentBadRequestResponseBody{
		Name:    res.Name,
		Message: res.Message,
	}
	return body
}

// NewUpdateStudentInternalErrorResponseBody builds the HTTP response body from
// the result of the "update_student" endpoint of the "student" service.
func NewUpdateStudentInternalErrorResponseBody(res *student.CustomError) *UpdateStudentInternalErrorResponseBody {
	body := &UpdateStudentInternalErrorResponseBody{
		Name:    res.Name,
		Message: res.Message,
	}
	return body
}

// NewUpdateStudentBadRequestResponseBody builds the HTTP response body from
// the result of the "update_student" endpoint of the "student" service.
func NewUpdateStudentBadRequestResponseBody(res *student.CustomError) *UpdateStudentBadRequestResponseBody {
	body := &UpdateStudentBadRequestResponseBody{
		Name:    res.Name,
		Message: res.Message,
	}
	return body
}

// NewDeleteStudentInternalErrorResponseBody builds the HTTP response body from
// the result of the "delete_student" endpoint of the "student" service.
func NewDeleteStudentInternalErrorResponseBody(res *student.CustomError) *DeleteStudentInternalErrorResponseBody {
	body := &DeleteStudentInternalErrorResponseBody{
		Name:    res.Name,
		Message: res.Message,
	}
	return body
}

// NewDeleteStudentNotFoundResponseBody builds the HTTP response body from the
// result of the "delete_student" endpoint of the "student" service.
func NewDeleteStudentNotFoundResponseBody(res *student.CustomError) *DeleteStudentNotFoundResponseBody {
	body := &DeleteStudentNotFoundResponseBody{
		Name:    res.Name,
		Message: res.Message,
	}
	return body
}

// NewDeleteStudentBadRequestResponseBody builds the HTTP response body from
// the result of the "delete_student" endpoint of the "student" service.
func NewDeleteStudentBadRequestResponseBody(res *student.CustomError) *DeleteStudentBadRequestResponseBody {
	body := &DeleteStudentBadRequestResponseBody{
		Name:    res.Name,
		Message: res.Message,
	}
	return body
}

// NewGetStudentPayload builds a student service get_student endpoint payload.
func NewGetStudentPayload(studentNumber uint32, authorization string) *student.GetStudentPayload {
	v := &student.GetStudentPayload{}
	v.StudentNumber = &studentNumber
	v.Authorization = authorization

	return v
}

// NewGetStudentsPayload builds a student service get_students endpoint payload.
func NewGetStudentsPayload(authorization string) *student.GetStudentsPayload {
	v := &student.GetStudentsPayload{}
	v.Authorization = authorization

	return v
}

// NewCreateStudentPayload builds a student service create_student endpoint
// payload.
func NewCreateStudentPayload(body *CreateStudentRequestBody, authorization string) *student.CreateStudentPayload {
	v := &student.CreateStudentPayload{
		Name:           *body.Name,
		Ruby:           *body.Ruby,
		StudentNumber:  *body.StudentNumber,
		DateOfBirth:    *body.DateOfBirth,
		Address:        *body.Address,
		ExpirationDate: *body.ExpirationDate,
	}
	v.Authorization = authorization

	return v
}

// NewUpdateStudentPayload builds a student service update_student endpoint
// payload.
func NewUpdateStudentPayload(body *UpdateStudentRequestBody, studentNumber uint32, authorization string) *student.UpdateStudentPayload {
	v := &student.UpdateStudentPayload{
		Name:           *body.Name,
		Ruby:           *body.Ruby,
		StudentNumber:  *body.StudentNumber,
		DateOfBirth:    *body.DateOfBirth,
		Address:        *body.Address,
		ExpirationDate: *body.ExpirationDate,
	}
	v.StudentNumber = studentNumber
	v.Authorization = authorization

	return v
}

// NewDeleteStudentPayload builds a student service delete_student endpoint
// payload.
func NewDeleteStudentPayload(studentNumber uint32, authorization string) *student.DeleteStudentPayload {
	v := &student.DeleteStudentPayload{}
	v.StudentNumber = &studentNumber
	v.Authorization = authorization

	return v
}

// ValidateCreateStudentRequestBody runs the validations defined on
// create_student_request_body
func ValidateCreateStudentRequestBody(body *CreateStudentRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Ruby == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ruby", "body"))
	}
	if body.StudentNumber == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("student_number", "body"))
	}
	if body.DateOfBirth == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("date_of_birth", "body"))
	}
	if body.Address == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("address", "body"))
	}
	if body.ExpirationDate == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("expiration_date", "body"))
	}
	if body.DateOfBirth != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.date_of_birth", *body.DateOfBirth, goa.FormatDateTime))
	}
	if body.ExpirationDate != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.expiration_date", *body.ExpirationDate, goa.FormatDateTime))
	}
	return
}

// ValidateUpdateStudentRequestBody runs the validations defined on
// update_student_request_body
func ValidateUpdateStudentRequestBody(body *UpdateStudentRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Ruby == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ruby", "body"))
	}
	if body.StudentNumber == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("student_number", "body"))
	}
	if body.DateOfBirth == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("date_of_birth", "body"))
	}
	if body.Address == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("address", "body"))
	}
	if body.ExpirationDate == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("expiration_date", "body"))
	}
	if body.DateOfBirth != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.date_of_birth", *body.DateOfBirth, goa.FormatDateTime))
	}
	if body.ExpirationDate != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.expiration_date", *body.ExpirationDate, goa.FormatDateTime))
	}
	return
}
