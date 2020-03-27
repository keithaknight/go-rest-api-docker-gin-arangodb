package main

func getEntityByID(executor queryExecutor, collectionName string, id string) (map[string][]map[string]interface{}, error) {
	queryResults, err := executor.execute(`
		FOR u IN @@coll
		FILTER u._key == @uId
		RETURN u`, map[string]interface{}{
		"@coll": collectionName,
		"uId":   id,
	})

	if err != nil {
		return nil, err
	}

	data := map[string][]map[string]interface{}{
		"data": queryResults,
	}

	return data, nil
}