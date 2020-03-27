package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/keithaknight/go-rest-api-docker-gin-arangodb/queries"
)

type idPathRoute struct {
	ID string `uri:"id" binding:"required,uuid"`
}

//RequestContext abstracts the gin.Context through an interface to simplify mocking for unit tests
type RequestContext interface {
	ShouldBindUri(obj interface{}) error
	JSON(code int, obj interface{})
}

//GetEntityByIDRoute handles a GET for a specific entity by id (example: GET /users/:id)
func GetEntityByIDRoute(c RequestContext, qe queries.QueryExecutor, collectionName string) {
	var idRoute idPathRoute
	err := c.ShouldBindUri(&idRoute)
	if err != nil || idRoute.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Invalid ID;  Must be a UUID"}})
		return
	}

	data, err := queries.GetEntityByID(qe, collectionName, idRoute.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": []string{"Request failed"}})
		return
	}

	c.JSON(http.StatusOK, data)
}