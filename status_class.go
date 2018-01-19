package httpstatusclass

import (
	"fmt"
	"net/http"
)

// StatusClass https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html
type StatusClass string

/*
	Informational Status Class: 1xx Response Codes
	Successful Status Class: 2xx Response Codes
	Rediction Status Class: 3xx Response Codes
	Client Error Status Class: 4xx Response Codes
	Server Error Status Class: 5xx Response Codes
	Unknown Status Class: Any Other Response Code
*/
const (
	Informational StatusClass = "Informational"
	Successful    StatusClass = "Successful"
	Redirection   StatusClass = "Redirection"
	ClientError   StatusClass = "Client Error"
	ServerError   StatusClass = "Server Error"
	Unknown       StatusClass = "Unknown"
)

// IdentifyStatusClass returns a StatusClass for a given http.Response object.
// For a http.Response object with an unexpected StatusCode, the returned
// error is non-nil and the StatusClass is Unknown.
func IdentifyResponseStatusClass(r http.Response) (StatusClass, error) {
	return IdentifyStatusClass(r.StatusCode)
}

// IdentifyStatusClass returns a StatusClass for a given integer status code.
// For a http.Response object with an unexpected StatusCode, the returned
// error is non-nil and the StatusClass is Unknown.
func IdentifyStatusClass(sc int) (StatusClass, error) {
	if sc >= 100 && sc < 200 {
		return Informational, nil
	} else if sc >= 200 && sc < 300 {
		return Successful, nil
	} else if sc >= 300 && sc < 400 {
		return Redirection, nil
	} else if sc >= 400 && sc < 500 {
		return ClientError, nil
	} else if sc >= 500 && sc < 600 {
		return ServerError, nil
	}

	return Unknown, fmt.Errorf("Unexpected status code: %d", sc)
}
