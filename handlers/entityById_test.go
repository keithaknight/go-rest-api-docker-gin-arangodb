package handlers

import (
	"net/http"
	"reflect"
	"testing"
	"github.com/keithaknight/go-rest-api-docker-gin-arangodb/queries"
)

type mockRequestContext struct {
	code int
	obj  interface{}
	expectedID string
	RequestContext
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
	qe := queries.MockQueryExecutor{}
	GetEntityByIDRoute(&c, &qe, "firms")

	if c.code != http.StatusBadRequest {
		t.Fail()
	}

	if c.obj == nil {
		t.Fail()
	}
}

func TestGetEntityByIDRouteWithId(t *testing.T) {
	c := mockRequestContext{ expectedID: "testId"}
	qe := queries.MockQueryExecutor{}
	GetEntityByIDRoute(&c, &qe, "firms")

	if c.code != http.StatusOK {
		t.Fail()
	}

	if c.obj == nil {
		t.Fail()
	}
}
