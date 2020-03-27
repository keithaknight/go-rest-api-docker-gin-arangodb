package main

import (
	"net/http"
	"os"
	"reflect"
	"testing"
)

type mockRequestContext struct {
	code int
	obj  interface{}
	expectedID string
	requestContext
}

func (c *mockRequestContext) JSON(code int, obj interface{}) {
	c.code = code
	c.obj = obj
}

func (c *mockRequestContext) ShouldBindUri(obj interface{}) error {
	ps := reflect.ValueOf(obj)
	s := ps.Elem()
	f := s.FieldByName("ID")
	f.SetString(c.expectedID)

	return nil
}

func TestGetEntityByIDRouteNullId(t *testing.T) {
	c := mockRequestContext{ expectedID: ""}
	qe := mockQueryExecutor{}
	getEntityByIDRoute(&c, &qe, "firms")

	if c.code != http.StatusBadRequest {
		t.Fail()
	}

	if c.obj == nil {
		t.Fail()
	}
}

func TestGetEntityByIDRouteWithId(t *testing.T) {
	c := mockRequestContext{ expectedID: "testId"}
	qe := mockQueryExecutor{}
	getEntityByIDRoute(&c, &qe, "firms")

	if c.code != http.StatusOK {
		t.Fail()
	}

	if c.obj == nil {
		t.Fail()
	}
}

func TestGetServerAddrDefault(t *testing.T) {
	os.Setenv("HOST", "")
	os.Setenv("PORT", "")
	addr := getServerAddr()

	if addr != "127.0.0.1:8080" {
		t.Fail()
	}
}

func TestGetServerAddrEnvVar(t *testing.T) {
	os.Setenv("HOST", "0.0.0.0")
	os.Setenv("PORT", "80")
	addr := getServerAddr()

	if addr != "0.0.0.0:80" {
		t.Fail()
	}
}