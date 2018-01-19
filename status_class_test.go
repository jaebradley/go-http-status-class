package httpstatusclass

import (
	"net/http"
	"testing"
)

func TestInformational(t *testing.T) {
	sc, err := IdentifyStatusClass(http.StatusContinue)
	if err != nil {
		t.Errorf("No error expected")
	}

	if sc != Informational {
		t.Errorf("Expected Informational Status Class instead of %s", sc)
	}
}

func TestSuccessful(t *testing.T) {
	sc, err := IdentifyStatusClass(http.StatusOK)
	if err != nil {
		t.Errorf("No error expected")
	}

	if sc != Successful {
		t.Errorf("Expected Successful Status Class instead of %s", sc)
	}
}

func TestRedirection(t *testing.T) {
	sc, err := IdentifyStatusClass(http.StatusMultipleChoices)
	if err != nil {
		t.Errorf("No error expected")
	}

	if sc != Redirection {
		t.Errorf("Expected Redirection Status Class instead of %s", sc)
	}
}

func TestClientError(t *testing.T) {
	sc, err := IdentifyStatusClass(http.StatusBadRequest)
	if err != nil {
		t.Errorf("No error expected")
	}

	if sc != ClientError {
		t.Errorf("Expected Client Error Status Class instead of %s", sc)
	}
}

func TestServerError(t *testing.T) {
	sc, err := IdentifyStatusClass(http.StatusInternalServerError)
	if err != nil {
		t.Errorf("No error expected")
	}

	if sc != ServerError {
		t.Errorf("Expected Server Error Status Class instead of %s", sc)
	}
}

func TestUnknown(t *testing.T) {
	sc, err := IdentifyStatusClass(-1)
	if err == nil {
		t.Errorf("Error expected")
	}

	if sc != Unknown {
		t.Errorf("Expected Unknown Status Class instead of %s", sc)
	}
}

func TestInformationalResponse(t *testing.T) {
	var r http.Response
	r.StatusCode = http.StatusContinue
	sc, err := IdentifyResponseStatusClass(r)
	if err != nil {
		t.Errorf("No error expected")
	}

	if sc != Informational {
		t.Errorf("Expected Informational Status Class instead of %s", sc)
	}
}

func TestSuccessfulResponse(t *testing.T) {
	var r http.Response
	r.StatusCode = http.StatusOK
	sc, err := IdentifyResponseStatusClass(r)
	if err != nil {
		t.Errorf("No error expected")
	}

	if sc != Successful {
		t.Errorf("Expected Successful Status Class instead of %s", sc)
	}
}

func TestRedirectionResponse(t *testing.T) {
	var r http.Response
	r.StatusCode = http.StatusMultipleChoices
	sc, err := IdentifyResponseStatusClass(r)
	if err != nil {
		t.Errorf("No error expected")
	}

	if sc != Redirection {
		t.Errorf("Expected Redirection Status Class instead of %s", sc)
	}
}

func TestClientErrorResponse(t *testing.T) {
	var r http.Response
	r.StatusCode = http.StatusBadRequest
	sc, err := IdentifyResponseStatusClass(r)
	if err != nil {
		t.Errorf("No error expected")
	}

	if sc != ClientError {
		t.Errorf("Expected Client Error Status Class instead of %s", sc)
	}
}

func TestServerErrorResponse(t *testing.T) {
	var r http.Response
	r.StatusCode = http.StatusInternalServerError
	sc, err := IdentifyResponseStatusClass(r)
	if err != nil {
		t.Errorf("No error expected")
	}

	if sc != ServerError {
		t.Errorf("Expected Server Error Status Class instead of %s", sc)
	}
}

func TestUnknownResponse(t *testing.T) {
	var r http.Response
	r.StatusCode = -1
	sc, err := IdentifyResponseStatusClass(r)
	if err == nil {
		t.Errorf("Error expected")
	}

	if sc != Unknown {
		t.Errorf("Expected Unknown Status Class instead of %s", sc)
	}
}
