/*
ResponseHelper
Peque#a libreria para facilitar las respuestas a los request sobre un ResponseWrite. Similar a como ASP .NET Core lo hace.

Pablo Ferreira 20200910
*/

package responseHelper

import (
	"encoding/json"
	"net/http"
)

// Estructura
type ResponseHelper struct {
	writer http.ResponseWriter
}

// Constructor
func New(writer http.ResponseWriter) *ResponseHelper {
	return &ResponseHelper{
		writer: writer,
	}
}

// Metodos
func (responseHelper *ResponseHelper) PrepareResponse(v interface{}) []byte {
	responseHelper.writer.Header().Set("Content-Type", "application/json")
	if v == nil {
		responseHelper.writer.WriteHeader(http.StatusNotFound)
		return nil
	}
	data, error := json.Marshal(v)
	if error != nil {
		responseHelper.writer.WriteHeader(http.StatusInternalServerError)
		return nil
	}
	return data
}

func (responseHelper *ResponseHelper) Ok(response interface{}) {
	data := responseHelper.PrepareResponse(response)
	if data == nil {
		return
	}
	responseHelper.writer.WriteHeader(http.StatusOK)
	responseHelper.writer.Write(data)
}

func (responseHelper *ResponseHelper) NotFound() {
	responseHelper.writer.WriteHeader(http.StatusNotFound)
}

func (responseHelper *ResponseHelper) BadRequest() {
	responseHelper.writer.WriteHeader(http.StatusBadRequest)
}

func (responseHelper *ResponseHelper) NoContent() {
	responseHelper.writer.WriteHeader(http.StatusNoContent)
}

func (responseHelper *ResponseHelper) Created(response interface{}) {
	data := responseHelper.PrepareResponse(response)
	if data == nil {
		return
	}
	responseHelper.writer.WriteHeader(http.StatusOK)
	responseHelper.writer.Write(data)
}
