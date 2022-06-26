// Code generated by goa v3.7.6, DO NOT EDIT.
//
// students HTTP server
//
// Command:
// $ goa gen github.com/is-hoku/goa-template/webapi/design

package server

import (
	"context"
	"net/http"

	students "github.com/is-hoku/goa-template/gen/students"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the students service endpoint HTTP handlers.
type Server struct {
	Mounts        []*MountPoint
	GetStudent    http.Handler
	GetStudents   http.Handler
	CreateStudent http.Handler
	CORS          http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the students service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *students.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"GetStudent", "GET", "/students/{id}"},
			{"GetStudents", "GET", "/students"},
			{"CreateStudent", "POST", "/students"},
			{"CORS", "OPTIONS", "/students/{id}"},
			{"CORS", "OPTIONS", "/students"},
		},
		GetStudent:    NewGetStudentHandler(e.GetStudent, mux, decoder, encoder, errhandler, formatter),
		GetStudents:   NewGetStudentsHandler(e.GetStudents, mux, decoder, encoder, errhandler, formatter),
		CreateStudent: NewCreateStudentHandler(e.CreateStudent, mux, decoder, encoder, errhandler, formatter),
		CORS:          NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "students" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.GetStudent = m(s.GetStudent)
	s.GetStudents = m(s.GetStudents)
	s.CreateStudent = m(s.CreateStudent)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the students endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountGetStudentHandler(mux, h.GetStudent)
	MountGetStudentsHandler(mux, h.GetStudents)
	MountCreateStudentHandler(mux, h.CreateStudent)
	MountCORSHandler(mux, h.CORS)
}

// Mount configures the mux to serve the students endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountGetStudentHandler configures the mux to serve the "students" service
// "get student" endpoint.
func MountGetStudentHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleStudentsOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/students/{id}", f)
}

// NewGetStudentHandler creates a HTTP handler which loads the HTTP request and
// calls the "students" service "get student" endpoint.
func NewGetStudentHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetStudentRequest(mux, decoder)
		encodeResponse = EncodeGetStudentResponse(encoder)
		encodeError    = EncodeGetStudentError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "get student")
		ctx = context.WithValue(ctx, goa.ServiceKey, "students")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGetStudentsHandler configures the mux to serve the "students" service
// "get students" endpoint.
func MountGetStudentsHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleStudentsOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/students", f)
}

// NewGetStudentsHandler creates a HTTP handler which loads the HTTP request
// and calls the "students" service "get students" endpoint.
func NewGetStudentsHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeGetStudentsResponse(encoder)
		encodeError    = EncodeGetStudentsError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "get students")
		ctx = context.WithValue(ctx, goa.ServiceKey, "students")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCreateStudentHandler configures the mux to serve the "students" service
// "create student" endpoint.
func MountCreateStudentHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleStudentsOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/students", f)
}

// NewCreateStudentHandler creates a HTTP handler which loads the HTTP request
// and calls the "students" service "create student" endpoint.
func NewCreateStudentHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeCreateStudentResponse(encoder)
		encodeError    = EncodeCreateStudentError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "create student")
		ctx = context.WithValue(ctx, goa.ServiceKey, "students")
		var err error
		res, err := endpoint(ctx, nil)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service students.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleStudentsOrigin(h)
	mux.Handle("OPTIONS", "/students/{id}", h.ServeHTTP)
	mux.Handle("OPTIONS", "/students", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// HandleStudentsOrigin applies the CORS response headers corresponding to the
// origin for the service students.
func HandleStudentsOrigin(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			h.ServeHTTP(w, r)
			return
		}
		if cors.MatchOrigin(origin, "http://localhost:8017") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "X-Time")
			w.Header().Set("Access-Control-Max-Age", "600")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			}
			h.ServeHTTP(w, r)
			return
		}
		h.ServeHTTP(w, r)
		return
	})
}
