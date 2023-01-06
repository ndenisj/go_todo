package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/ndenisj/go_todo/db/sqlc"
)

// Server: serve http request for the todo services
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer: create a new server instance and a new HTTP server and routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// add route to the server
	v1 := router.Group("/v1")
	{
		v1.POST("/todos", server.createTodo)
		v1.GET("/todos/:id", server.getTodo)
		v1.GET("/todos", server.listTodos)
		v1.DELETE("/todos/:id", server.deleteTodo)
		v1.PUT("/todos", server.updateTodo)
	}
	// router.POST("/todos", server.createTodo)
	// router.GET("/todos/:id", server.getTodo)
	// router.GET("/todos", server.listTodos)
	// router.DELETE("/todos/:id", server.deleteTodo)
	// router.PUT("/todos", server.updateTodo)

	server.router = router
	return server

}

// Start: runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err_message string) gin.H {
	return gin.H{
		"status":  false,
		"message": err_message,
	}
}

func successResponse(message string, data interface{}) gin.H {
	return gin.H{
		"status":  true,
		"message": message,
		"data":    data,
	}
}
