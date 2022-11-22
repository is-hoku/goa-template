// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package sample

import (
	"context"
	"github.com/is-hoku/goa-sample/webapi/datastore"
	"github.com/is-hoku/goa-sample/webapi/interactor"
)

// Injectors from wire.go:

func NewStudentApp(ctx context.Context) (*StudentApp, error) {
	config := newMySQLConfig()
	db, err := datastore.NewMySQL(config)
	if err != nil {
		return nil, err
	}
	getStudentByNumberMedia := datastore.NewGetStudentByNumberMedia(db)
	getStudentByNumberOption := &interactor.GetStudentByNumberOption{
		StudentByNumberGetter: getStudentByNumberMedia,
	}
	getStudentByNumber := interactor.NewGetStudentByNumber(getStudentByNumberOption)
	getStudentsMedia := datastore.NewGetStudentsMedia(db)
	getStudentsOption := &interactor.GetStudentsOption{
		StudentsGetter: getStudentsMedia,
	}
	getStudents := interactor.NewGetStudents(getStudentsOption)
	getStudentByIDMedia := datastore.NewGetStudentByIDMedia(db)
	createStudentMedia := datastore.NewCreateStudentMedia(db)
	createStudentOption := &interactor.CreateStudentOption{
		StudentByIDGetter: getStudentByIDMedia,
		StudentCreator:    createStudentMedia,
	}
	createStudent := interactor.NewCreateStudent(createStudentOption)
	updateStudentMedia := datastore.NewUpdateStudentMedia(db)
	updateStudentOption := &interactor.UpdateStudentOption{
		StudentUpdater:        updateStudentMedia,
		StudentByNumberGetter: getStudentByNumberMedia,
	}
	updateStudent := interactor.NewUpdateStudent(updateStudentOption)
	deleteStudentMedia := datastore.NewDeleteStudentMedia(db)
	deleteStudentOption := &interactor.DeleteStudentOption{
		StudentDeleter: deleteStudentMedia,
	}
	deleteStudent := interactor.NewDeleteStudent(deleteStudentOption)
	studentApp, err := newStudentApp(ctx, getStudentByNumber, getStudents, createStudent, updateStudent, deleteStudent)
	if err != nil {
		return nil, err
	}
	return studentApp, nil
}
