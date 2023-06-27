package users

import (
	"github.com/gin-gonic/gin"
	"github.com/obadalov/bookstore-users-api/domain/users"
	"github.com/obadalov/bookstore-users-api/services"
	"github.com/obadalov/bookstore-users-api/utils/errors"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: return bad request to user
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("User id should be a number")
		c.JSON(restErr.Status, restErr)
		return
	}

	user, e := services.GetUser(userId)
	if e != nil {
		c.JSON(e.Status, e)
		return
	}
	c.JSON(http.StatusCreated, user)
}
