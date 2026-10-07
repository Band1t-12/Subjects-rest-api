package myhttp

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	HTTPHandlers *HTTPHandlers
}

func NewServer(HTTPHandlers *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		HTTPHandlers: HTTPHandlers,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/subjects").Methods("POST").HandlerFunc(s.HTTPHandlers.HandleAddSubject)
	router.Path("/subjects/{id}").Methods("GET").HandlerFunc(s.HTTPHandlers.HandleGetSubject)
	router.Path("/subjects").Methods("GET").HandlerFunc(s.HTTPHandlers.HandleGetAllSubjects)
	router.Path("/subjects/{id}").Methods("PUT").HandlerFunc(s.HTTPHandlers.HandleUpdateSubject)
	router.Path("/subjects/{id}").Methods("DELETE").HandlerFunc(s.HTTPHandlers.HandleDeleteSubject)
	router.Path("/teachers").Methods("POST").HandlerFunc(s.HTTPHandlers.HandleAddTeacher)
	router.Path("/teachers/{id}").Methods("PUT").HandlerFunc(s.HTTPHandlers.HandleUpdateTeacher)

	if err := http.ListenAndServe(":8080", LoggingMiddleware(router)); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}

	return nil
}
