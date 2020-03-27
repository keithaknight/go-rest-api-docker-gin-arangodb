package main

import (
	"testing"
)

func TestArangoDbQueryExecutor(t *testing.T) {
	config := newArangoDbConfig()

	var qe queryExecutor
	if config.password != "" {
		qe = &arangoDbQueryExecutor{}
	} else {
		qe = &mockQueryExecutor{}
	}

	results, err := qe.execute("FOR u IN users LIMIT 1 RETURN u", nil)
	if err != nil {
		t.Fail()
	}

	if results == nil || len(results) == 0 {
		t.Fail()
	}

	if results[0]["_key"] == nil {
		t.Fail()
	}
}
