package queries

import (
	"testing"
)

func TestArangoDbQueryExecutor(t *testing.T) {
	config := newArangoDbConfig()

	var qe QueryExecutor
	if config.password != "" {
		qe = &ArangoDbQueryExecutor{}
	} else {
		qe = &MockQueryExecutor{}
	}

	results, err := qe.Execute("FOR u IN users LIMIT 1 RETURN u", nil)
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
