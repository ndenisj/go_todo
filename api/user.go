package api

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
	db "github.com/ndenisj/go_todo/db/sqlc"
	"github.com/ndenisj/go_todo/utils"
)

type createUserRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Phone    string `json:"phone" binding:"required,e164"`
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type createUserResponse struct {
	ID                int64     `json:"id"`
	FullName          string    `json:"fullName"`
	Phone             string    `json:"phone"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"passwordChangedAt"`
	IsAdmin           bool      `json:"isAdmin"`
	CreatedAt         time.Time `json:"createdAt"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		var ve validator.ValidationErrors
		log.Println(err)
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		// ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	hashedPassword, err := utils.HashedPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}

	arg := db.CreateUserParams{
		FullName: req.FullName,
		Phone: sql.NullString{
			String: req.Phone,
			Valid:  true,
		},
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: hashedPassword,
	}
	// Check validations

	//
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		// if isUniqueContraintViolation(err) {
		// 	ctx.JSON(http.StatusForbidden, errorResponse("Duplicate records"))
		// 	return
		// }

		if pqErr, ok := err.(*pq.Error); ok {
			log.Println(pqErr.Code)
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(duplicateError(err.Error())))
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}

	response := createUserResponse{
		ID:                user.ID,
		FullName:          user.FullName,
		Phone:             user.Phone.String,
		Username:          user.Username,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		IsAdmin:           user.IsAdmin,
		CreatedAt:         user.CreatedAt,
	}

	ctx.JSON(http.StatusCreated, successResponse("User created successfully!", response))
}
