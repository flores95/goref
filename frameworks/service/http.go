package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/flores95/goref/frameworks/config"
)

// HTTPServicer is an http endpoint for a REST or other http based API
type HTTPServicer struct {
	port   string
	Router *mux.Router
}

// NewHTTPServicer provides a Servicer with the given configuration
func NewHTTPServicer(conf config.Configurator) (s HTTPServicer) {
	s.port = ":8080"
	if p := conf.GetValue("PORT"); p != "" {
		s.port = fmt.Sprintf(":%v", p)
	}
	s.Router = mux.NewRouter()
	return s
}

// Read attempts to unmarshal the incoming json body into the given interface
func (s *HTTPServicer) Read(r *http.Request, obj interface{}) {
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, obj)
}

// Respond will marshal the given interface into json and send it with the provided response code
func (s *HTTPServicer) Respond(writer http.ResponseWriter, obj interface{}, statusCode int) {
	writer.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(obj)
	writer.WriteHeader(statusCode)
	writer.Write(j)
}

// Health is used to verify that the service is functioning
func (s *HTTPServicer) Health() Status {
	return Up
}

// Start in this case just gives the unique URI of the service
func (s *HTTPServicer) Start() {
	http.ListenAndServe(s.port, s.Router)
}
