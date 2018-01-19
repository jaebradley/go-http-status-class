[![Go Report Card](https://goreportcard.com/badge/github.com/jaebradley/go-http-status-class)](https://goreportcard.com/report/github.com/jaebradley/go-http-status-class)

# go-http-status-class

## Introduction

A very simple API that returns [different `HTTP Status Classes`](https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html) for [various `HTTP Status Codes`](https://golang.org/src/net/http/status.go) (like `Successful` for `http.StatusOK` or `ClientError` for `http.StatusBadRequest`).

## Usage
```golang
import (
  "net/http"

  "github.com/jaebradley/go-http-status-class"
)

var r http.Response;
r.StatusCode = http.StatusOK
statusClass, err := httpstatusclass.IdentifyResponseStatusClass(r)

if err != nil {
  fmt.Printf("Unknown Status Class: %s\n", r.Status)
}

fmt.Println(statusClass) // This should print "Successful"

statusClass2, err2 := httpstatusclass.IdentifyStatusClass(-1)

if err2 != nil {
  log.Fatal(err) // This will print "Unexpected status code: -1"
}
```
