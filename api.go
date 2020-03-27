package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

var qe queryExecutor

type idPathRoute struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type requestContext interface {
	ShouldBindUri(obj interface{}) error
	JSON(code int, obj interface{})
}

func getEntityByIDRoute(c requestContext, qe queryExecutor, collectionName string) {
	var idRoute idPathRoute
	err := c.ShouldBindUri(&idRoute)
	if err != nil || idRoute.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Invalid ID;  Must be a UUID"}})
		return
	}

	data, err := getEntityByID(qe, collectionName, idRoute.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": []string{"Request failed"}})
		return
	}

	c.JSON(http.StatusOK, data)
}

func configureRoutes(router *gin.Engine) {
	qe := arangoDbQueryExecutor{}
	router.GET("/users/:id", func(c *gin.Context) { getEntityByIDRoute(c, &qe, "users") })
	router.GET("/companies/:id", func(c *gin.Context) { getEntityByIDRoute(c, &qe, "companies") })
}

func getServerAddr() string {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if host == "" {
		host = "127.0.0.1"
	}
	if port == "" {
		port = "8080"
	}
	addr := host + ":" + port

	return addr
}

func initializeServer(router *gin.Engine) *http.Server {
	addr := getServerAddr()

	svr := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	println("Starting Server At Address: " + addr)

	return svr
}

func main() {
	router := gin.Default()
	configureRoutes(router)
	svr := initializeServer(router)
	svr.ListenAndServe()
}
