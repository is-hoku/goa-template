// Code generated by goa v3.7.6, DO NOT EDIT.
//
// students HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/is-hoku/goa-template/webapi/design

package server

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	students "github.com/is-hoku/goa-template/gen/students"
	studentsviews "github.com/is-hoku/goa-template/gen/students/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeGetStudentResponse returns an encoder for responses returned by the
// students get student endpoint.
func EncodeGetStudentResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*studentsviews.Student)
		enc := encoder(ctx, w)
		body := NewGetStudentResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetStudentRequest returns a decoder for requests sent to the students
// get student endpoint.
func DecodeGetStudentRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  int64
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseInt(idRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "integer"))
			}
			id = v
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetStudentPayload(id)

		return payload, nil
	}
}

// EncodeGetStudentError returns an encoder for errors returned by the get
// student students endpoint.
func EncodeGetStudentError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "internal_error":
			var res *students.CustomError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetStudentInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "not_found":
			var res *students.CustomError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetStudentNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetStudentsResponse returns an encoder for responses returned by the
// students get students endpoint.
func EncodeGetStudentsResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*studentsviews.Students)
		enc := encoder(ctx, w)
		body := NewGetStudentsResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeGetStudentsError returns an encoder for errors returned by the get
// students students endpoint.
func EncodeGetStudentsError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "internal_error":
			var res *students.CustomError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetStudentsInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeCreateStudentResponse returns an encoder for responses returned by the
// students create student endpoint.
func EncodeCreateStudentResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*studentsviews.Student)
		enc := encoder(ctx, w)
		body := NewCreateStudentResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeCreateStudentError returns an encoder for errors returned by the
// create student students endpoint.
func EncodeCreateStudentError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "internal_error":
			var res *students.CustomError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewCreateStudentInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "bad_request":
			var res *students.CustomError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewCreateStudentBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalStudentsviewsStudentViewToStudentResponseBody builds a value of type
// *StudentResponseBody from a value of type *studentsviews.StudentView.
func marshalStudentsviewsStudentViewToStudentResponseBody(v *studentsviews.StudentView) *StudentResponseBody {
	res := &StudentResponseBody{
		ID:             *v.ID,
		Name:           *v.Name,
		Ruby:           *v.Ruby,
		StudentNumber:  *v.StudentNumber,
		DateOfBirth:    *v.DateOfBirth,
		Address:        *v.Address,
		ExpirationDate: *v.ExpirationDate,
	}

	return res
}
