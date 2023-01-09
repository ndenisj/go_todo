package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/ndenisj/go_todo/db/sqlc"
)

type createTodoRequest struct {
	Owner   string `json:"owner" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// createTodo: receive request with payload process and send to db
func (server *Server) createTodo(ctx *gin.Context) {
	var req createTodoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	arg := db.CreateTodoParams{
		Owner:   req.Owner,
		Title:   req.Title,
		Content: req.Content,
	}

	todo, err := server.store.CreateTodo(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, successResponse("successful", todo))
	// ctx.JSON(http.StatusOK, todo)
}

type getTodoRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getTodo(ctx *gin.Context) {
	var req getTodoRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	todo, err := server.store.GetTodo(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse("Todo not found"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, todo)
	// ctx.JSON(http.StatusOK, successResponse("successful", todo))
}

type listTodosRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listTodos(ctx *gin.Context) {
	var req listTodosRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	arg := db.ListTodosParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	todos, err := server.store.ListTodos(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, successResponse("Successful", todos))
}

type deleteTodoRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteTodo(ctx *gin.Context) {
	var req deleteTodoRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	err := server.store.DeleteTodo(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, successResponse("Deleted", nil))
}

type updateTodoRequest struct {
	ID      int64   `json:"id" binding:"required"`
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

func (server *Server) updateTodo(ctx *gin.Context) {
	var req updateTodoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	arg := db.UpdateTodoParams{
		ID: req.ID,
		Title: sql.NullString{
			String: req.getTitle(),
			Valid:  req.getTitle() != "",
		},
		Content: sql.NullString{
			String: req.getContent(),
			Valid:  req.getContent() != "",
		},
	}
	todo, err := server.store.UpdateTodo(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, successResponse("updated", todo))
}

func (x *updateTodoRequest) getTitle() string {
	if x.Title != nil {
		return *x.Title
	}
	return ""
}
func (x *updateTodoRequest) getContent() string {
	if x.Content != nil {
		return *x.Content
	}
	return ""
}
