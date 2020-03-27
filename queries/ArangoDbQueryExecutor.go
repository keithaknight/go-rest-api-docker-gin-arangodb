package queries

import (
	"context"
	"crypto/tls"
	driver "github.com/arangodb/go-driver"
	arangoHttp "github.com/arangodb/go-driver/http"
)

//ArangoDbQueryExecutor is the primary type for executing queries
type ArangoDbQueryExecutor struct {
	db driver.Database
	QueryExecutor
}

func initializeArangoDb() (driver.Database, error) {
	config := newArangoDbConfig()
	conn, err := arangoHttp.NewConnection(arangoHttp.ConnectionConfig{
		Endpoints: config.endpointUrls,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
	})
	if err != nil {
		return nil, err
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(config.userName, config.password),
	})
	if err != nil {
		return nil, err
	}

	db, err := client.Database(context.Background(), config.databaseName)
	if err != nil {
		return nil, err
	}

	return db, nil
}

//Execute runs the specified query
func (q *ArangoDbQueryExecutor) Execute(queryText string, bindVars map[string]interface{}) ([]map[string]interface{}, error) {
	if q.db == nil {
		db, err := initializeArangoDb()
		if err != nil {
			return nil, err
		}
		q.db = db
	}

	ctx := driver.WithQueryCount(context.Background())
	cursor, err := q.db.Query(ctx, queryText, bindVars)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	count := cursor.Count()
	data := make([]map[string]interface{}, count)

	idx := 0
	for {
		var doc map[string]interface{}
		_, err := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			// handle other errors
		}

		data[idx] = doc
		idx++
	}

	return data, nil
}

//NewQueryExecutor returns a new query executor
func NewQueryExecutor() ArangoDbQueryExecutor {
	qe := ArangoDbQueryExecutor{}
	return qe
}
