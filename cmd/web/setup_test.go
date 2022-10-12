package main

import(
	"os"
	"testing"
	"net/http"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

type myHandler struct{}

func (mh *myHandler) ServeHTTP(responce http.ResponseWriter, request *http.Request) {}
