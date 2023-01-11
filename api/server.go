package api

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	db "github.com/ndenisj/go_todo/db/sqlc"
)

// Server: serve http request for the todo services
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer: create a new server instance and a new HTTP server and routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// add route to the server
	v1 := router.Group("/v1")
	{
		// todo
		v1.POST("/todos", server.createTodo)
		v1.GET("/todos/:id", server.getTodo)
		v1.GET("/todos", server.listTodos)
		v1.DELETE("/todos/:id", server.deleteTodo)
		v1.PUT("/todos", server.updateTodo)
		// user
		v1.POST("/user", server.createUser)
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

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	case "e164":
		return "Should be a valid phone number with country code"
	}
	return "Unknown error"
}

// func isUniqueContraintViolation(err error) bool {
// 	if pgError, ok := err.(*pq.Error); ok && errors.Is(err, pgError) {
// 		if pgError.Code == "23505" {
// 			return true
// 		}
// 	}

// 	return false
// }

func duplicateError(err string) string {
	if strings.Contains(err, "users_username_key") {
		return "username already taken"
	} else if strings.Contains(err, "users_phone_key") {
		return "phone already taken"
	} else if strings.Contains(err, "users_email_key") {
		return "email already taken"
	}

	return "Duplicate record"
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
