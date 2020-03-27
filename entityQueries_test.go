package main

import (
	"testing"
)

type mockQueryExecutor struct {
	queryExecutor
}

func (qe *mockQueryExecutor) execute(_ string, _ map[string]interface{}) ([]map[string]interface{}, error) {
	data := make([]map[string]interface{}, 1)
	data[0] = map[string]interface{} {
		"_key": "abc",
	}
	return data, nil
}

func TestGetEntityByID(t *testing.T) {
	qe := mockQueryExecutor{}

	results, err := getEntityByID(&qe, "Test Query", "")
	if err != nil {
		t.Fail()
	}

	if results == nil {
		t.Fail()
	}
}