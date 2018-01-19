package httpstatusclass

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInformational(t *testing.T) {
	sc, err := IdentifyStatusClass(http.StatusContinue)
	assert.Nil(t, err, "No error expected")
	assert.Equalf(t, Informational, sc, "Expected Informational Status Class instead of %s", sc)
}

func TestSuccessful(t *testing.T) {
	sc, err := IdentifyStatusClass(http.StatusOK)
	assert.Nil(t, err, "No error expected")
	assert.Equalf(t, Successful, sc, "Expected Successful Status Class instead of %s", sc)
}

func TestRedirection(t *testing.T) {
	sc, err := IdentifyStatusClass(http.StatusMultipleChoices)
	assert.Nil(t, err, "No error expected")
	assert.Equal(t, Redirection, sc, "Expected Redirection Status Class instead of %s", sc)
}

func TestClientError(t *testing.T) {
	sc, err := IdentifyStatusClass(http.StatusBadRequest)
	assert.Nil(t, err, "No error expected")
	assert.Equal(t, ClientError, sc, "Expected Client Error Status Class instead of %s", sc)
}

func TestServerError(t *testing.T) {
	sc, err := IdentifyStatusClass(http.StatusInternalServerError)
	assert.Nil(t, err, "No error expected")
	assert.Equal(t, ServerError, sc, "Expected Server Error Status Class instead of %s", sc)
}

func TestUnknown(t *testing.T) {
	sc, err := IdentifyStatusClass(-1)
	assert.NotNil(t, err, "No error expected")
	assert.Equal(t, Unknown, sc, "Expected Unknown Status Class instead of %s", sc)
}

func TestInformationalResponse(t *testing.T) {
	var r http.Response
	r.StatusCode = http.StatusContinue
	sc, err := IdentifyResponseStatusClass(r)
	assert.Nil(t, err, "No error expected")
	assert.Equalf(t, Informational, sc, "Expected Informational Status Class instead of %s", sc)
}

func TestSuccessfulResponse(t *testing.T) {
	var r http.Response
	r.StatusCode = http.StatusOK
	sc, err := IdentifyResponseStatusClass(r)
	assert.Nil(t, err, "No error expected")
	assert.Equalf(t, Successful, sc, "Expected Successful Status Class instead of %s", sc)
}

func TestRedirectionResponse(t *testing.T) {
	var r http.Response
	r.StatusCode = http.StatusMultipleChoices
	sc, err := IdentifyResponseStatusClass(r)
	assert.Nil(t, err, "No error expected")
	assert.Equalf(t, Redirection, sc, "Expected Redirection Status Class instead of %s", sc)
}

func TestClientErrorResponse(t *testing.T) {
	var r http.Response
	r.StatusCode = http.StatusBadRequest
	sc, err := IdentifyResponseStatusClass(r)
	assert.Nil(t, err, "No error expected")
	assert.Equalf(t, ClientError, sc, "Expected Client Error Status Class instead of %s", sc)
}

func TestServerErrorResponse(t *testing.T) {
	var r http.Response
	r.StatusCode = http.StatusInternalServerError
	sc, err := IdentifyResponseStatusClass(r)
	assert.Nil(t, err, "No error expected")
	assert.Equal(t, ServerError, sc, "Expected Server Error Status Class instead of %s", sc)
}

func TestUnknownResponse(t *testing.T) {
	var r http.Response
	r.StatusCode = -1
	sc, err := IdentifyResponseStatusClass(r)
	assert.NotNil(t, err, "No error expected")
	assert.Equal(t, Unknown, sc, "Expected Unknown Class instead of %s", sc)
}
