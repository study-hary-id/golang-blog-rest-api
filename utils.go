package main

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

// responseJSON construct generic `data` into json bytes and send response back.
func responseJSON(
	w http.ResponseWriter,
	_ *http.Request,
	data interface{},
	statusCode int,
) {
	jsonBytes, err := json.Marshal(constructItems(data))
	if err != nil {
		internalServerError(w, err)
		return
	}
	w.WriteHeader(statusCode)
	_, err = w.Write(jsonBytes)
	if err != nil {
		internalServerError(w, err)
		return
	}
}

// internalServerError gives InternalServerError response.
func internalServerError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// getRequestType returns the type of request based on the URL.
func getRequestType(r *http.Request) string {
	return strings.Split(r.URL.Path, "/")[1]
}

// URLConv converts string to proper data types,
// since the values from URL queries are always string.
func URLConv(value string) interface{} {
	var (
		boolRe = regexp.MustCompile(`^([Tt][Rr][Uu][Ee]|[Ff][Aa][Ll][Ss][Ee])$`)
		intRe  = regexp.MustCompile(`^\d+$`)
	)
	if boolRe.MatchString(value) {
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		return boolValue
	} else if intRe.MatchString(value) {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		return intValue
	}
	return nil
}
