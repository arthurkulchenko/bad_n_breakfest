package main

import (
	"fmt"
	"testing"
	"net/http"
)

func TestNoSurve(test *testing.T) {
	var myH myHandler
	handler := NoSurf(&myH)

	switch v := handler.(type) {
	case http.Handler:
		// do nothing
	default:
		test.Error(fmt.Sprintf("type is not http.Handler, but is %T", v))
	}
}
