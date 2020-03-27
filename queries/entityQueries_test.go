package queries

import (
	"testing"
)

func TestGetEntityByID(t *testing.T) {
	qe := MockQueryExecutor{}

	results, err := GetEntityByID(&qe, "Test Query", "")
	if err != nil {
		t.Fail()
	}

	if results == nil {
		t.Fail()
	}
}