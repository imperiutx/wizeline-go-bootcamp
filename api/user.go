package api

import (
	"context"
	"net/http"
	"time"

	db "wizeline-go-bootcamp/db/sqlc"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type userResponse struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

func (server *Server) createUser(ctx echo.Context) error {
	var req createUserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	arg := db.CreateUserParams{
		Username: req.Username,
		Email:    req.Email,
	}

	user, err := server.store.CreateUser(context.Background(), arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return ctx.JSON(http.StatusForbidden, err.Error())
			}
		}
		return ctx.JSON(http.StatusInternalServerError, err.Error())

	}

	rsp := newUserResponse(user)
	return ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) getUser(ctx echo.Context) error {
	user, err := server.store.GetUser(context.Background(), ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	rsp := newUserResponse(user)
	return ctx.JSON(http.StatusOK, rsp)
}
