package rest

import (
	"encoding/json"
	"net/http"
)

func ResponseSuccess(resp http.ResponseWriter, payload interface{}) {
	responseJSON(resp, http.StatusOK, payload)
}

func ResponseCreated(resp http.ResponseWriter, payload interface{}) {
	responseJSON(resp, http.StatusCreated, payload)
}

func ResponseBadRequest(resp http.ResponseWriter) {
	responseJSON(resp, http.StatusBadRequest, map[string]string { "error": "Bad Request" })
}

func ResponseUnauthorized(resp http.ResponseWriter) {
	responseJSON(resp, http.StatusUnauthorized, map[string]string { "error": "Unauthorized" })
}


func ResponseNotFound(resp http.ResponseWriter) {
	responseJSON(resp, http.StatusNotFound, map[string]string { "error": "Data not found" })
}

func ResponseInternalServerError(resp http.ResponseWriter) {
	responseJSON(resp, http.StatusInternalServerError, map[string]string { "error": "Server internal error" })
}

func responseJSON(respWriter http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	respWriter.Header().Set("Content-Type", "application/json")
	respWriter.WriteHeader(code)
	respWriter.Write(response)
}