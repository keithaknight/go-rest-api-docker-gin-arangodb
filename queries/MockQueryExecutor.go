package queries

//MockQueryExecutor provides a mock query executor that can be used for testing
type MockQueryExecutor struct {
	QueryExecutor
}

//Execute returns static data that can be used for testing
func (qe *MockQueryExecutor) Execute(_ string, _ map[string]interface{}) ([]map[string]interface{}, error) {
	data := make([]map[string]interface{}, 1)
	data[0] = map[string]interface{} {
		"_key": "abc",
	}
	return data, nil
}
