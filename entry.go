package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
	"github.com/keithaknight/go-rest-api-docker-gin-arangodb/handlers"
	"github.com/keithaknight/go-rest-api-docker-gin-arangodb/queries"
)

func configureRoutes(router *gin.Engine) {
	qe := queries.NewQueryExecutor()
	router.GET("/users/:id", func(c *gin.Context) { handlers.GetEntityByIDRoute(c, &qe, "users") })
	router.GET("/companies/:id", func(c *gin.Context) { handlers.GetEntityByIDRoute(c, &qe, "companies") })
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
